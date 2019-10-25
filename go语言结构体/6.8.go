package main

import "fmt"

// 结构体的匿名字段， 字段的类型是必须的

type inner struct{
	in1 int
	in2 int
}

type outer struct {
	in3 int
	int
	inner
}

func main()  {
	out := &outer{3,5,inner{1,2}}

	// go语言结构体内嵌有性质：
	// 1、内嵌的结构体可以直接访问其成员变量
	// 2、内嵌结构体的字段名就是它的类型名

	fmt.Println(out.in1)
	fmt.Println(out.in2)
	fmt.Println(out.in3)
	fmt.Println(out.int)
	fmt.Println(out.inner)

}


