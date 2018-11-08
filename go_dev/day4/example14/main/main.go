package main

import (
	"fmt"
)
//这是通过数组方法来实现斐波那契数 
func num(n int) {
	var a []uint64
	a = make([]uint64,n)//数组是必须要写准确的数值，所以这里先定义来切片，然后通过make函数通过n变量来声明
	
	a[0] = 1
	a[1] = 1
	for i:=2;i<n;i++ {
		a[i] = a[i-1] + a[i-2]
		//fmt.Println(a[i])
	}
	
	for _,v := range a { //range方法遍历输出，不需要使用的值可以通过_来消除
		fmt.Println(v)
	}
	
}

func test() {
	var a = [...]int{1,2,3,4,5}//...这里的会自动统计定义的时候有多少个数值 也可以var a [5]int
	var b = [...]string{1:"hello",3:"world"} 
	// [...]int(1:200, 3:400)//可以定义对应下标的值 注意一定要是{} 符号包括着 而不是()
	fmt.Println(a)
	fmt.Println(b)
	
}

func main(){
	//n := 30
	//num(n)
	test()
}