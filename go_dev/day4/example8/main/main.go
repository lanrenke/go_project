package main

import (
	"time"
	"fmt"
)

func recusive(n int){
	fmt.Println("hello")
	time.Sleep(time.Second)
	if n >10 {
		return
	}
	recusive(n+1)//自己调用函数本身 就变成递归
}
//这里可以通过递归算输入数字的累加
func factor(n int) int {
	if n == 1 {
		return 1
	}
	return factor(n - 1)*n
}

func main() {
	//recusive(0)
	fmt.Println(factor(10))
}