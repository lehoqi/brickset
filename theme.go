/**
 * @Package brickset
 * @Time: 2022/2/8 5:09 PM
 * @Author: wuhb
 * @File: theme.go
 */

package brickset

type Themes struct {
	Theme         string
	SetCount      int
	SubthemeCount int
	YearFrom      int
	YearTo        int
}

type Subthemes struct {
	Theme    string
	Subtheme string
	SetCount int
	YearFrom int
	YearTo   int
}
