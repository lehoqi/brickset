/**
 * @Package brickset
 * @Time: 2022/2/8 9:47 PM
 * @Author: wuhb
 * @File: response.go
 */

package brickset

import "errors"

type CommonResponse struct {
	Status           string         `json:"status"`
	Message          string         `json:"message"`
	Themes           []*Themes      `json:"themes"`
	Matches          int            `json:"matches"`
	Sets             []*Sets        `json:"sets"`
	Hash             string         `json:"hash"`
	Reviews          []*Review      `json:"reviews"`
	Subthemes        []*Subthemes   `json:"subthemes"`
	Instructions     []*Instruction `json:"instructions"`
	AdditionalImages []*Image       `json:"additionalImages"`
	Years            []*Years
}

func (c CommonResponse) IsSuccess() bool {
	return c.Status == "success"
}
func (c CommonResponse) Error() error {
	if c.IsSuccess() {
		return nil
	}
	return errors.New(c.Message)
}
