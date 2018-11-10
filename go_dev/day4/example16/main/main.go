package main

import (
	"fmt"
)

//定义一个结构体
type slice struct { //这里是一个切片的内部结构
	ptr *[100]int //写死大小 纯粹演示
	len int
	cap int
}

//定一个make 仅限演示 不能直接命名make make是系统的函数
func make1(s slice,cap int) slice {
	s.ptr = new([100]int)//这里使用new是因为在结构体声明的时候 ptr是指针
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice) {
	s.ptr[1] = 1000 //这里是生效的 会被修改 因为这里是引用类型的，本身是有指针指定来修改 值就会修改 了。
}

func testSlice1() {
	var s1 slice
	s1 = make1(s1,10)
	
	s1.ptr[0] = 100
	modify(s1)
	
	fmt.Println(s1.ptr)
}

//这里简单写法，切片是引用类型，传值修改是会修改对应地址的值。而不像值类型的那样传值是会进行副本复制。
func modify1(a []int) {
	a[1] = 1000
}
func testSlice2() {
	var b []int = []int{1,2,3,4} 
	modify1(b)
	fmt.Println(b)
}



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
	//test()
	//testSlice1()
	testSlice2()
}