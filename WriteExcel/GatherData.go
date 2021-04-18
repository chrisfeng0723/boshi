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
	"strings"
	"time"
)

func  GatherData(LineSlice []int,ColumnSlice []int,ResultSlice []ReadFile.Content) (fileName string){
	//获取所有的行并排序转成Map

	LineMap := utils.SliceToMap(LineSlice)
	//过滤掉重复的内容

	ColumnMap := utils.SliceToMap(ColumnSlice)
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

func CalcAllHeats(fileName string,Content map[float64]int,SortValue []float64)(numberLocation map[int]string){
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
	numberLocation = make(map[int]string,len(SortValue))
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

		numberLocation[Content[val]] = "G"+CurrentLine
		start++

	}

	sumFormula := fmt.Sprintf("SUM(F1:%s)","F"+cast.ToString(totalLine))
	f.SetCellFormula("Sheet2","F"+sumLine,sumFormula)
	f.SetActiveSheet(index)
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}
	return

}

func CalcFinalWeigth(fileName string,sumCoefficient map[int]string,columnSlice []int,LineSlice []int){
	f, err := excelize.OpenFile(fileName)
	defer f.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	lenLine := len(columnSlice)
	lenColumn := len(LineSlice)
	index := f.NewSheet("Sheet2")

	for i:=0;i<lenLine;i++{
		//var sumFormula string
		sumFormulaSlice := make([]string,0)
		for key,val :=range LineSlice{
			location :=utils.TransferNumber(key+2,i+3)
			temp :=fmt.Sprintf("%s*Sheet2!%s",location,sumCoefficient[val])
			sumFormulaSlice = append(sumFormulaSlice,temp)
		}
		//fmt.Println(strings.Join(sumFormulaSlice,","))
		//fmt.Println("----------")
		resultFormula :=fmt.Sprintf("SUM(%s)",strings.Join(sumFormulaSlice,","))
		resultLocation :=utils.TransferNumber(lenColumn+2,i+3)
		f.SetCellFormula("Sheet1",resultLocation,resultFormula)
	}

	f.SetActiveSheet(index)
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}
	return

}
