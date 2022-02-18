/**
 * @Package example
 * @Time: 2022/2/9 5:29 PM
 * @Author: wuhb
 * @File: main.go
 */

package main

import (
	"context"
	"log"

	"github.com/lehoqi/brickset"
)

const (
	apiKey   = "your api key"
	username = "your username"
	password = "your password"
)

func main() {
	ctx := context.Background()
	svc := brickset.New(apiKey, username, password, brickset.WithImagePath("/tmp/brickset/", "/brickset/"))
	_, themes, err := svc.GetThemes(ctx)
	if err != nil {
		panic(err)
	}
	for _, theme := range themes {
		_, sets, err := svc.GetSets(ctx, &brickset.GetSetRequest{Theme: theme.Theme, PageSize: 1})
		if err != nil {
			panic(err)
		}
		for _, set := range sets {
			log.Printf("%+v", set)
		}
	}
}
