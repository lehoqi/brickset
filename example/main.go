/**
 * @Package example
 * @Time: 2022/2/9 5:29 PM
 * @Author: wuhb
 * @File: main.go
 */

package main

import (
	"context"
	"github.com/wuhongbing/brickset"
	"log"
)

func main() {
	ctx := context.Background()
	svc := brickset.New("api-key", "username", "password", brickset.WithDebug(true))
	_, themes, err := svc.GetThemes(ctx)
	if err != nil {
		panic(err)
	}
	for _, theme := range themes {
		_, sets, err := svc.GetSets(ctx, &brickset.GetSetRequest{Theme: theme.Theme, PageSize: 500})
		if err != nil {
			panic(err)
		}
		for _, set := range sets {
			log.Println(set.Name)
		}
	}
}
