/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2021/4/15 19:00
 */

package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func RemoveDuplicateElement(languages []int) []int {
	result := make([]int, 0)
	temp := map[int]int{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = 0
			result = append(result, item)
		}
	}
	return result
}

func TransferNumber(column,line int) string {
	result,err :=excelize.CoordinatesToCellName(column,line)
	if err !=nil{
		fmt.Println(err)
	}
	return result
}

func SliceToMap(arr []int) map[int]int{
	result :=make(map[int]int)
	for key,val := range arr{
		result[val] = key
	}
	return result
}
