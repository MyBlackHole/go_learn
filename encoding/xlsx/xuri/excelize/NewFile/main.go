package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Italic: true,
			Family: "Times New Roman",
			Size:   36,
			Color:  "#777777",
		},
	})

	err = f.SetCellStyle("Sheet1", "A1", "H9", style)

	if err != nil {
		fmt.Println(err)
	}

	err = f.SetRowHeight("Sheet1", 2, 50)
	if err != nil {
		return
	}

	err = f.SetColWidth("Sheet1", "A", "A", 10)
	if err != nil {
		return
	}
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	if err := f.AddChart("Sheet1", "E1", `{
	       "type": "col3DClustered",
	       "series": [
	       {
	           "name": "Sheet1!$A$2",
	           "categories": "Sheet1!$B$1:$D$1",
	           "values": "Sheet1!$B$2:$D$2"
	       },
	       {
	           "name": "Sheet1!$A$3",
	           "categories": "Sheet1!$B$1:$D$1",
	           "values": "Sheet1!$B$3:$D$3"
	       },
	       {
	           "name": "Sheet1!$A$4",
	           "categories": "Sheet1!$B$1:$D$1",
	           "values": "Sheet1!$B$4:$D$4"
	       }],
	       "title":
	       {
	           "name": "Fruit 3D Clustered Column Chart"
	       }
	   }`); err != nil {
		fmt.Println(err)
		return
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
