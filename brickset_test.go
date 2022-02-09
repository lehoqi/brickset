/**
 * @Package brickset
 * @Time: 2022/2/8 11:43 PM
 * @Author: wuhb
 * @File: brickset_test.go
 */

package brickset

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	apiKey   = ""
	userName = ""
	password = ""
)

func TestGetSetsByTheme(t *testing.T) {
	b := New(apiKey, userName, password, WithDebug(true))
	_, sets, err := b.GetSets(context.Background(), &GetSetRequest{Theme: "Pirates"})
	assert.NoError(t, err)
	assert.NotEmpty(t, sets)
}
func TestGetThemes(t *testing.T) {
	b := New(apiKey, userName, password, WithDebug(true))
	_, themes, err := b.GetThemes(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, themes)
}

func TestGetSetsWithDownload(t *testing.T) {
	b := New(apiKey, userName, password, WithDebug(true), WithImagePath("/images", "p"))
	_, sets, err := b.GetSets(context.Background(), &GetSetRequest{Theme: "Pirates"})
	assert.NoError(t, err)
	assert.NotEmpty(t, sets)
}

func TestAuthCheckKey(t *testing.T) {
	a := NewAuth(apiKey, userName, password, NewClient(baseURL, true))
	ok, err := a.CheckKey(context.Background(), "")
	assert.NoError(t, err)
	assert.True(t, ok)
}
