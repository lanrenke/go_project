package main

import (
	"strconv"
	"fmt"
	"strings"
)

//去掉字符串首尾空白字符
func test_1(n string) string {
	str := strings.TrimSpace(n)
	return str
}
//去掉字符串首尾 对应 字符
func test_2(n string) string {
	str := strings.Trim(n,"h")//注意一定要是双引号 不然会报错
	return str
}
//去掉字符串左边部分 对应 字符
func test_3(n string) string {
	str := strings.TrimLeft(n,"h")
	return str
}
//去掉字符串右边部分 对应 字符
func test_4(n string) string {
	str := strings.TrimRight(n,"h")
	return str
}
//返回str空隔分隔的所有子串的 字符
func test_5(n string) []string {
	str := strings.Fields(n)
	return str
}
//使用指定的内容分隔开所有子串
func test_6(n string) []string {
	str := strings.Split(n,",")//这里指定的内容是，号
	return str
}

func main() {
	
	test1 := test_1(" hello ") //这里两边的空格会去掉
	fmt.Printf("test1=%s\n",test1)
	
	test2 := test_2("hhelloh") //这里两边所有的h都会去掉
	fmt.Printf("test2=%s\n",test2)
		
	test3 := test_3("hhelloh") //这里左边所有的h都会去掉
	fmt.Printf("test3=%s\n",test3)
	
	test4 := test_4("hhelloh") //这里右边所有的h都会去掉
	fmt.Printf("test4=%s\n",test4)	
		
	test5 := test_5("hello world i am ck")	
	fmt.Printf("test5=%q\n",test5)
	
	test6 := test_6("hello,world,i,am,ck")	
	fmt.Printf("test6=%q\n",test6)
	
	test7 := strconv.Itoa(1000)//把int转换成string型
	fmt.Printf("test7=%s\n",test7)
	
	test8,err := strconv.Atoi("151")
	if err !=nil {
		fmt.Println("转化失败",err)
	}else {
		fmt.Printf("test8=%v\n",test8)
	}
	
	
	
}