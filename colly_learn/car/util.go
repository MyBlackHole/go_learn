package main

import (
	"fmt"
	"regexp"
)

func getYear(str string) (yearStr, src string) {
	year_compile, err := regexp.Compile(`\d{4}款`)
	if err != nil {
		fmt.Println(err)
	}

	indexs := year_compile.FindStringIndex(str)
	if len(indexs) != 0 {
		yearStr = str[indexs[0] : indexs[0]+4]
		// src = year_compile.ReplaceAllString(str, "")
		src = str
	} else {
		src = str
	}
	return
}

func isZw(str string) (b bool) {
	zm_compile, err := regexp.Compile("[\u2E80-\u9FFF]")
	if err != nil {
		fmt.Println(err)
	}

	indexs := zm_compile.FindStringIndex(str)
	if len(indexs) != 0 {
		b = true
	}
	return
}

func isZf(str string) (b bool) {
	zm_compile, err := regexp.Compile(`[~!@#$%^&*()_\-+=<>?:"{}|,.\/;'\\[\]·~！@#￥%……&*（）——\-+={}|《》？：“”【】、；‘'，。、]`)
	if err != nil {
		fmt.Println(err)
	}

	indexs := zm_compile.FindStringIndex(str)
	if len(indexs) != 0 {
		b = true
	}
	return
}

func isZm(str string) (b bool) {
	zm_compile, err := regexp.Compile(`[A-Za-z]`)
	if err != nil {
		fmt.Println(err)
	}

	indexs := zm_compile.FindStringIndex(str)
	if len(indexs) != 0 {
		b = true
	}
	return
}

func isSz(str string) (b bool) {
	sz_compile, err := regexp.Compile(`\d`)
	if err != nil {
		fmt.Println(err)
	}

	indexs := sz_compile.FindStringIndex(str)
	if len(indexs) != 0 {
		b = true
	}
	return
}

func getG(str string) (src string) {
	g_compile, err := regexp.Compile(`国[A-Z]{1,2}`)
	if err != nil {
		fmt.Println(err)
	}

	indexs := g_compile.FindStringIndex(str)
	if len(indexs) != 0 {
		src = g_compile.ReplaceAllString(str, "")
	} else {
		src = str
	}
	return
}

func getYG(str string) (yearStr, src string) {
	yearStr, src = getYear(str)
	src = getG(src)
	return
}

func Reverse(r []rune) []rune {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func textFunc() {
	fmt.Println(isSz("1232"))
	fmt.Println(isZf(".adf"))
	fmt.Println(isZm("ad.."))
	fmt.Println(isZw("快速的佛教.."))
}
