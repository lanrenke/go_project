package main

import (
	"fmt"
)
//遍历多维数组
func test() {
	var a [2][3]int = [...][3]int{{1,2},{4,5,6}}
	for row,v1 := range a {
		for cow,v2 := range v1 {
			fmt.Printf("(%d,%d)=%d\n",row,cow,v2)
		}
	}
}

func main() {
	test()
}