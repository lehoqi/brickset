/**
 * @Package brickset
 * @Time: 2022/2/8 5:06 PM
 * @Author: wuhb
 * @File: config.go
 */

package brickset

import "time"

var defaultHashExpires = time.Hour * 24

type config struct {
	debug       bool
	username    string
	password    string
	apiKey      string
	hashExpires time.Duration
	imagePath   *imagePath
	storage     IBrickStorage
}
type imagePath struct {
	base   string
	prefix string
}
