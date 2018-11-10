package main

import (
	"fmt"
)



func main() {
	i := 0
	/*var p *int  这里做了一下试验，在defer声明的时候传入i的地址，然后在下面代码修改i的值，在执行defer的时候并不会输出修改后的值，而是一开始声明的值
	p = &i*/
	defer fmt.Println("1",i)//defer 声明的 是要函数return的时候执行（函数结束的时候）要写在return之前
	defer fmt.Println("2",i)//先声明后执行，是通过栈来存放 先入后输出。所以程序运行的时候输出的内容是先输出2 再到 1
	//defer的作用 是在程序结束之后执行一些代码 比如在声明的时候锁定了某些资源，如果不解锁go是不会自动解锁的，所以在什么声明的同时写上defer来解锁就会解决这个问题。
	
	i = 10
	/*var x *int
	x = &i
	fmt.Println(*x)*/
	fmt.Println(i)
	
	for j:=0;j<5;j++ {
		defer fmt.Printf("%d\n",j)
	}
	
}