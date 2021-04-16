/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2021/4/13 19:25
 */

package main

import (
	"boshi/ReadFile"
	"boshi/utils"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"sort"
)

const PATH = "./txt.data"

func main() {
	LineSlice := make([]int, 0)
	ColumnSlice := make([]int, 0)
	ResultSlice := make([]ReadFile.Content, 0)
	files, _ := ioutil.ReadDir(PATH)
	for _, f := range files {
		//fmt.Println("正在处理"+f.Name())
		Line, Column, Result := ReadFile.ReadFile(PATH + "/" + f.Name())
		LineSlice = append(LineSlice, Line)
		ColumnSlice = append(ColumnSlice, Column...)
		ResultSlice = append(ResultSlice, Result...)

	}
	//获取所有的行并排序转成Map
	sort.Ints(LineSlice)
	LineMap := utils.SliceToMap(LineSlice)
	//过滤掉重复的内容
	FilterColumnSlice :=utils.RemoveDuplicateElement(ColumnSlice)
	fmt.Println(FilterColumnSlice)
	sort.Ints(FilterColumnSlice)
	ColumnMap := utils.SliceToMap(FilterColumnSlice)
	//获取所有的可能出现的数据
	fmt.Println(LineMap, ColumnMap,len(ResultSlice))
	fmt.Print(" "," ")
	for _,lval := range LineSlice{
		fmt.Print("c"+cast.ToString(lval),"  ")
	}
}
