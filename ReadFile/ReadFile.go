/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2021/4/13 19:30
 */

package ReadFile

const PATH = "./txt.data"
var Line,column []int
var result map[int]map[int]float64

//获取原始数据文件名
/**
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
*/







/**
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
*/


