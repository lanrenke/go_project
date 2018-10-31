package main

import (
	"fmt"
)
//也可以使用math.sqrt()函数求n的平方根 然后判断i的值到平方根即可
func isPrime(n int) bool {
	for i:=2;i<n;i++ {
		if n%i == 0 {//素数的基本概念，除了 1 和自身 外不能被其他数字整除（就是取余数结果不会为零），这里小于n的原因：比数值自身大的数 余数不可能为零200
			return false  //return 会打断当前函数运行
		}
	}
	return true
}
//输出一个数值区间内到素数
func main() {
	var n int
	var m int
	fmt.Scanf("%d%d",&n,&m)//scanf 从终端输入的数值 且传到地址 直接传到n数不会修改外部定义好到n，所以需要传到地址才可以修改对应地址的值
	
	for i := n;i<m;i++ {
		if isPrime(i) == true {
			fmt.Printf("%d\n",i)
			continue
		}
	}
	
}