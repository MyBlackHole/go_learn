package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	err := f.SetRowHeight("Sheet1", 2, 50)
	if err != nil {
		return
	}

	err = f.SetColWidth("Sheet1", "A", "A", 10)
	if err != nil {
		return
	}

	// Insert a picture.
	if err := f.AddPicture("Sheet1", "A2", "image.png", `{"autofit": true}`); err != nil {
		fmt.Println(err)
	}
	// Insert a picture to worksheet with scaling.
	if err := f.AddPicture("Sheet1", "D2", "image.jpg",
		`{"x_scale": 0.5, "y_scale": 0.5, "autofit": true}`); err != nil {
		fmt.Println(err)
	}
	// Insert a picture offset in the cell with printing support.
	if err := f.AddPicture("Sheet1", "H2", "image.gif", `{
        "x_offset": 15,
        "y_offset": 10,
        "print_obj": true,
        "lock_aspect_ratio": false,
        "autofit": true,
        "locked": false
    }`); err != nil {
		fmt.Println(err)
	}

	// Save the spreadsheet with the origin path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
