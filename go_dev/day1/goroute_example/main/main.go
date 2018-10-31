package main

import (
	"go_dev/day1/goroute_example/goroute"
	"fmt"
)

func main() {
	var pipe chan int//定义管道
	pipe = make(chan int,1)//定位管道大小
	go goroute.Add(100,300,pipe)
	
	sum := <- pipe
	
	fmt.Println("sum=",sum)
}