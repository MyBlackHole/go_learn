package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	jsonFile, err := os.Open("user.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Println(string(byteValue))

	//var users Users

	//json.Unmarshal(byteValue, &users)

	//for i := 0; i < len(users.Users); i++ {
	//      fmt.Println("User Type: " + users.Users[i].Type)
	//      fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
	//      fmt.Println("User Name: " + users.Users[i].Name)
	//      fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	//}

	// var result map[string]interface{}
	var result Users

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	json.Unmarshal(byteValue, &result)

	fmt.Println(result)

}
