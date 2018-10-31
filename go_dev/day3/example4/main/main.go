package main

import (
	"strconv"
	"fmt"
)
//通过ascii 码来做到判断水仙花数的功能
func main(){
	/*str := "abcd"
	for i:=0;i< len(str);i++{
		fmt.Printf("%d\n",str[i])//输出对应字符的ascii编码
	}*/
	var str string
	fmt.Scanf("%s",&str)
	
	var result = 0
	for i:=0;i<len(str);i++ {
		num := int(str[i] -'0'） //这里是用对应获取到到字符到ascii码减去数字 0 的 ascii 码来还原当前获取到的字是什么数字
		result += num*num*num   //ascii码数连续的
	}
	
	number ,err := strconv.Atoi(str)//把string型转换为int型  如果无法转换则报错
	if (err != nil) { //注意判断为空的写法是nil 不是null
		fmt.Println("can not convert %s to int",str)
		return
	}
	if result == number {
		fmt.Println(number,"is shuixianhua")
	}else {
		fmt.Println(number,"isn't shuixianhua")
	}
	
}