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
	apiKey   = "your own api key"
	userName = "your own username"
	password = "your own password"
)

func TestGetSetsByTheme(t *testing.T) {
	b := New(apiKey, userName, password, WithDebug(true))
	_, sets, err := b.GetSets(context.Background(), true, &GetSetRequest{Theme: "Pirates"})
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
	b := New(apiKey, userName, password, WithDebug(true), WithImagePath("/Users/wuhb/Downloads/brickset/images", "wuhb"))
	_, sets, err := b.GetSets(context.Background(), true, &GetSetRequest{Theme: "Pirates"})
	assert.NoError(t, err)
	assert.NotEmpty(t, sets)
}
