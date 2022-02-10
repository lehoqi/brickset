/**
 * @Package brickset
 * @Time: 2022/2/8 11:17 PM
 * @Author: wuhb
 * @File: hash.go
 */

package brickset

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type sHash struct {
	auth    IBrickAuth
	storage IBrickStorage
	expires time.Duration
	lock    *sync.RWMutex
}

func (s *sHash) GetHash(ctx context.Context, username string) (string, error) {
	v, err := s.storage.Get(ctx, fmt.Sprintf("%s-hash", username))
	if err == nil {
		return v.(string), nil
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	h, err := s.auth.Login(ctx)
	if err != nil {
		return "", err
	}
	if err = s.storage.Set(ctx, fmt.Sprintf("%s-hash", username), h, s.expires); err != nil {
		return "", err
	}
	return h, nil
}

func NewHash(auth IBrickAuth, store IBrickStorage, expires time.Duration) IBrickHash {
	if expires <= 0 {
		expires = time.Hour
	}
	return &sHash{
		auth:    auth,
		storage: store,
		expires: expires,
		lock:    new(sync.RWMutex)}
}
