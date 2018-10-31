package test

import (
	"fmt"
)

var Name string = "this is in test package"
var Age int = 100  //这里变量命名和add包里面函数命名是一样到，但是不会用影响，从调用到时候就可以看到 add包到变量调用是add.Age,而test包里面到调用
					//是test.Age,是不会相互影响到 

func init() {
	fmt.Println("this is a test")
	fmt.Println("test.package.Name=",Name)
	fmt.Println("test.package.Age=",Age)
	
	Age = 10
	fmt.Println("test.package.Age=",Age)
}