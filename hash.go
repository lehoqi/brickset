/**
 * @Package brickset
 * @Time: 2022/2/8 11:17 PM
 * @Author: wuhb
 * @File: hash.go
 */

package brickset

import (
	"context"
	"sync"
	"time"
)

type sHash struct {
	auth        IBrickAuth
	hash        string
	lastUpdated time.Time
	expires     time.Duration
	lock        *sync.RWMutex
}

func (s *sHash) GetHash(ctx context.Context) (string, error) {
	if s.hash != "" && s.lastUpdated.Add(s.expires).After(time.Now()) {
		return s.hash, nil
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	h, err := s.auth.Login(ctx)
	if err != nil {
		return "", err
	}
	s.hash = h
	s.lastUpdated = time.Now()
	return s.hash, nil
}

func NewHash(auth IBrickAuth, expires time.Duration) IBrickHash {
	if expires == 0 {
		expires = time.Hour * 24
	}
	return &sHash{auth: auth, expires: expires, lock: new(sync.RWMutex)}
}
