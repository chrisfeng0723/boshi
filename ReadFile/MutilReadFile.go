/**
* @Author:fengxinlei
* @Description:
* @Version 1.0.0
* @Date: 2021/4/15 17:34
 */

package ReadFile

import (
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"os"
	"strings"
)

type Content struct {
	Line   int
	Column int
	Result float64
}

func ReadFile(fileName string) (Line int, Column []int, Result []Content) {
	Line = cast.ToInt(GetFileNumber(fileName))
	Result,Column = GetLineContent(fileName, Line)
	return
}

//读取每一行的内容

func GetLineContent(fileName string, number int) ([]Content, []int) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open"+fileName+"error!", err)
		return []Content{}, []int{}
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)
	buf := bufio.NewReader(file)
	//1.去除#开头的注释
	//2.获取每列的内容
	ContentSlice := make([]Content, 0)
	ColumnSlice := make([]int, 0)
	for {
		line, err := buf.ReadString('\n')
		if err == nil && line[0:1] != "#" {
			s := strings.Fields(strings.TrimSpace(line))
			//fmt.Println(len(s),s[0],s[1],s[2])
			LineTemp := strings.Split(s[2], ",")
			for _, val := range LineTemp {
				ColumnSlice = append(ColumnSlice, cast.ToInt(val))
				ContentSlice = append(ContentSlice, Content{
					Line:   number,
					Column: cast.ToInt(val),
					Result: cast.ToFloat64(s[0]),
				})
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(fileName+"Read file error!", err)
				return ContentSlice, ColumnSlice
			}
		}
	}
	return ContentSlice, ColumnSlice

}
