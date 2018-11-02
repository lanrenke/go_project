package main

import (
	"fmt"
)

func main() {
	str := "hello world,中国"
	for i,v := range str { //这里使用range 方法遍历了一个字符串，i是数组下标，v是当前对应下标的值
		if i > 2 {
			continue
		}
		if (i > 3) {
			break
		}
		fmt.Printf("index[%d] val[%c] len[%d]\n",i,v,len([]byte(string(v))))//这里有两个类型转换 先转换为string类型 然后在转换成数组类型 然后len输出当前字符所占用的字节
	}
}