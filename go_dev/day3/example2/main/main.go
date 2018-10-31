package main

import (
	"fmt"
)
//水仙花数 就是 每个位数的3次方 加起来到值位其本身 如 153  = 1的3次方 + 5的3次方+ 3的3次方
func isNumber(n int) bool {
	var i,j,k int
	i = n%10 //取余数 10 就是个位
	j = (n/10) % 10 //十位
	k = (n/100) % 10 //百位
	
	sum := i*i*i + j*j*j + k*k*k
	return sum == n //这里可以直接输出true 还false
}

func main() {
	var n int
	fmt.Scanf("%d",&n)
	/*fmt.Scanf("%d%d",&n,&m) 这部分可以输出一个区间内的水仙花数
	for i:=n;i<m;i++ {
		if isNumber(n) == true {
			fmt.Println(i)
		}
	}*/
	
	if isNumber(n) == true {
		fmt.Println(n,"is shuixianhua")
	}else {
		fmt.Println(n,"isn't shuixianhua")
	}
}