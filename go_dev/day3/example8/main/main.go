package main

import (
	"fmt"
	"time"
)

func test() {
	time.Sleep(100*time.Millisecond)
}

func main() {
	now := time. Now()
	fmt.Println(now.Format("02/1/2006 15:04:05"))//注意在go内的时间格式 必须写这个时间 2006 1 02 15 04 05
	fmt.Println(now.Format("2006/1/02"))
	
	start :=t ime.Now().UnixNano()//获取当前纳秒数
	test()
	end := time.Now().UnixNano()
	
	fmt.Printf("cost:%d\n",(end- start)/1000)
	
}