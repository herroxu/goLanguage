package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"os"
	"bufio"
)

// go语言Writer和Reader接口简述
// uuid 通用唯一识别码
// 内建变量类型
// byte(8位), bool, string, rune（32位）, uintptr, int(), uint()， complex
// 强制类型转换， 枚举类型, 自增值枚举类型， 普通枚举类型
func menu()  {
	const (
		cpp  = iota
		_
		python
		golang
		java
	)
	// 自增值枚举类型
	const (
		b  = 1 << (10 * iota)
		kb
		mb
	)

	fmt.Println(cpp, python, golang, java)
	fmt.Println(b, kb, mb)
}




// 作用域，包内
var aa = 3
var ss = "kk"

func ini()  {
	var a, b int = 3, 4
	var s string = "bb"
	var f = "aa"
	var d,g = 23, false
	h, i, j := 3,"ji",true
	y := "jiu"
	fmt.Println(a,b,s,f,d,g, h, i,j, y)
}

func main()  {
	ini()
	menu()
	if content, err := ioutil.ReadFile("/Users/jumboo/Go/goTest/go语言接口/a.log"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content))
	}
	fmt.Println(strconv.Itoa(34)) // 数字转字符串
}

// case后自动break， 除非添加fallthrough
// switch后面没有表达式
func PrintFile(filename string)  {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scanner.Text()
	}

}

// 函数式编程
