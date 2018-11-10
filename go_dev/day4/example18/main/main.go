package main

import (
	"fmt"
)


//通过make创建slice 区别是 make创建其实也是在数组基础上创建，但是这个数组是看不到的
func main() {
	//var slice []int = make([]int,len) //不同的写法
	
	//slice1 := make([]int,len)
	
	//slice2 := make([]int,len,cap)
	
	var a [5]int = [...]int{1,2,3,4,5} //通过数组方法来创建
	s := a[1:]
	
	s[1] = 100
	fmt.Println(s)
	fmt.Println("扩容前",len(s),cap(s))
	fmt.Println("1",a) //这里输出可以看到通过切片修改下标1的值，a里面对应下标值会修改，因为地址是相同的
	//append函数可以在原有的切片上添加，如果超过大小会新建空间存放 这个时候跟一开始的地址就不一样了
	s = append(s,10)
	s = append(s,10)
	s = append(s,10)
	fmt.Println("扩容后",len(s),cap(s))	
	s = append(s,10)
	s = append(s,10)
	s = append(s,10)
	
	s[1]=10000
	fmt.Println("3",a)//而这里在再次输出a是就会发现并没有修改，是因为s已经新建了内存存放，地址已经不同了
	fmt.Println(s)
}