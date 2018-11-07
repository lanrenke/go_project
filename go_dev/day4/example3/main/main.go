package main

import (
	"fmt"
)
//判断输入的字符是否数回文。回文的意思就如 abba  aba这样 前后两个对应对称位置的值相同
func process(str string) bool {
	for i:=0;i<len(str);i++ { //这里对中文无效
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
//这个方法适用中文加英文
func check(str string) bool {
	t :=[]rune(str)//把输入的参数转变成一个一个字符形式，无论是中文还是英文 字符长度多少
	length := len(t)
	for i,_ := range t {//通过range来获取一个字符 
		if(i == length/2){
			break
		}
		last := length - i - 1
		if t[i] != t[last] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scanf("%s",&str)
	if check(str) {
		fmt.Println("yes")
	}else {
		fmt.Println("no")
	}
}