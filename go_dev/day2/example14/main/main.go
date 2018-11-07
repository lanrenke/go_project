package main

import (
	"time"
	"fmt"
	"math/rand"
)

func init(){//使用init来初始化，可以在程序正式运行前初始化内容，提高程序运行效率。
	//rand.Seed(time.Now().Unix())//通过秒来生成随机数，不然下面代码生成到数值并不会改变，只会生成一个值。（这里到参数time.Now().Unix()是时间戳 单位是秒）
	rand.Seed(time.Now().UnixNano())//使用纳秒来确保随机生成数基本不会出现重复
}

func main() {
	
	for i:=0;i<10;i++ {
		a := rand.Int()//生成随机整数
		fmt.Println(a)
	}
	for i:=0;i<10;i++ {
		a := rand.Intn(100)//生成小于100到随机整数
		fmt.Println(a)
	} 
	for i:=0;i<10;i++ {
		a := rand.Float32()//生成浮点数
		fmt.Println(a)
	} 
}