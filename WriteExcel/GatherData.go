/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2021/4/16 11:30
 */

package WriteExcel

import (
	"boshi/ReadFile"
	"boshi/utils"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/spf13/cast"
	"sort"
	"time"
)

func  GatherData(LineSlice []int,ColumnSlice []int,ResultSlice []ReadFile.Content){
	//获取所有的行并排序转成Map
	sort.Ints(LineSlice)
	LineMap := utils.SliceToMap(LineSlice)
	//过滤掉重复的内容
	FilterColumnSlice :=utils.RemoveDuplicateElement(ColumnSlice)
	sort.Ints(FilterColumnSlice)
	ColumnMap := utils.SliceToMap(FilterColumnSlice)
	//获取所有的可能出现的数据



	//write excel

	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	// Set first Line
	for _,v := range LineSlice{
		location :=utils.TransferNumber(LineMap[v]+2,1)
		f.SetCellValue("Sheet1", location,"C"+cast.ToString(v))
	}

	//set first column
	for _,v := range ColumnSlice{
		location :=utils.TransferNumber(1,ColumnMap[v]+3)
		f.SetCellValue("Sheet1", location,cast.ToString(v))
	}

	//set the data
	for _,val := range ResultSlice{
			excelColumn := ColumnMap[val.Column]+3
			excelLine  := LineMap[val.Line]+2
			location:= utils.TransferNumber(excelLine,excelColumn)
		    f.SetCellValue("Sheet1", location,val.Result)
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	fileName := time.Now().Format("20060102150405")
	if err := f.SaveAs(fileName+".xlsx"); err != nil {
		fmt.Println(err)
	}
}
