package main 

import (
	"fmt"
)
//判断是否为完全数 完全数如 6= 1+2+3
func prefect(n int) bool {
	var sum int
	for i:=1;i<n;i++ {
		if n%i ==0 { //如果是一个数的因数 去余数数为零的
			sum += i
		}
	}
	return n == sum //这里输出判断值，这里的判断是判断当前输入的数与求余的数的相加值是否相等，相等则这个输入数值为完全数
}

func process(n int){
	
	for i:=1;i<(n+1);i++ {
		if (prefect(i) == true) {
			fmt.Println(i)
		}
	}
	
}

func main() {
	var n int
	fmt.Scanf("%d",&n)
	process(n)
}