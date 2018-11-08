package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func (string) string {
	return func(name string) string {
		if !strings.HasSuffix(name,suffix) {
			return name + suffix
		}
		return name
	}
}

//闭包例子
func main() {
	func1 := makeSuffixFunc(".bmp")//这里绑定了suffix 值 然后返回一个函数给到foun1
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))
}