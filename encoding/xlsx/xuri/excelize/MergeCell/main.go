package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("./abs.xlsx")
	// f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	f.SetCellValue("Sheet1", "B4", "Hello world.")
	f.SetCellValue("Sheet1", "B5", "Hello world.")

	err = f.Save()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
