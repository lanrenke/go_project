package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main() {
	
	
	rand.Seed(time.Now().UnixNano())//通过seed以生成不同的随机数，如果不设置每次生成都是一样的
	num := rand.Intn(100)	
	
	for {
		var n int
		fmt.Scanf("%d",&n)
		falg := false//定义一个判断条件，不然for循环会一直运行
		switch { //条件判断的时候可以不写参数
			case (n == num):
			 	fmt.Println("您猜对了!")
				falg = true
			case (n > num):
				fmt.Println("数值大了")
			case (n < num):
				fmt.Println("数值小了")
		}
		if falg {
			break
		}
	}
	
	
	
}