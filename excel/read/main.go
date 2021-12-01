package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("demo.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	sheetName := "Sheet2"

	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

}
