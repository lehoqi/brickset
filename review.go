/**
 * @Package brickset
 * @Time: 2022/2/9 10:19 PM
 * @Author: wuhb
 * @File: review.go
 */

package brickset

type Review struct {
	Author     string  `json:"author"`
	DatePosted string  `json:"datePosted"`
	Rating     *Rating `json:"rating"`
	Title      string  `json:"title"`
	Review     string  `json:"review"`
	HTML       bool    `json:"HTML"`
}

type Rating struct {
	Overall            int `json:"overall"`
	Parts              int `json:"parts"`
	BuildingExperience int `json:"buildingExperience"`
	Playability        int `json:"playability"`
	ValueForMoney      int `json:"valueForMoney"`
}
