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
	"path/filepath"
)

func getPathFromURL(u string) (string, error) {
	s, err := url.ParseRequestURI(u)
	if err != nil {
		return "", err
	}
	return s.Path, nil
}

func downloadFile(url string, target string) error {
	_ = os.MkdirAll(filepath.Dir(target), os.ModePerm)
	resp, err := http.Get(url)
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
