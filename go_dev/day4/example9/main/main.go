package main

import (
	"fmt"
)

func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n - 1) + fab(n-2)//递归都使用场景 尽量是在关键值在递减都时候
}

//斐波那契数  0 1 位都是1 然后从2 位开始 都数前两项之和
func main() {
	
	for i:=0;i<10;i++ {
		fmt.Println(fab(i))
	}
	
}