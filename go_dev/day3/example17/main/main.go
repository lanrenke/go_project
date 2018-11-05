package main 

import (
	"fmt"
)

func add(a int,arg ...int) int {//arg 可变参数列表 可传入多个参数 当然也可以不传入
	var sum int = a 
	for i:=0;i<len(arg);i++ { //这里可以通过len函数来获取传入参数的个数
		sum += arg[i] //这里可以通过i来获取对应index的值
	}
	return sum
}

func joint(a string,arg ...string) string {
	var str string = a
	for i:=0;i<len(arg);i++ {
		str += " "+ arg[i]
	}
	return str
}

func main() {
	num := add(10,20,40,50)
	fmt.Println(num)
	
	str := joint("hello","world","i","am","ck")
	fmt.Println(str)
}