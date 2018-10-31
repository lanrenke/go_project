package main

import (
	"fmt"
)

//字符串反转
func reverse(str string) string {
	var result string
	strLen := len(str)
	for i:=0;i<strLen;i++ {
		result = result + fmt.Sprintf("%c",str[strLen - i -1])//str[xx]当xx是数字到时候 str[xx]返回到是二进制码，通过sprintf来转换为对应到单词，这样就可以拼接起来了
	}
	return result
}
//字符串反转 数组方法
func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)//转换成数组类型
	lenght := len(str)
	for i:=0;i<lenght;i++ {
		result = append(result,tmp[lenght -i -1])
	}
	return string(result)
}

func main(){
	var str1 = "hello"
	str2 := "world"
	//str3 := str1 + " "+ str2
	str3 := fmt.Sprintf("%s %s",str1,str2)//格式化拼接
	n := len(str3)//获取字符长度
	
	
	fmt.Println(str3)
	fmt.Printf("len(str3)=%d\n",n)
	
	substr := str3[0:5]//切片取值，在取前字符中从哪个位置到哪个位置到值
	fmt.Println(substr)
	
	substr = str3[6:]//不填后面到值则会直接获取剩余到部分
	fmt.Println(substr)
	
	result := reverse(str3)
	fmt.Println(result)
	
	result = reverse1(str3)
	fmt.Println(result)
	
}