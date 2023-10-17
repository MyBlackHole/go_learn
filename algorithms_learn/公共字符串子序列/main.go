package main

import (
	// "encoding/json"
	"fmt"
	// "time"
	// "fmt"
	// "io/ioutil"
)

type point struct {
	x     int
	y     int
	count int
	value rune
}

func LCS(s1, s2 string) (isRune []rune, notIsRune []rune) {

	if len(s1) == 0 || len(s2) == 0 {
		notIsRune = []rune(s1)
		return
	}

	runeS1 := []rune(s1)
	runeS2 := []rune(s2)

	var List = make([][]point, len(runeS1))

	for i := range List {
		List[i] = make([]point, len(runeS2))
	}

	for i1, value1 := range runeS1 {
		for i2, value2 := range runeS2 {
			index_i1 := i1 - 1
			index_i2 := i2 - 1
			if value1 == value2 {
				switch {
				case index_i1 < 0 || index_i2 < 0:
					List[i1][i2].count = 1
					List[i1][i2].value = value1
				default:
					List[i1][i2].count = List[index_i1][index_i2].count + 1
					List[i1][i2].x = index_i1
					List[i1][i2].y = index_i2
					List[i1][i2].value = value1
				}
			} else {
				switch {
				case index_i1 < 0 && index_i2 < 0:
				case index_i1 < 0:
					List[i1][i2].count = List[0][index_i2].count
					List[i1][i2].x = List[0][index_i2].x
					List[i1][i2].y = List[0][index_i2].y
					List[i1][i2].value = List[0][index_i2].value
				case index_i2 < 0:
					List[i1][i2].count = List[index_i1][0].count
					List[i1][i2].x = List[index_i1][0].x
					List[i1][i2].y = List[index_i1][0].y
					List[i1][i2].value = List[index_i1][0].value
				default:
					if List[i1][index_i2].count > List[index_i1][i2].count {
						List[i1][i2].count = List[i1][index_i2].count
						List[i1][i2].x = List[i1][index_i2].x
						List[i1][i2].y = List[i1][index_i2].y
						List[i1][i2].value = List[i1][index_i2].value
					} else {
						List[i1][i2].count = List[index_i1][i2].count
						List[i1][i2].x = List[index_i1][i2].x
						List[i1][i2].y = List[index_i1][i2].y
						List[i1][i2].value = List[index_i1][i2].value
					}
				}
			}
		}
	}
	var ii int
	i := len(runeS1) - 1
	j := len(runeS2) - 1

	for List[i][j].count > 0 {
		isRune = append(isRune, List[i][j].value)
		ii = List[i][j].x + 1
		j = List[i][j].y

		for ; i != ii; i-- {
			notIsRune = append(notIsRune, runeS1[i])
		}
		i--
	}

	for ; i >= 0; i-- {
		notIsRune = append(notIsRune, runeS1[i])
	}

	return
}

func testLCS() {
	isRune, notIsRune := LCS("fdop", "bac")
	fmt.Print(string(isRune), string(notIsRune))
}

type CarStyle struct {
	STYLE string
	BRAND string
	MODEL string
}

type LcsResp struct {
	S1    string
	S2    string
	COUNT int
}

func main() {
	testLCS()

	// filebyt, _ := ioutil.ReadFile("car_brand_model_style.json")
	// // filebyt, _ := ioutil.ReadFile("car_style.json")

	// var carStyleList []carStyle

	// _ = json.Unmarshal(filebyt, &carStyleList)

	// for i := 0; i < len(carStyleList); i++ {
	// 	var lcsRespList []lcsResp
	// 	for j := i + 1; j < len(carStyleList); j++ {
	// 		S1 := carStyleList[i].BRAND + " " + carStyleList[i].MODEL + " " + carStyleList[i].STYLE
	// 		S2 := carStyleList[j].BRAND + " " + carStyleList[j].MODEL + " " + carStyleList[j].STYLE
	// 		resp := lcsDB(S1, S2)
	// 		lcsRespList = append(lcsRespList, resp)
	// 	}
	// 	data, _ := json.Marshal(lcsRespList)
	// 	name := fmt.Sprintf("lcs_resp_%d.json", i)
	// 	ioutil.WriteFile(name, data, 0644)
	// }
}
