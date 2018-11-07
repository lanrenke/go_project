package main

import (
	"errors"
	"fmt"
)


func initConfig()(err error){
	return errors.New("init config failed")//导入错误包，可以创建一个实例来输出需要指定的内容
}


func test(){
	
	defer func(){//匿名函数 
		if err := recover();err != nil { //通过recover来抓取错误 并跳过这个错误 从而不影响整体程序的运行
			fmt.Println(err)
		}
	}()
	
	b := 0
	a := 100/b
	fmt.Println(a)
	return
}

func main(){
	//test()
	err := initConfig()
	if err != nil{
		panic(err)//panic会终止程序的运行 可以与reover配合是程序可以继续运行，特别是在使用goroute 并发的时候
	}
	
	
	var i int
	fmt.Println(i)
	
	j := new(int)//通过new来分配内存创建 创建的是地址 是用来分配值类型 make是分配引用类型 如chan map slice
	*j = 100
	fmt.Println(*j)
	
	var a []int
	a = append(a,10,23,45)//append的作用是把元素追加到数组 slice中
	a = append(a,a...)//这里...是把a数组原有到内容是简写到里面  a...的意思其实就是 10 23 45
	fmt.Println(a)
	
	
}