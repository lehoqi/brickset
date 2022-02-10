/**
 * @Package brickset
 * @Time: 2022/2/8 9:27 PM
 * @Author: wuhb
 * @File: brickset.go
 */

package brickset

import (
	"context"
	"github.com/wuhongbing/brickset/storage"
	"log"
	"net/url"
	"path"
	"strconv"
	"time"
)

type IBrickAuth interface {
	Login(ctx context.Context) (string, error)
	CheckUserHash(ctx context.Context, hash string) (bool, error)
	CheckKey(ctx context.Context, key string) (bool, error)
}

type IBrickHash interface {
	GetHash(ctx context.Context, username string) (string, error)
}
type IBrickSet interface {
	GetSets(ctx context.Context, params *GetSetRequest) (int, []*Sets, error)
	GetThemes(ctx context.Context) (int, []*Themes, error)
	GetReviews(ctx context.Context, setID int) (int, []*Review, error)
	GetSubthemes(ctx context.Context, theme string) (int, []*Subthemes, error)
	GetInstructions(ctx context.Context, setID int) (int, []*Instruction, error)
	GetInstructions2(ctx context.Context, setNumber string) (int, []*Instruction, error)
	GetAdditionalImages(ctx context.Context, setID int) (int, []*Image, error)
	GetYears(ctx context.Context, theme string) (int, []*Years, error)
}

type IBrickStorage interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
}

var Logger = log.Default()

type Option func(conf *config)
type brickSet struct {
	conf   *config
	hash   IBrickHash
	auth   IBrickAuth
	client IClient
}

func (b brickSet) GetYears(ctx context.Context, theme string) (int, []*Years, error) {
	response := &CommonResponse{}
	params := url.Values{}
	params.Set("apiKey", b.conf.apiKey)
	params.Set("theme", theme)
	err := b.client.GetJSON(ctx, getYearsURL, params, response)
	if err != nil {
		return 0, nil, err
	}
	return response.Matches, response.Years, response.Error()
}

func (b brickSet) GetAdditionalImages(ctx context.Context, setID int) (int, []*Image, error) {
	response := &CommonResponse{}
	params := url.Values{}
	params.Set("apiKey", b.conf.apiKey)
	params.Set("setID", strconv.Itoa(setID))
	err := b.client.GetJSON(ctx, getAdditionalImagesURL, params, response)
	if err != nil {
		return 0, nil, err
	}

	if b.conf.imagePath != nil && response.AdditionalImages != nil {
		for _, image := range response.AdditionalImages {
			err := image.Download(b.conf.imagePath)
			if err != nil {
				Logger.Printf("download image error:%v", err)
			}
		}
	}
	return response.Matches, response.AdditionalImages, response.Error()
}
func (b brickSet) GetInstructions(ctx context.Context, setID int) (int, []*Instruction, error) {
	response := &CommonResponse{}
	params := url.Values{}
	params.Set("apiKey", b.conf.apiKey)
	params.Set("setID", strconv.Itoa(setID))
	err := b.client.GetJSON(ctx, getInstructionsURL, params, response)
	if err != nil {
		return 0, nil, err
	}
	return response.Matches, response.Instructions, response.Error()
}
func (b brickSet) GetInstructions2(ctx context.Context, setNumber string) (int, []*Instruction, error) {
	response := &CommonResponse{}
	params := url.Values{}
	params.Set("apiKey", b.conf.apiKey)
	params.Set("setNumber", setNumber)
	err := b.client.GetJSON(ctx, getInstructions2URL, params, response)
	if err != nil {
		return 0, nil, err
	}
	return response.Matches, response.Instructions, response.Error()
}

func (b brickSet) GetSubthemes(ctx context.Context, theme string) (int, []*Subthemes, error) {
	response := &CommonResponse{}
	params := url.Values{}
	params.Add("apiKey", b.conf.apiKey)
	params.Set("theme", theme)
	err := b.client.GetJSON(ctx, getSubthemesURL, params, response)
	if err != nil {
		return 0, nil, err
	}
	return response.Matches, response.Subthemes, response.Error()
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

func (b brickSet) GetReviews(ctx context.Context, setID int) (int, []*Review, error) {
	response := &CommonResponse{}
	request := url.Values{}
	request.Set("apiKey", b.conf.apiKey)
	request.Set("setID", strconv.Itoa(setID))
	err := b.client.GetJSON(ctx, getReviewsURL, request, response)
	if err != nil {
		return 0, nil, err
	}
	return response.Matches, response.Reviews, response.Error()
}

func (b brickSet) GetSets(ctx context.Context, params *GetSetRequest) (int, []*Sets, error) {
	response := &CommonResponse{}
	request := url.Values{}
	request.Set("params", params.JSON())
	request.Set("apiKey", b.conf.apiKey)
	hash, err := b.hash.GetHash(ctx, b.conf.username)
	if err != nil {
		return 0, nil, err
	}

	request.Set("userHash", hash)
	err = b.client.GetJSON(ctx, getSetsURL, request, response)
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
		hashExpires: defaultHashExpires,
		username:    username,
		password:    password,
	}
	for _, opt := range opts {
		opt(conf)
	}

	if conf.storage == nil {
		conf.storage = storage.NewMemory()
	}

	c := NewClient(baseURL, conf.debug)
	a := NewAuth(conf.apiKey, conf.username, conf.password, c)
	h := NewHash(a, conf.storage, conf.hashExpires)
	return &brickSet{
		conf:   conf,
		hash:   h,
		auth:   a,
		client: c,
	}
}

func WithAuth(username, password string) Option {
	return func(conf *config) {
		conf.username = username
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

func WithStorage(s IBrickStorage) Option {
	return func(conf *config) {
		conf.storage = s
	}
}
