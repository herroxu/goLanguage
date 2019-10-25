package main

import (
	"fmt"
	"bytes"
)

func MyPrint(args ...interface{})  {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "int value")
		case string:
			fmt.Println(arg, "string value")
		}
	}
}

func joinStrings(slist ...string) string {
	// 定义一个字节缓冲，快速的连接字符串
	var b bytes.Buffer
	for _,s := range slist {
		b.WriteString(s)
	}
	return b.String()
}


func main() {

	s := "ming"
	i := 12
	MyPrint(s, i)

	joinStrings("zu", "ming", "dong")
}
