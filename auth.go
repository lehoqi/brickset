/**
 * @Package brickset
 * @Time: 2022/2/8 5:35 PM
 * @Author: wuhb
 * @File: auth.go
 */

package brickset

import (
	"context"
	"net/url"
)

type auth struct {
	apiKey   string
	username string
	password string
	client   IClient
}

func (a auth) CheckKey(ctx context.Context, key string) (bool, error) {
	if key == "" {
		key = a.apiKey
	}
	resp := &CommonResponse{}
	req := url.Values{}
	req.Set("apiKey", key)
	err := a.client.GetJSON(ctx, checkKeyURL, req, resp)
	if err != nil {
		return false, err
	}
	return resp.IsSuccess(), resp.Error()
}

func (a auth) Login(ctx context.Context) (string, error) {
	response := &CommonResponse{}
	req := url.Values{}
	req.Add("apiKey", a.apiKey)
	req.Add("username", a.username)
	req.Add("password", a.password)
	err := a.client.GetJSON(ctx, loginURL, req, response)
	if err != nil {
		return "", err
	}

	if response.IsSuccess() {
		return response.Hash, nil
	}

	return "", response.Error()
}

func (a *auth) CheckUserHash(ctx context.Context, hash string) (bool, error) {
	response := &CommonResponse{}
	req := url.Values{}
	req.Add("apiKey", a.apiKey)
	req.Add("userHash", hash)
	err := a.client.GetJSON(ctx, checkUserHashURL, req, response)
	if err != nil {
		return false, err
	}

	if response.IsSuccess() {
		return true, nil
	}

	return false, response.Error()
}

func NewAuth(apiKey, username, password string, client IClient) IBrickAuth {
	return &auth{
		apiKey:   apiKey,
		username: username,
		password: password,
		client:   client,
	}
}
