package main

import (
	"fmt"
)

func main(){
	
	for i:=0;i<9;i++ {
		for j:=0;j<i;j++ {
			fmt.Printf("%d*%d=%d ",i+1,j+1,(i+1)*(j+1))//输出九九乘数表
		}
		fmt.Println()
	}
	
}