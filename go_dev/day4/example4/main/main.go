package main

import (
	"os"
	"fmt"
	"bufio"
)

func check(str string) (n,m,num,other int) {
	
	t := []rune(str)
	for _,v := range t {
		switch {
			case v >= 'a' && v <= 'z'://要使用单引号  使用双引号就变成字符串
				fallthrough
			case v >= 'A' && v <= 'Z':
				n++
			case v== ' ':
				m++
			case v >= '0' && v<= '9':
				num++
			default:
				other++
		}
	}
	
	return 
}


func main(){
    reader := bufio.NewReader(os.Stdin)//建立实例从终端输入读取输入的数据
	result, _,err := reader.ReadLine()//读取输入的一行
	if err != nil {
		fmt.Println("读取输入失败")
		return
	}
	n,m,num,other := check(string(result))
	fmt.Printf("大小写字母%d个\n空格%d个\n数字%d个\n其他%d个\n",n,m,num,other)
}