package main

import (
	"fmt"
)

func test() {
	var a []int = []int{1,2,3,4,5,6}//定义切片 有值的ß
	b := make([]int,10) //定义一个空值切片
	
	copy(b,a) //从a中拷贝值到b，如果拷贝值不超过大小 则自动补0 当如果超过大小 则只会拷贝对应大小的值 ，如b只有10位 ，超过10位的值不会被拷贝到
	//拷贝是不会扩容的
	fmt.Println( b)
	
	
	//string字符输出
	s := "hello world"
	s1 :=s[0:5]
	s2 := s[6:]
	fmt.Println(s1)
	fmt.Println(s2)
	
	
	
}

func test1() {
	//特别提醒string是没有办法修改的，字符串是不可变的 如s1[1] = a 是不能通过编译的
	//所以要修改一个字符串的字符 就需要转化一下 转化成数组
	s :="我hello world"
	s1 := []rune(s)//通过rune函数转化成数组 用其他的函数 非英文会有问题 
	
	s1[0] = '0'
	str := string(s1)
	fmt.Println(str)
}
//切片拷贝
func main() {
	//test()
	test1()
}