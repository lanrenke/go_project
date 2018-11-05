package main

import (
	"fmt"
)
//判断输入的字符是否数回文。回文的意思就如 abba  aba这样 前后两个对应对称位置的值相同
func process(str string) bool {
	for i:=0;i<len(str);i++ {
		if(i == len(str)/2){
			break
		}
		last := len(str) - i - 1//这里是赋值index 对应指的位置的当前i位置的对称的index的的位置 减去1 是因为数位是0开始的
		if (str[i] != str[last]) {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scanf("%s",&str)
	if process(str) {
		fmt.Println("yes")
	}else {
		fmt.Println("no")
	}
}