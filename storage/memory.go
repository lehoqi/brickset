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

type Memory struct {
	data map[string]*memoryCmd
	lock sync.Locker
}

func (m Memory) Set(_ context.Context, key string, value interface{}, expiration time.Duration) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data[key] = &memoryCmd{
		val:    value,
		expire: time.Now().Add(expiration),
	}
	return nil
}

func (m Memory) Get(_ context.Context, key string) (interface{}, error) {
	if cmd, ok := m.data[key]; ok {
		if cmd.expire.After(time.Now()) {
			return cmd.val, nil
		}
		delete(m.data, key)
	}

	return nil, errors.New("not found")
}

func NewMemory() *Memory {
	return &Memory{
		data: make(map[string]*memoryCmd),
		lock: new(sync.Mutex),
	}
}
