package main

import (
    "fmt"

    "github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
    f, err := excelize.OpenFile("./clue_raw_data_pos.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Get value from cell by given worksheet name and axis.
    // cell, err := f.GetCellValue("result 1", "B2")
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // fmt.Println(cell)
    // Get all the rows in the Sheet1.
    rows, err := f.GetRows("result 1")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}
