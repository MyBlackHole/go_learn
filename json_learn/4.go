package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Family struct {
    LastName string
}
type Location struct {
    City string
}
type Person struct {
    Family    `mapstructure:",squash"`
    Location  `mapstructure:",squash"`
    FirstName string
}

func main() {
    input := map[string]interface{}{
        "FirstName": "Mitchell",
        "LastName":  "Hashimoto",
        "City":      "San Francisco",
    }

    var result Person
    err := mapstructure.Decode(input, &result)
    if err != nil {
        panic(err)
    }

    fmt.Println(result.FirstName)
    fmt.Println(result.LastName)
    fmt.Println(result.City)
}