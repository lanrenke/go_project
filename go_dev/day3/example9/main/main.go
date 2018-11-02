package main 

import (
	"fmt"
)
 
func modify(p *int) {
	fmt.Println(p)
	*p = 100
	return
}

func main() {
	var a int
	a = 10
	fmt.Println(&a) //输出a变量的地址
	
	var p *int //这样写是声明指针
	p = &a
	
	fmt.Println(*p) //这样就可以输出指针指向的值
	
	var b int = 999
	p = &b
	*p = 5//这个时候修改的是b的值而不是a的值了
	
	fmt.Println(a)
	fmt.Println(b)
	
	modify(&a)
	fmt.Println(a)
}