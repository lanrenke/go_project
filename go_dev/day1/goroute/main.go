package main

import (
	"time"
	
 	//"fmt"
)

func main() {
	for i:=0;i<100;i++{
		go test_goroute(i)
	}
	time.Sleep(time.Second)
	//fmt.Println(a)
}