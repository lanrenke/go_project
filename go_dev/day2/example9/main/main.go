package main

import (
	"fmt"
)

func swap(a *int,b *int){//修改指向地址到值 这样就可以做到 a b值 ‘互换’。变量不带* 就是地址到值
	tmp := *a
	*a = *b
	*b = tmp
	return
}

func swap1(a int,b int)(int,int) {//这里使用返回值到方法完成值交换 
	return b,a
}

func test() {
	var a = 100
	fmt.Println(a)
	
	for i:=0;i<100;i++ {//b是在语句块里面声明的（就是在大括号里面声明到），只要离开语句块b就会被内存回收，下面到代码是无法访问到b到内容
		var b = i*2 
		fmt.Println(b)
	}
}

func main(){
	first := 100
	second := 200
	//swap(&first,&second)//注意 传地址要使用 & 符号，这是传地址修改对应地址值的方法
	first,second = swap1(first,second)//这里是返回值方法，直接修改当前参数但值
	//first,second = second,first 这个是方法最直接，直接交换即可
	fmt.Println("first=",first)
	fmt.Println("second=",second)
	
	test()
}