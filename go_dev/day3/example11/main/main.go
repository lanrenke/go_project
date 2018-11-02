package main

import (
	"fmt"
)

func main() {
	var a int = 10
	switch a { //go的switch 是不需要break的
		case 0://执行条件的写法可以很多  例如 a>0 and a<10
			fmt.Println("a is equal 0")
			//fallthrough 这个关键字就是强制往下执行代码 就是会执行10内的代码
		case 10: 
			fmt.Println("a is equal 10")
		default :
			fmt.Println("a is equal default")
	}
}