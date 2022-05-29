package main

import "fmt"

type point struct {
	x     int
	y     int
	count int
	value rune
}

func LCS(s1, s2 string) (isRune, notIsRune []rune) {

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
					List[i1][i2].x = -1
					List[i1][i2].y = -1
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

	x := len(runeS1) - 1
	y := len(runeS2) - 1

	for ; x >= 0 && y >= 0 && List[x][y].count > 0; {
		isRune = append(isRune, List[x][y].value)
		item := List[x][y]
		x = item.x
		y = item.y
	}

	isRune = Reverse(isRune)
	i := 0
	j := 0
	for ; i < len(isRune); i++ {
		for ; j < len(runeS1) ;j++ {
			if isRune[i] != runeS1[j] {
				notIsRune = append(notIsRune, runeS1[j])
			} else {
				j++ 
				break
			}
		}
	}

	for ; j < len(runeS1) ;j++ {
		notIsRune = append(notIsRune, runeS1[j])
	}

	return
}

func testLCS() {
	isRune, notIsRune := LCS("1232020", "2020")
	fmt.Println(string(isRune), "---", string(notIsRune))
	isRune, notIsRune = LCS("2020", "1232020")
	fmt.Println(string(isRune), "---", string(notIsRune))
	isRune, notIsRune = LCS("20 20", "20 20")
	fmt.Println(string(isRune), "---", string(notIsRune))
	isRune, notIsRune = LCS("2020 1.5T 手动风尚型 5座", "2020 领尚型")
	fmt.Println(string(isRune), "---", string(notIsRune))
	isRune, notIsRune = LCS(" 2020 领尚型", " 2020 动风尚型 5座")
	fmt.Println(string(isRune), "---", string(notIsRune))
	isRune, notIsRune = LCS("林肯进口", "长安林肯")
	fmt.Println(string(isRune), "---", string(notIsRune))
	isRune, notIsRune = LCS("旗胜F1", "大柴神")
	fmt.Println(string(isRune), "---", string(notIsRune))
	isRune, notIsRune = LCS("2020款 55 TFSI e", "2020款 A5 ABT Sportline")
	fmt.Println(string(isRune), "---", string(notIsRune))
	isRune, notIsRune = LCS("2020款 A5 ABT Sportline", "2020款 55 TFSI e")
	fmt.Println(string(isRune), "---", string(notIsRune))
}
