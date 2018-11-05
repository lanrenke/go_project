package main

import (
	"fmt"
)

type add_func func(int,int) int //这里定义一个类型 类型是函数 函数也可以是一种类型
	
func add(a,b int) int {
	return a + b
}
	
func operator(op add_func,a int ,b int) int {//这里又定义了op的类型是一个函数
	return op(a,b)
}

func main() {
	c := add
	fmt.Println(c)//这输出的是add函数的地址
	sum := operator(c,100,200)
	fmt.Println(sum)
	
}