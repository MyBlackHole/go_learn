package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("导出待盘点和导入模板2.xlsx")
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
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B20")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell, "-----")
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-----------")

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	// f.SetCellValue("Sheet1", "C5", "武器将阿夫斯克的解放拉斯的")
	// f.SaveAs("Book3.xlsx")
	// f.WriteToBuffer()
}
