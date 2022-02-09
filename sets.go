/**
 * @Package brickset
 * @Time: 2022/2/8 5:16 PM
 * @Author: wuhb
 * @File: sets.go
 */

package brickset

import (
	"encoding/json"
	"path"
	"time"
)

type GetSetRequest struct {
	SetID      int    `json:"setID,omitempty"`
	Query      string `json:"query,omitempty"`
	Theme      string `json:"theme,omitempty"`
	SetNumber  string `json:"setNumber,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	PageNumber int    `json:"pageNumber,omitempty"`
}

func (g GetSetRequest) JSON() string {
	s, _ := json.Marshal(g)
	return string(s)
}

type Sets struct {
	SetID                int          `json:"setID"`
	Number               string       `json:"number"`
	NumberVariant        int          `json:"numberVariant"`
	Name                 string       `json:"name"`
	Year                 int          `json:"year"`
	Theme                string       `json:"theme"`
	ThemeGroup           string       `json:"themeGroup"`
	Subtheme             string       `json:"subtheme"`
	Category             string       `json:"category"`
	Released             bool         `json:"released"`
	Pieces               int          `json:"pieces"`
	Minifigs             int          `json:"minifigs"`
	Image                Image        `json:"image"`
	BricksetURL          string       `json:"bricksetURL"`
	Collection           Collection   `json:"collection"`
	Collections          Collections  `json:"collections"`
	LEGOCom              LEGOCom      `json:"LEGOCom"`
	Rating               float32      `json:"rating"`
	ReviewCount          int          `json:"reviewCount"`
	PackagingType        string       `json:"packagingType"`
	Availability         string       `json:"availability"`
	InstructionsCount    int          `json:"instructionsCount"`
	AdditionalImageCount int          `json:"additionalImageCount"`
	AgeRange             AgeRange     `json:"ageRange"`
	Dimensions           Dimensions   `json:"dimensions"`
	Barcode              Barcodes     `json:"barcode"`
	ExtendedData         ExtendedData `json:"extendedData"`
	LastUpdated          string       `json:"lastUpdated"`
}

type LEGOCom struct {
	US LEGOComDetails `json:"US"`
	UK LEGOComDetails `json:"UK"`
	CA LEGOComDetails `json:"CA"`
	DE LEGOComDetails `json:"DE"`
}

type LEGOComDetails struct {
	RetailPrice        float32   `json:"retailPrice"`
	DateFirstAvailable time.Time `json:"dateFirstAvailable"`
	DateLastAvailable  time.Time `json:"dateLastAvailable"`
}

type Dimensions struct {
	Height float32 `json:"height"`
	Width  float32 `json:"width"`
	Depth  float32 `json:"depth"`
	Weight float32 `json:"weight"`
}

type ExtendedData struct {
	Notes       string   `json:"notes"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
}

type Image struct {
	ThumbnailURL string `json:"thumbnailURL"`
	ImageURL     string `json:"imageURL"`
}

func (i *Image) Download(c *imagePath) error {
	tpPath, err := getPathFromURL(i.ThumbnailURL)
	if err != nil {
		return err
	}
	err = downloadFile(i.ThumbnailURL, path.Join(c.base, tpPath))
	if err != nil {
		return err
	}
	i.ThumbnailURL = c.prefix + tpPath
	ipath, err := getPathFromURL(i.ImageURL)
	if err != nil {
		return err
	}
	err = downloadFile(i.ImageURL, path.Join(c.base, ipath))
	if err != nil {
		return err
	}
	i.ImageURL = c.prefix + ipath
	return nil
}

type Collections struct {
	OwnedBy  int `json:"ownedBy"`
	WantedBy int `json:"wantedBy"`
}

type Collection struct {
	Owned    bool   `json:"owned"`
	Wanted   bool   `json:"wanted"`
	QtyOwned int    `json:"qtyOwned"`
	Rating   int    `json:"rating"`
	Notes    string `json:"notes"`
}

type AgeRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Barcodes struct {
	EAN string `json:"EAN"`
	UPC string `json:"UPC"`
}
