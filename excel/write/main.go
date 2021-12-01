package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()

	sheetName := "Sheet2"

	// 创建一个工作表
	index := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	headers := map[string]string{"A1": "名称", "B1": "类型", "C1": "开始时间", "D1": "结束时间", "E1": "备注"}
	for k, v := range headers {
		f.SetCellValue(sheetName, k, v)
	}

	// 设置宽度
	if err := f.SetColWidth(sheetName, "A", "E", 20); err != nil {
		fmt.Println(err)
	}

	values := []map[string]interface{}{
		{
			"A2": "测试",
			"B2": "test",
			"C2": "2021-08-15",
			"D2": "2021-09-16",
			"E2": "这里是备注",
		},
		{
			"A3": "第二行",
			"B3": "6666",
			"C3": "2021-01-02",
			"D3": "2020-10-16",
			"E3": "这里是备注",
		},
	}

	for _, rows := range values {
		for k, v := range rows {
			f.SetCellValue(sheetName, k, v)
		}
	}

	// 设置样式
	style, err := f.NewStyle(`{"alignment": {"horizontal": "center", "vertical": "center"}}`)
	if err != nil {
		fmt.Println(err)
	}
	f.SetCellStyle(sheetName, "A1", "E10", style)
	headerStyle, err := f.NewStyle(`{"font":{"bold": true}, "alignment": {"horizontal": "center", "vertical": "center"}}`)
	if err != nil {
		fmt.Println(err)
	}
	f.SetCellStyle(sheetName, "A1", "E1", headerStyle)

	if err := f.SaveAs("demo.xlsx"); err != nil {
		fmt.Println(err)
	}
}
