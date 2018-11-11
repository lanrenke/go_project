package main

import (
	"fmt"
)

func testMap() {
	//var a map[string]string //这里是声明 还没有分配内存的
	//a = make(map[string]string,10)//这里make来分配内存 初始化 方法1
	 
	var a map[string]string = map[string]string {//第二种方法
		"1":"abc",
		"2":"def",
		"3":"ghk",
	}
	
	//a["abc"] = "efg"
	//a["abc2"] = "efg2"
	//a["abc3"] = "efg3"
	fmt.Println(a)
	fmt.Println(a["abc"])
}

func testMap2() {
	a := make(map[string]map[string]string,100)//可以多重嵌套
	a["key1"] = make(map[string]string) //多重嵌套 里面嵌套的map也是需要声明
	a["key1"]["key1"] = "abc"
	a["key1"]["key2"] = "def"
	a["key1"]["key3"] = "ghk"
	fmt.Println(a)
		
}

func modify(a map[string]map[string]string) {
	_,ok := a["zhansan"] //查询是否有对应的内容
	if !ok { //这里案例可以判断一个用户 用户名是否存在 是 则执行一定的操作 否 这新建这个用户	
		a["zhansan"] = make(map[string]string) //这里简写了代码，因为不论用户是否存在 密码和昵称设置都是需要的。这里就是体现到要考虑代码维护方便的情况	
	}
	a["zhansan"]["passworld"] = "123456"
	a["zhansan"]["nickname"] = "pp"
	return
}

func testMap3() {
	
	a := make(map[string]map[string]string,100)
	modify(a)
	fmt.Println(a)
}

func  testMap4() {
	a := make(map[string]map[string]string,100)//可以多重嵌套
	a["key1"] = make(map[string]string) //多重嵌套 里面嵌套的map也是需要声明
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "def"
	a["key1"]["key4"] = "ghk"
	for k,v1 := range a {
		fmt.Println(k)
		for i,v2 := range v1 {
			fmt.Println(i,v2)
		}
	}
}

func main(){
	//testMap()
	//testMap2()
	//testMap3()
	testMap4()
}