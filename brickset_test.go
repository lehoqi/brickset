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
	apiKey   = "your_api_key"
	username = "your_username"
	password = "your_password"
)
var bs = New(apiKey, username, password)

func TestGetSetsByTheme(t *testing.T) {
	_, sets, err := bs.GetSets(context.Background(), &GetSetRequest{Theme: "Pirates", PageSize: 1})
	assert.NoError(t, err)
	assert.NotEmpty(t, sets)
}
func TestGetThemes(t *testing.T) {
	_, themes, err := bs.GetThemes(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, themes)
}

func TestAuthCheckKey(t *testing.T) {
	a := NewAuth(apiKey, username, password, NewClient(baseURL, true))
	ok, err := a.CheckKey(context.Background(), "")
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestGetReviews(t *testing.T) {
	_, reviews, err := bs.GetReviews(context.Background(), 7308)
	assert.NoError(t, err)
	assert.NotEmpty(t, reviews)
}

func TestBrickSet_GetAdditionalImages(t *testing.T) {
	_, images, err := bs.GetAdditionalImages(context.Background(), 7308)
	assert.NoError(t, err)
	assert.NotEmpty(t, images)
}

func TestBrickSet_GetInstructions(t *testing.T) {
	_, instructions, err := bs.GetInstructions(context.Background(), 7308)
	assert.NoError(t, err)
	assert.NotEmpty(t, instructions)
}

func TestBrickSet_GetSubthemes(t *testing.T) {
	_, subthemes, err := bs.GetSubthemes(context.Background(), "Action Wheelers")
	assert.NoError(t, err)
	assert.NotEmpty(t, subthemes)
}

func TestBrickSet_GetYears(t *testing.T) {
	_, years, err := bs.GetYears(context.Background(), "Action Wheelers")
	assert.NoError(t, err)
	assert.NotEmpty(t, years)
}
