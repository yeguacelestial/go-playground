package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func ejemplo1() {
	// Create file in memory
	var f *excelize.File
	f = excelize.NewFile()

	f.SetCellValue("Sheet1", "B2", "Valor1")
	f.SetCellValue("Sheet1", "B3", "Valor2")

	// Save file
	dir := "example.xlsx"
	if err := f.SaveAs(dir); err != nil {
		fmt.Println(err)
	}
}

func ejemplo2() {
	// Create file in memory
	var f *excelize.File
	f = excelize.NewFile()
	sheetName := "nuevo"
	index := f.NewSheet(sheetName)

	f.SetCellValue(sheetName, "A2", "Hello world")
	f.SetActiveSheet(index)
	CreateExcel(f, "example2.xlsx")
}

func CreateExcel(f *excelize.File, dir string) {
	if err := f.SaveAs(dir); err != nil {
		fmt.Println(err)
	}
}

func main() {
	ejemplo2()
}
