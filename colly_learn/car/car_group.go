package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	"strconv"
	// "time"
	// "fmt"
	// "io/ioutil"
)

func carToCarGroup(car1, car2 Car) (carGroup CarGroup, err error) {
	if car1.Source == car2.Source {
		err = errors.New("站点相同不比较")
		return
	}

	carGroup.LabelID = 0
	carGroup.BrandPro = car1.BrandPro
	carGroup.AID = car1.ID
	carGroup.BID = car2.ID
	if car1.ID > car2.ID {
		carGroup.Path = 1
	}


	car1Str := car1.BrandPro + " " + car1.Brand + " " + car1.Model + " " + car1.Style
	year1Str, car1Str := getYear(car1Str)

	car2Str := car2.BrandPro + " " + car2.Brand + " " + car2.Model + " " + car2.Style
	year2Str, car2Str := getYear(car2Str)

	if year1Str != year2Str || len(year1Str) == 0 {
		err = errors.New("年份不相同或没有年份")
		return
	}

	yearInt, err := strconv.Atoi(year1Str[0:4])
	if err != nil {
		fmt.Println(err)
	}

	carGroup.Year = yearInt

	isRune, notIsRune := LCS(car1Str, car2Str)
	carGroup.EqualStr = string(isRune)
	carGroup.EqualCount = len(isRune)
	carGroup.NotEqualStr = string(notIsRune)
	carGroup.NotEqualCount = len(notIsRune)
	return
}

func carToCarGroup2(car1, car2 Car) (carGroup CarGroup2, err error) {

	if car1.Source == car2.Source {
		err = errors.New("站点相同不比较")
		return
	}

	carGroup.BrandPro = car1.BrandPro
	carGroup.AID = car1.ID
	carGroup.BID = car2.ID

	var car1Year string
	var car2Year string
	car1Year, car1.Style = getYear(car1.Style)

	car2Year, car2.Style = getYear(car2.Style)

	if car1Year != car2Year || len(car2Year) == 0 {
		err = errors.New("年份不相同或没有年份")
		return
	}

	// 年份 string 转 int
	yearInt, err := strconv.Atoi(car1Year[0:4])
	if err != nil {
		fmt.Println(err)
	}
	carGroup.Year = yearInt

	// brand a - b
	isRune, abNotIsRune := LCS(car1.Brand, car2.Brand)
	// brand b - a
	isRune, baNotIsRune := LCS(car2.Brand, car1.Brand)
	carGroup.BrandEqualStr = string(isRune)
	carGroup.BrandEqualCount = len(isRune)
	carGroup.AbBrandNes = string(abNotIsRune)
	carGroup.AbBrandNec = len(abNotIsRune)
	carGroup.BaBrandNes = string(baNotIsRune)
	carGroup.BaBrandNec = len(baNotIsRune)
	// flag
	abBrandFlag := 0
	abNotIs := string(abNotIsRune)
	if isSz(abNotIs) {
		abBrandFlag |= 1 << 0
	}
	if isZm(abNotIs) {
		abBrandFlag |= 1 << 1
	}
	if isZf(abNotIs) {
		abBrandFlag |= 1 << 2
	}
	if isZw(abNotIs) {
		abBrandFlag |= 1 << 3
	}
	carGroup.AbBrandFlag = abBrandFlag

	baBrandFlag := 0
	baNotIs := string(baNotIsRune)
	if isSz(baNotIs) {
		baBrandFlag |= 1 << 0
	}
	if isZm(baNotIs) {
		baBrandFlag |= 1 << 1
	}
	if isZf(baNotIs) {
		baBrandFlag |= 1 << 2
	}
	if isZw(baNotIs) {
		baBrandFlag |= 1 << 3
	}
	carGroup.BaBrandFlag = baBrandFlag

	// model a - b
	isRune, abNotIsRune = LCS(car1.Model, car2.Model)
	// model b - a
	isRune, baNotIsRune = LCS(car2.Model, car1.Model)
	carGroup.ModelEqualStr = string(isRune)
	carGroup.ModelEqualCount = len(isRune)
	carGroup.AbModelNes = string(abNotIsRune)
	carGroup.AbModelNec = len(abNotIsRune)
	carGroup.BaModelNes = string(baNotIsRune)
	carGroup.BaModelNec = len(baNotIsRune)
	// flag
	abModelFlag := 0
	abNotIs = string(abNotIsRune)
	if isSz(abNotIs) {
		abModelFlag |= 1 << 0
	}
	if isZm(abNotIs) {
		abModelFlag |= 1 << 1
	}
	if isZf(abNotIs) {
		abModelFlag |= 1 << 2
	}
	if isZw(abNotIs) {
		abModelFlag |= 1 << 3
	}
	carGroup.AbModelFlag = abModelFlag

	baModelFlag := 0
	baNotIs = string(baNotIsRune)
	if isSz(baNotIs) {
		baModelFlag |= 1 << 0
	}
	if isZm(baNotIs) {
		baModelFlag |= 1 << 1
	}
	if isZf(baNotIs) {
		baModelFlag |= 1 << 2
	}
	if isZw(baNotIs) {
		baModelFlag |= 1 << 3
	}
	carGroup.BaModelFlag = baModelFlag

	// style a - b
	isRune, abNotIsRune = LCS(car1.Style, car2.Style)
	// style b - a
	isRune, baNotIsRune = LCS(car2.Style, car1.Style)
	carGroup.StyleEqualStr = string(isRune)
	carGroup.StyleEqualCount = len(isRune)
	carGroup.AbStyleNes = string(abNotIsRune)
	carGroup.AbStyleNec = len(abNotIsRune)
	carGroup.BaStyleNes = string(baNotIsRune)
	carGroup.BaStyleNec = len(baNotIsRune)
	// flag
	abStyleFlag := 0
	abNotIs = string(abNotIsRune)
	if isSz(abNotIs) {
		abStyleFlag |= 1 << 0
	}
	if isZm(abNotIs) {
		abStyleFlag |= 1 << 1
	}
	if isZf(abNotIs) {
		abStyleFlag |= 1 << 2
	}
	if isZw(abNotIs) {
		abStyleFlag |= 1 << 3
	}
	carGroup.AbStyleFlag = abStyleFlag

	baStyleFlag := 0
	baNotIs = string(baNotIsRune)
	if isSz(baNotIs) {
		baStyleFlag |= 1 << 0
	}
	if isZm(baNotIs) {
		baStyleFlag |= 1 << 1
	}
	if isZf(baNotIs) {
		baStyleFlag |= 1 << 2
	}
	if isZw(baNotIs) {
		baStyleFlag |= 1 << 3
	}
	carGroup.BaStyleFlag = baStyleFlag

	return
}
