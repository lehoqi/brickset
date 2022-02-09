/**
 * @Package storage
 * @Time: 2022/2/9 9:37 PM
 * @Author: wuhb
 * @File: memory.go
 */

package storage

import (
	"context"
	"errors"
	"sync"
	"time"
)

type memoryCmd struct {
	val    interface{}
	expire time.Time
}
type memory struct {
	data map[string]*memoryCmd
	lock sync.Locker
}

func (m memory) Set(_ context.Context, key string, value interface{}, expiration time.Duration) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data[key] = &memoryCmd{
		val:    value,
		expire: time.Now().Add(expiration),
	}
	return nil
}

func (m memory) Get(_ context.Context, key string) (interface{}, error) {
	if cmd, ok := m.data[key]; ok {
		if cmd.expire.After(time.Now()) {
			return cmd.val, nil
		}
		m.lock.Lock()
		defer m.lock.Unlock()
		delete(m.data, key)
	}
	return nil, errors.New("not found")
}

func NewMemory() *memory {
	return &memory{
		data: make(map[string]*memoryCmd),
		lock: new(sync.Mutex),
	}
}
