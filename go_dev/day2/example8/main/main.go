package main

import (
	"fmt"
)

func modifr(a int){
	a = 10
	return
}

func modifr1(a *int){//加*是指向内存地址
	*a = 10
}

func main(){
	a:= 5
	b:=make(chan int,1)
	
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	
	modifr(a)//调用了函数 但是a到值是复制后传到内存里面的，指向到地址不同了，所以值是不会改变
	fmt.Println("a=",a)
	modifr1(&a)//这里是传指针，对应地址到值改了，指针指过去到值也会修改
	fmt.Println("a=",a)
}