package main

import (
	"fmt"
)

// 字符串翻转
func reverse(str string) string {
	result := ""
	for i := len(str); i > 0; i-- {
		result = result + str[i-1:i]
	}
	return result
}

func reverseStr(str string) string {
	result := ""
	for i := len(str); i > 0; i-- {
		result = result + fmt.Sprintf("%c", str[i-1])
	}
	return result
}

func reverseArr(str string) string {
	var result []byte //定义一个空数组
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, str[length-i-1]) //通过append往数组里追加值
	}
	return string(result) //最后在string化
}

func main() {
	var str = "liujingzhao"
	result := reverse(str)
	resStr := reverseStr(str)
	res := reverseArr(str)
	fmt.Println(result, resStr, res)

}
