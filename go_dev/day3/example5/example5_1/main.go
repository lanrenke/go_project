package main 

import (
	"strings"
	"fmt"
)

func urlProcess(url string) string {
	result := strings.HasPrefix(url,"http://")//判断开头是否含有https://
	if !result {//不含有则补全
		url = fmt.Sprintf("http://%s",url) //格式化处理
	}
	return url
}

func pathProcess(path string) string {
	result := strings.HasSuffix(path,"/")//判断结尾是否包含 /
	if !result {
		path = fmt.Sprintf("%s/",path)
	}
	return path
}



//strings包的使用
func main() {
	var (
		url string
		path string
		
	)
	fmt.Scanf("%s%s",&url,&path)
	
	url = urlProcess(url)
	path = pathProcess(path)
	fmt.Println(url)
	fmt.Println(path)
	
		
}


