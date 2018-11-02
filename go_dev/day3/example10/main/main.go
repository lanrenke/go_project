package main

import (
	"strconv"
	"fmt"
)
//讲从终端输入的值转换成数值类型 
func main() {
	var n string
	fmt.Scanf("%s",&n)
	num,err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("can not convert to int")
	}else {
		fmt.Printf("%d\n",num)
	}
}