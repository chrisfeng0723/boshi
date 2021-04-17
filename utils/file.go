package utils


import "strings"

//获取文件名称的序号
func GetFileNumber(fileName string) string{
s:=strings.Split(fileName,"_")
return s[1]
}