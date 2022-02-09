/**
 * @Package brickset
 * @Time: 2022/2/8 10:35 PM
 * @Author: wuhb
 * @File: client.go
 */

package brickset

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
)

type IClient interface {
	GetJSON(ctx context.Context, api string, params url.Values, response interface{}) error
	PostJSON(ctx context.Context, api string, request interface{}, response interface{}) error
	PostForm(ctx context.Context, api string, params url.Values, response interface{}) error
}

type client struct {
	r *resty.Client
}

func (c client) PostForm(ctx context.Context, api string, params url.Values, response interface{}) error {
	resp, err := c.r.R().SetContext(ctx).SetFormDataFromValues(params).Post(api)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(http.StatusText(resp.StatusCode()))
	}
	return c.r.JSONUnmarshal(resp.Body(), response)
}

func (c client) GetJSON(ctx context.Context, api string, params url.Values, response interface{}) error {
	resp, err := c.r.R().SetContext(ctx).SetQueryParamsFromValues(params).Get(api)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(http.StatusText(resp.StatusCode()))
	}
	return c.r.JSONUnmarshal(resp.Body(), response)
}

func (c client) PostJSON(ctx context.Context, api string, request interface{}, response interface{}) error {
	resp, err := c.r.R().SetContext(ctx).SetBody(request).Post(api)
	if err != nil {
		return err
	}
	return c.r.JSONUnmarshal(resp.Body(), response)
}

func NewClient(baseUrl string, debug bool) IClient {
	c := resty.New()
	c = c.SetBaseURL(baseUrl).SetDebug(debug).SetHeader("User-Agent", "Brickset-Go-Client")
	return &client{r: c}
}
