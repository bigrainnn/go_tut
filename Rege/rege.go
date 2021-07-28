package main

import (
	"fmt"
	"regexp"
)

const text = "my email is dygao@gmail.com@123\n" +
	"jiod 2 is aa@123.com\n" +
	"sded 3 is ccc@3213.com.cn"


func main() {
	// . 匹配任何一个字符
	// + 一个或者多个
	// .* 0个或多个
	// [a-zA-Z0-9]
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
	fmt.Println(match)
}