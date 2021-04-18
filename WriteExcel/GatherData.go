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

func  GatherData(LineSlice []int,ColumnSlice []int,ResultSlice []ReadFile.Content) (fileName string){
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
	fileName = time.Now().Format("20060102150405")+".xlsx"
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
	return fileName
}

func CalcAllHeats(fileName string,Content map[float64]int,SortValue []float64){
	f, err := excelize.OpenFile(fileName)
	defer f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	index := f.NewSheet("Sheet2")
	start :=1
	totalLine :=len(Content)
	sumLine :=cast.ToString(totalLine+1)
	for _,val := range SortValue{
		CurrentLine := cast.ToString(start)
		f.SetCellValue("Sheet2", "A"+CurrentLine,Content[val])
		Column2 :="B"+CurrentLine
		f.SetCellValue("Sheet2",Column2,val)
		formula3 :=	Column2+"-B1"
		f.SetCellFormula("Sheet2","C"+CurrentLine,formula3)
		formula4 :=	"C"+CurrentLine+"*627.5"
		f.SetCellFormula("Sheet2","D"+CurrentLine,formula4)
		formula5 :=	"-D"+CurrentLine+"/(0.0019858955*298.15)"
		f.SetCellFormula("Sheet2","E"+CurrentLine,formula5)
		formula6 :=	"EXP(E"+CurrentLine+")"
		f.SetCellFormula("Sheet2","F"+CurrentLine,formula6)
		formula7 :=	"F"+CurrentLine+"/"+"F"+sumLine
		f.SetCellFormula("Sheet2","G"+CurrentLine,formula7)
		start++

	}

	sumFormula := fmt.Sprintf("SUM(F1:%s)","F"+cast.ToString(totalLine))
	f.SetCellFormula("Sheet2","F"+sumLine,sumFormula)
	f.SetActiveSheet(index)
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}

}
