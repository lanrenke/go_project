package main

import (
	"fmt"
	"sort"
)

func test() {
	s := []int{3,45,3,5,6,12,56} //这里要用切片，不能使用数组 数组不能排序
	sort.Ints(s[:])
	
	fmt.Println("数字 ",s)
}

func testString() {
	s := []string{"a","h","d","a","df","gt","e"}
	sort.Strings(s[:])
	
	fmt.Println("字符串 ",s)
}

func testFloat() {
	s := []float64{3.223,234.456,2345.111,55.11,6.11}
	sort.Float64s(s[:])
	
	fmt.Println("浮点数 ",s)
}

func testIntserch() {
	s := []int{3,45,3,5,6,12,56} 
	index1 := sort.SearchInts(s[:],6)//未排序的时候查找
	fmt.Println("数字未排序查找",index1)//输出的是所查找的下标
	
	sort.Ints(s[:])
	index2 := sort.SearchInts(s[:],6)//排序的时候查找
	fmt.Println("数字排序查找",index2)
	fmt.Println("数字 ",s)
}


//sort包函数使用 排序查找
func main() {
	test()
	testString()
	testFloat()
	testIntserch()
}