package main

import (
	"fmt"
)

func test() {
	var a [10]int
	b := a
	b[0] = 100
	fmt.Println(a[0])//这里输出是0 因为数组是值类型 赋值给b到时候是复制了一个副本 ，修改副本b是不会修改a到值的
					//可以通过传地址的方法来修改
}

func main() {
	var a [10]int //定义数组 而且 var a [5]int 是两种不同到类型 数组定义完成后 里面对应到所有值都会默认为0     
	a[0] = 10
	fmt.Println(a[0])
	
	
	for i:=0;i<len(a);i++ {
		fmt.Println(a[i])
	}
	//不同的遍历方式
	for index,val := range a {
		fmt.Println(index,val)
	}
}