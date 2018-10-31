package add

import(
	_ "go_dev/day2/example2/test" //引入到包只想运行init函数而不是使用是可以这样使用别名
)


var Name string 
var Age int //如果在这里赋值给了age 但是在init函数里面还是会定义一次 所以在这里定义是没有效果到


func init() {//init函数在导入包到时候就会被调用
	Name = "hello world"
	Age = 10
}

