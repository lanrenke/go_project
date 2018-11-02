package main 

import (
	"fmt"
)

func Print(n int) { //嵌套两个循环就可以做到阶梯输出内容 
	 for i:=1;i<n+1;i++{
		for j:=0;j<i;j++ {
			fmt.Print("A")//这里是输出a
		}
		fmt.Println()//这里输出回车
	}
}

func main() {
	
	Print(6)
	
}