/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2021/4/13 19:25
 */

package main

import (
	"boshi/ReadFile"
	"boshi/WriteExcel"
	"fmt"
	"io/ioutil"
)

const PATH = "./txt.data"

func main() {
	LineSlice := make([]int, 0)
	ColumnSlice := make([]int, 0)
	ResultSlice := make([]ReadFile.Content, 0)
	files, _ := ioutil.ReadDir(PATH)
	for _, f := range files {
		fmt.Println("正在处理"+f.Name())
		Line, Column, Result := ReadFile.ReadFile(PATH + "/" + f.Name())
		LineSlice = append(LineSlice, Line)
		ColumnSlice = append(ColumnSlice, Column...)
		ResultSlice = append(ResultSlice, Result...)

	}
	//one step
	WriteExcel.GatherData(LineSlice,ColumnSlice,ResultSlice)

}
