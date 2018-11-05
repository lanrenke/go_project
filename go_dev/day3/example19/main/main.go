package main

import (
	"fmt"
)

func test(a int, b int) int {
	result := func(x int,y int)int { //这里是匿名函数 调用要使用后面加括号传参
		return x + y
	}(a,b)
	
	return result
}

func main() {
	num := test(5,6)
	fmt.Println(num)
}