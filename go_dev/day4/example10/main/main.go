package main

import (
	"math/rand"
	"time"
	"fmt"
)
//输入数字然后在这个数值内生成随机数
func randnum(n int) int  {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func main() {
	var n int
	fmt.Scanf("%d",&n)
	
	fmt.Println(randnum(n))
	
}