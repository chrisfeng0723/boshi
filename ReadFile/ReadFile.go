/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2021/4/13 19:30
 */

package ReadFile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"github.com/spf13/cast"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)
const PATH = "./txt.data"
var Line,column []int
var result map[int]map[int]float64

//获取原始数据文件名
func ReadFile(){
	result = make(map[int]map[int]float64,0)
	files,_:=ioutil.ReadDir(PATH)
	for _,f :=range files{
		fmt.Println("正在处理"+f.Name())
		number := cast.ToInt(GetFileNumber(f.Name()))
		Line = append(Line,number)
		GetLineContent(PATH+"/"+f.Name(),number)

	}
	//fmt.Println(len(Line))
	sort.Ints(Line)
	//mLine := sliceToMap(Line)
	tempColumn :=removeDuplicateElement(column)
	sort.Ints(tempColumn)
	mColumn :=sliceToMap(tempColumn)
	//fmt.Println(Line,tempColumn,result)
	WriteExcel(mColumn,result)
}

//获取文件名称的序号
func GetFileNumber(fileName string) string{
	s:=strings.Split(fileName,"_")
	return s[1]
}


//读取每一行的内容

func GetLineContent(fileName string,number int) {
	//fileName = "./txt.data/"+fileName
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size)
	buf := bufio.NewReader(file)
	temp := make(map[int]float64,0)
	for {
		line, err := buf.ReadString('\n')
		if err == nil && line[0:1] != "#"{
			s :=strings.Fields(strings.TrimSpace(line))
			//fmt.Println(len(s),s[0],s[1],s[2])
			LineTemp :=strings.Split(s[2],",")
			for _,val := range LineTemp{
				column = append(column,cast.ToInt(val))
				tempKey := cast.ToInt(val)
				tempValue:= cast.ToFloat64(s[0])
				temp[tempKey]= tempValue
			}
		}
		if err != nil {
			if err == io.EOF {
				//fmt.Println(fileName+"File read ok!")
				break
			} else {
				fmt.Println(fileName+"Read file error!", err)
				return
			}
		}else{

		}
	}
	result[number] = temp
}

func removeDuplicateElement(languages []int) []int {
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

func WriteExcel(column map[int]int,result map[int]map[int]float64){
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	for key,val := range result{
		for skey,sval :=range val{
			fmt.Println(key,skey,sval)
			//实际的行数
			fcloumn :=column[skey]
			fmt.Println(fcloumn)
			location :=TransferNumber(key+1,fcloumn+1+2)
			fmt.Println(location,"---------------")
			//f.SetCellValue("Sheet1",location , sval)
		}
	}
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func TransferNumber(column,line int) string {
	result,err :=excelize.CoordinatesToCellName(column,line)
	if err !=nil{
		fmt.Println(err)
	}
	return result
}

func sliceToMap(arr []int) map[int]int{
	result :=make(map[int]int)
	for key,val := range arr{
		result[val] = key
	}
	return result
}