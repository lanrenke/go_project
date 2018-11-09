package main

import (
	"fmt"
)

func test() {
	var slice []int //这里是定义切片
	var arr [5]int = [...]int{1,2,3,4,5}
					 //这里已经超过下标4了 所以就把后面的全部输出
	slice = arr[2:5]//注意要非常注意 是包含左边不包含右边边界 就是包含 2 3 而不包含4 还要注意 下标是从0开始的
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))//长度len和容量cap不一定相同
	
	slice = arr[0:1]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))//长度len和容量cap不一定相同
	
	var slice2 []int
	slice2 = arr[:]//当包含整个数组的时候 直接冒号即可
	fmt.Println(slice2)
}

func main(){
	test()
}