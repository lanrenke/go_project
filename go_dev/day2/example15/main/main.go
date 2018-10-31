package main

import (
	"fmt"
)


func main(){
	var str = "hello world\n"
	//反引号使用
	var str1 = `
	床前明月光，
	疑是地上霜，
	举头望明月，
	我是郭德纲。
	`
	var b byte = 'c'
	
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(b)//直接输出的话就是c的ascll编码
	fmt.Printf("%c\n",b)//格式输出就会正常输出值了
}