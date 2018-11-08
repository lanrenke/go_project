package main

import (
	"fmt"
)


func Adder() func (int) int {
	var x int
	return func(delta int) int {//闭包中 如果函数中使用了外部定义的变量 就是一个闭包 而这个x的值是会保存之前所运行的值
		x += delta  			//这部分相当于一个类 把x当成一个公共参数
	 	return x
	}
}

//闭包：一个函数与其相关的引用环境组合而成的实体
func main(){
	var f = Adder()//注意这里是声明f的是个函数，但是声明的函数是return回来的函数而不是adder函数体本身
	fmt.Print(f(1)," - ")
	fmt.Print(f(20)," - ")
	fmt.Print(f(300),"\n")
}
//输出结果 1-21-321