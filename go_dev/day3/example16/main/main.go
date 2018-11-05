package main

import (
	"fmt"
)

func add(a,b int) int {
	return a+b
}

func main() {
	c := add
	fmt.Println(c)
	 
	sum := c(10,20)
	fmt.Println(sum)
	/*if (c == add){ //无法直接比较 无法编译
		fmt.Println("c equal add")
	}*/
	
}