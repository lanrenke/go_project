package main

import(
	"fmt"
)

//应该有的程序逻辑
func list(n int){
	for i:=0;i<=n;i++{
		fmt.Printf("%d+%d=%d\n",i,n-i,n)
	}
	
}

//main函数作为执行函数 里面不放执行方法是毕竟好的选择
func main() {
	list(10)
}


//这是一种写法 
//func main() {
//	n := 5
//	for i:=0;i<=n;i++{
		
//		fmt.Println(i,"+",(n - i),"=",n)
		
//	}
//}