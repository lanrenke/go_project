package main

import (
	"fmt"
)

func test() {
	var a = [10]int{1,2,3,4,5}
	
	b := a[1:5]
	//验证切片使用的时候 取的地址是否为数组对应下标的地址
	fmt.Printf("%p\n",b) 
	fmt.Printf("%p\n",&a[1])
}

func main() {
	test()
}