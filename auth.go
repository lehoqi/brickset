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
	userName string
	password string
	client   IClient
}

func (a auth) Login(ctx context.Context) (string, error) {
	resp := &CommonResponse{}
	req := url.Values{}
	req.Add("apiKey", a.apiKey)
	req.Add("userName", a.userName)
	req.Add("password", a.password)
	err := a.client.GetJSON(ctx, loginURL, req, resp)
	if err != nil {
		return "", err
	}
	if resp.IsSuccess() {
		return resp.Hash, nil
	}
	return "", resp.Error()
}

func (a *auth) CheckUserHash(ctx context.Context, hash string) (bool, error) {
	resp := &CommonResponse{}
	req := url.Values{}
	req.Add("apiKey", a.apiKey)
	req.Add("userHash", hash)
	err := a.client.GetJSON(ctx, checkUserHashURL, req, resp)
	if err != nil {
		return false, err
	}
	if resp.IsSuccess() {
		return true, nil
	}
	return false, resp.Error()
}

func NewAuth(apiKey, userName, password string, client IClient) IBrickAuth {
	return &auth{
		apiKey:   apiKey,
		userName: userName,
		password: password,
		client:   client,
	}
}
