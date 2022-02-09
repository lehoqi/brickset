/**
 * @Package brickset
 * @Time: 2022/2/8 9:27 PM
 * @Author: wuhb
 * @File: brickset.go
 */

package brickset

import (
	"context"
	"log"
	"net/url"
	"path"
	"time"
)

type IBrickAuth interface {
	Login(ctx context.Context) (string, error)
	CheckUserHash(ctx context.Context, hash string) (bool, error)
}

type IBrickHash interface {
	GetHash(ctx context.Context) (string, error)
}
type IBrickSet interface {
	GetSets(ctx context.Context, useHash bool, params *GetSetRequest) (int, []*Sets, error)
	GetThemes(ctx context.Context) (int, []*Themes, error)
}

var Logger log.Logger

type Option func(conf *config)
type brickSet struct {
	conf   *config
	hash   IBrickHash
	client IClient
}

func (b brickSet) GetThemes(ctx context.Context) (int, []*Themes, error) {
	response := &CommonResponse{}
	request := url.Values{}
	request.Set("apiKey", b.conf.apiKey)
	err := b.client.GetJSON(ctx, getThemesURL, request, response)
	if err != nil {
		return 0, nil, err
	}
	return response.Matches, response.Themes, response.Error()
}

func (b brickSet) GetSets(ctx context.Context, useHash bool, params *GetSetRequest) (int, []*Sets, error) {
	response := &CommonResponse{}
	request := url.Values{}
	request.Set("params", params.JSON())
	request.Set("apiKey", b.conf.apiKey)
	if useHash {
		hash, err := b.hash.GetHash(ctx)
		if err != nil {
			return 0, nil, err
		}
		request.Set("userHash", hash)
	}
	err := b.client.GetJSON(ctx, getSetsURL, request, response)
	if err != nil {
		return 0, nil, err
	}
	if b.conf.imagePath != nil {
		for _, set := range response.Sets {
			err := set.Image.Download(b.conf.imagePath)
			if err != nil {
				Logger.Printf("download image error:%v", err)
			}
		}
	}
	return response.Matches, response.Sets, response.Error()
}

func New(apiKey, username, password string, opts ...Option) IBrickSet {
	conf := &config{
		apiKey:      apiKey,
		hashExpires: time.Hour * 24,
		userName:    username,
		password:    password,
	}
	for _, opt := range opts {
		opt(conf)
	}
	c := NewClient(baseURL, conf.debug)
	return &brickSet{
		conf:   conf,
		hash:   NewHash(NewAuth(conf.apiKey, conf.userName, conf.password, c), conf.hashExpires),
		client: c,
	}
}

func WithAuth(userName, password string) Option {
	return func(conf *config) {
		conf.userName = userName
		conf.password = password
	}
}

func WithDebug(debug bool) Option {
	return func(conf *config) {
		conf.debug = debug
	}
}

func WithHashExpires(expires time.Duration) Option {
	return func(conf *config) {
		conf.hashExpires = expires
	}
}

func WithImagePath(basePath string, prefix string) Option {
	return func(conf *config) {
		if !path.IsAbs(basePath) {
			Logger.Fatal("path is not absolute")
			return
		}
		conf.imagePath = &imagePath{
			base:   basePath,
			prefix: prefix,
		}
	}
}
