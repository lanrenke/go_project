package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func multi(str1 string,str2 string) (result string) {
	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}
	var index1 = len(str1) - 1
	var index2 = len(str2) - 1
	var left int //记录进位
	
	for index1 >= 0 && index2 >=0 {
		c1 := str1[index1] - '0' //通过acsii码相减来获取对应的数字
		c2 := str2[index2] - '0' //从最后一位开始计算
		
		sum := int(c1 + c2) + left //需要加上进位 如果两个数相加超过10 下一位相加数需要加上进位的
		if sum >=10 {
			left = 1  //如果有进位就赋值为1
		}else {
			left = 0
		}
		c3 := (sum % 10) + '0' //先求余数然后加回0的acsii码变回字符
		result = fmt.Sprintf("%c%s",c3,result)//将结果拼接起来
		
		index1-- 
		index2--
	}
	//fmt.Println(index1)
	//fmt.Println(index2)
	
	//下面两种情况就是 当两个数的数位不等时需要单独处理 上面的for循环已经把对应的index的值减掉 代码执行到这部分到时候 index到值已经过滤到 两个数 相同位数到部分 
	//剩下到位数由下面代码处理
	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >=10 {
			left = 1  //如果有进位就赋值为1
		}else {
			left = 0
		}
		c3 := (sum % 10) + '0' //先求余数然后加回0的acsii码变回字符
		result = fmt.Sprintf("%c%s",c3,result)
		fmt.Println(index1)
		index1--
	}
	
	for index2 >= 0 {
		c2 := str2[index2] - '0'
		sum := int(c2) + left
		if sum >=10 {
			left = 1  //如果有进位就赋值为1
		}else {
			left = 0
		}
		c3 := (sum % 10) + '0' //先求余数然后加回0的acsii码变回字符
		result = fmt.Sprintf("%c%s",c3,result)
		index2--
	}
	
	if left == 1 {//上面的循环已经输出来所有的数 但是进位没有输出 所以需要另外处理
		result = fmt.Sprintf("1%s",result)
	}
	
	return
}

//面试考题 两个大数相加，值数要能超过int64
func main() {
	reader := bufio.NewReader(os.Stdin)//建立实例从终端输入读取输入的数据
	result, _,err := reader.ReadLine()//读取输入的一行
	if err != nil {
		fmt.Println("读取输入失败")
		return
	}
	
	strSlice := strings.Split(string(result),"+") //通过 + 符号区分开两个字符
	if len(strSlice) != 2 {
		fmt.Println("输入错误，需要输入两个大数相加且有 + 符号")
		return
	}
	strNumber1 := strings.TrimSpace(strSlice[0])//去除空格
	strNumber2 := strings.TrimSpace(strSlice[1])
	
	fmt.Println(multi(strNumber1,strNumber2))
}