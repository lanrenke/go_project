package	main

import (
	"time"
	"fmt"
)

const (
	Man = 1 
	Female = 2
)

func main() {
	for  {
		second := time.Now().Unix()//获取当前时间到秒数到方法 
		if(second % Female == 0){
			fmt.Println("Female")
		}else {
			fmt.Println("Man")
		}
		time.Sleep(1000*time.Millisecond)//Millisecond是毫秒
	}
	
}