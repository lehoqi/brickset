/**
 * @Package brickset
 * @Time: 2022/2/9 9:40 AM
 * @Author: wuhb
 * @File: tool.go
 */

package brickset

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func getPathFromURL(u string) (string, error) {
	s, err := url.ParseRequestURI(u)
	if err != nil {
		return "", err
	}
	return s.Path, nil
}

func downloadFile(source string, target string) error {
	resp, err := http.Get(source)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(target)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err
}