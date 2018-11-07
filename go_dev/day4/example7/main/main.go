package main

import (
	"fmt"
)

func test() {
	s1 := new([]int)//这里是定义 还没初始化
	fmt.Println(s1)//输出地址
	
	s2 := make([]int,10)
	fmt.Println(s2)//输出列表内容(就是输出list 值)
	
	*s1 = make([]int,5)
	(*s1)[0] = 100
	s2[0] = 100
	fmt.Println(s1)
	
	return
}

func main() {
	test()
}