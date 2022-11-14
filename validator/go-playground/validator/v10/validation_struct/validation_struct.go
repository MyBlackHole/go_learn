package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName string     `validate:"required"`
	LastName  string     `validate:"required"`
	Age       uint8      `validate:"gte=0,lte=130"`
	Email     string     `validate:"required,email"`
	Addresses []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func main() {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName: "Badger",
		LastName:  "Smith",
		Age:       135,
		Email:     "Badger.Smith@gmail.com",
		Addresses: []*Address{address},
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println("=== error msg ====")
		fmt.Println(err)

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		fmt.Println("\r\n=========== error field info ====================")
		for _, err := range err.(validator.ValidationErrors) {
           // 列出效验出错字段的信息
			fmt.Println("Namespace: ", err.Namespace())
			fmt.Println("Fild: ", err.Field())
			fmt.Println("StructNamespace: ", err.StructNamespace())
			fmt.Println("StructField: ", err.StructField())
			fmt.Println("Tag: ", err.Tag())
			fmt.Println("ActualTag: ", err.ActualTag())
			fmt.Println("Kind: ", err.Kind())
			fmt.Println("Type: ", err.Type())
			fmt.Println("Value: ", err.Value())
			fmt.Println("Param: ", err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}
}
