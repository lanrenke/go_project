package main

import (
	"fmt"
)
//求n个数阶乘之和
func sum(n int) uint64 {
	var s uint64 = 1
	var sum uint64 = 0
	for i:=1;i<=n;i++ {
		s = s * uint64(i) //举例理解： 4的阶乘数 1*2*3*4 其实就是等于 3的阶乘 乘于 4 以此类推 3的阶乘等于 2的阶乘乘于 3  2的阶乘就是1 乘于2
		sum += s
	}
	return s
}

func main() {
	var n int
	fmt.Scanf("%d",&n)

	s := sum(n)
	fmt.Println(s)
}