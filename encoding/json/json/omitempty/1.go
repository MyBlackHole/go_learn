package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Id string `json:"id,omitempty"`
	A  string `json:"a,omitempty"`
}

type Org struct {
	Name string `json:"name"`
}

type AppWithOrg struct {
	App *App    `json:"app,omitempty"`
	Org Org     `json:"org,omitempty"`
	S   *string `json:"s"`
	R   *string `json:"r,omitempty"`
	E   string  `json:"e"`
	D   string  `json:"d,omitempty"`
}

type AppList struct {
	applist []App
}

func main() {
	data := []byte(`
        {
            "id": "k34rAT4",
            "name": "My Awesome Org"
        }
    `)

	var b AppWithOrg

	json.Unmarshal(data, &b)
	fmt.Printf("%#v\n", b)

	// a := AppWithOrg{
	// 	App: App{
	// 		Id: "k34rAT4",
	// 	},
	// 	Org: Org{
	// 		Name: "My Awesome Org",
	// 	},
	// }

	var a AppWithOrg
	a.App = &App{Id: "sdjf"}

	data, _ = json.Marshal(a)
	fmt.Println(string(data))

	appList := AppList{}
	data, _ = json.Marshal(appList)
	fmt.Println(string(data))
}
