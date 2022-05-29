package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)


func main() {
	// brandList := getGB()

	// var carGroupList []CarGroup2

	// for _, brand := range brandList {
	// 	carList := getIBBMS(brand)
	// 	for i := 0; i < len(carList); i++ {
	// 		for ii := i + 1; ii < len(carList); ii++ {
	// 			carGroup, err := carToCarGroup2(carList[i], carList[ii])

	// 			if err != nil {
	// 			}

	// 			if carGroup.Year != 0 {
	// 				carGroupList = append(carGroupList, carGroup)
	// 			}

	// 		}
	// 	}

	// }

	// for i := 0; i < len(carGroupList)-1; {
	// 	j := i + 200
	// 	if j > len(carGroupList)-1 {
	// 		j = len(carGroupList) - 1
	// 	}
	// 	batch(carGroupList[i:j])
	// 	i = j
	// }


	// carList := getIBBMS("ABT")
	// for i := 0; i < len(carList); i++ {
	// 	for ii := i + 1; ii < len(carList); ii++ {
	// 		carGroup, err := carToCarGroup2(carList[i], carList[ii])

	// 		if err != nil {
	// 		}

	// 		if carGroup.Year != 0 {
	// 			fmt.Println(carGroup)
	// 			// carGroupList = append(carGroupList, carGroup)
	// 		}

	// 	}
	// }

	// testLCS()

	// // filebyt, _ := ioutil.ReadFile("car_brand_model_style.json")
	// // // filebyt, _ := ioutil.ReadFile("car_style.json")

	// // var carStyleList []carStyle

	// // _ = json.Unmarshal(filebyt, &carStyleList)

	// // for i := 0; i < len(carStyleList); i++ {
	// // 	var lcsRespList []lcsResp
	// // 	for j := i + 1; j < len(carStyleList); j++ {
	// // 		S1 := carStyleList[i].BRAND + " " + carStyleList[i].MODEL + " " + carStyleList[i].STYLE
	// // 		S2 := carStyleList[j].BRAND + " " + carStyleList[j].MODEL + " " + carStyleList[j].STYLE
	// // 		resp := lcsDB(S1, S2)
	// // 		lcsRespList = append(lcsRespList, resp)
	// // 	}
	// // 	data, _ := json.Marshal(lcsRespList)
	// // 	name := fmt.Sprintf("lcs_resp_%d.json", i)
	// // 	ioutil.WriteFile(name, data, 0644)
	// // }

	byt, err := ioutil.ReadFile("CILI.json")

	if err != nil {
		fmt.Printf(err.Error())
	}

	imgInfo := &[]ImgInfo{}

	json.Unmarshal(byt, imgInfo)
	fmt.Printf("%d", len(*imgInfo))
}
