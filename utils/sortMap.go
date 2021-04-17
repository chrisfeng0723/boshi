/**
 * @Author: fxl
 * @Description: 
 * @File:  sortMap.go
 * @Version: 1.0.0
 * @Date: 2021/4/17 17:53
 */
package utils

import (
	"sort"
)

func SortMapByKey(heads map[float64]int) []float64{
	value := make([]float64,0)
	for key := range heads{
		value =append(value,key)
	}
	sort.Float64s(value)
	return value
}
