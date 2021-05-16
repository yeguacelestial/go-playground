package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type memberStat struct {
	Total     int `title:"total" sheet:"Test Sheet 1"`
	New       int `title:"new"`
	Effective int `title:"effective"`
}

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

func CreateSheet(f *excelize.File, sheetName string, data [][]map[string]string) *excelize.File {
	if f == nil {
		f = excelize.NewFile()
	}

	index := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	for _, col := range data {
		for _, row := range col {
			for index, value := range row {
				f.SetCellValue(sheetName, index, value)
			}
		}
	}

	return f
}

func ArrayForExcel(file [][]map[string]string, row ...string) [][]map[string]string {
	var data []map[string]string
	var ExcelRow map[string]string = make(map[string]string)

	col := []string{
		"A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U",
		"X", "Y", "Z"}

	for index, value := range row {
		ExcelCol := col[index] + fmt.Sprint(len(file)+1)
		ExcelRow[ExcelCol] = value
		data = append(data, ExcelRow)
	}

	file = append(file, data)

	return file
}

func ejemplo3() {
	var result [][]map[string]string

	result = ArrayForExcel(result, "Col 1 Row 1", "Col 2 Row 1", "Col 3 Row 1", "Col 4 Row 1")
	result = ArrayForExcel(result, "Col 1 Row 2", "Col 2 Row 2", "Col 3 Row 2", "Col 4 Row 2")
	result = ArrayForExcel(result, "Col 1 Row 3", "Col 2 Row 3", "Col 3 Row 3", "Col 4 Row 3")

	var arrtostring []string

	result = ArrayForExcel(result, arrtostring...)

	f := CreateSheet(nil, "Sheet test", result)
	CreateExcel(f, "example3.xlsx")
}

func main() {
	ejemplo3()
}
