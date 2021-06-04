package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get value from cell by given worksheet and axis.
	cell, err := f.GetCellValue("CALCULO FINIQUITO", "B10")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cell)

	// Get all the rows in the Sheet
	rows, err := f.GetRows("CALCULO FINIQUITO")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
