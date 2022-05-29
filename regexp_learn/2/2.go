package main

import (
	"fmt"
	"regexp"
)

func main() {
	// var s string = "http://weibo.com/ttarticle/p/show?id=2309404620387190177901"
	// // resp, err := regexp.MatchString("=([0-9]*$)", s)
	// r, _ := regexp.Compile("([0-9]*$)")
	// resp := r.FindString(s)
	// // resp := r.ReplaceAllLiteralString(s, "${2}")
	// fmt.Print(resp)

	// var s string = "http://www12365auto.com/series/605/options.shtml"
	// r, _ := regexp.Compile("http://www.12365auto.com/series/([0-9]*)/options.shtml")
	// p_list := r.FindStringSubmatch(s)
	// if len(p_list) > 0 {
	// 	url := fmt.Sprintf("http://www.12365auto.com/js/series/param_%s.js", p_list[len(p_list) - 1])
	// 	fmt.Print(url)
	// }

	// url := "http://www.12365auto.com/pic/series-1-1416-0-9-2021-0-1.shtml"
	// r, _ := regexp.Compile(`series-\d+-\d+-0-\d+-\d+-\d+-\d+.shtml`)
	// p_list := r.FindStringSubmatch(url)
	// if len(p_list) <= 0 {
	// 	fmt.Printf("%v", p_list)
	// }

	// url := "国VIii"
	// g_compile, _ := regexp.Compile(`国[A-Z]{1,2}`)
	// b := bytes.NewReader([]byte(url))
	// p_list := g_compile.FindReaderIndex(b)
	// // fmt.Printf("%v\n", p_list)
	// fmt.Printf("%v\n", url[p_list[0]:p_list[1]])
	// fmt.Println(g_compile.ReplaceAllString(url, ""))

	str := "2021款 ko"
	year_compile, err := regexp.Compile(`\d{4}款`)
	if err != nil {
		fmt.Println(err)
	}

	indexs := year_compile.FindStringIndex(str)
	if len(indexs) != 0 {
		fmt.Println(str[indexs[0]:indexs[1]]) 
		fmt.Println(year_compile.ReplaceAllString(str, "")) 
	} else {
		fmt.Println(nil) 
	}

	if "ok" == "ok" {
		fmt.Println("ok")
	}
}
