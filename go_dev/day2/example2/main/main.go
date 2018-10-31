package main

import(
	a "go_dev/day2/example2/add" //这里使用了别名
	
	"fmt"
)

func main(){
	
	//a.Test()调用函数 不然值会是初始化的值，而且还会报包没有使用到到错误.如果在init函数里面有声明 这不需要调用
	fmt.Println("Name=",a.Name)
	fmt.Println("Age=",a.Age)
}