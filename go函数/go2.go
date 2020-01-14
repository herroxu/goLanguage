package main

import (
	"fmt"
	"strings"
)

// 斐波那契数列生成器
func fibonaqi() intGen {
	a ,b := 0 , 1
	return func() int {
		for i:=1; i<10; i++ {
			a, b = b, a+b
			fmt.Println(a)
		}
	}
}

// 为函数生成接口
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	s := fmt.Sprintf("%d", next)
	return strings.NewReader(s).Read(p)
}

func main()  {

	fu := fibonaqi()
	fu()
}
