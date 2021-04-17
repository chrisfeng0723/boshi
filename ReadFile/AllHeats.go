/**
 * @Author: fxl
 * @Description: 
 * @File:  AllHeats.go
 * @Version: 1.0.0
 * @Date: 2021/4/17 17:24
 */
package ReadFile

import (
	"boshi/utils"
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"os"
	"strings"
)

func GetAllHeats(fileName string) (map[float64]int) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open"+fileName+"error!", err)
		return map[float64]int{}
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	//1.去除#开头的注释
	//2.获取每列的内容
	HeatsMap := make(map[float64]int, 0)
	for {
		line, err := buf.ReadString('\n')
		if err == nil{
			s := strings.Fields(strings.TrimSpace(line))
			fileNumber :=utils.GetFileNumber(s[0])
			value := cast.ToFloat64(s[1])
			HeatsMap[value] = cast.ToInt(fileNumber)

		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(fileName+"Read file error!", err)
				return map[float64]int{}
			}
		}
	}
	return HeatsMap

}

