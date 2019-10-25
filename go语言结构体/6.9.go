package main

import "fmt"

// 结构体内嵌 实现 对象特性 组合
type Flying struct {}

func (f *Flying) Fly()  {
	fmt.Println("I can fly")
}

type Walking struct {}

func (w Walking) Walk()  {
	fmt.Println("I can walk")
}

type Human struct {
	Walking
}

type Bird struct {
	Walking
	Flying
}

func main()  {
	h := Human{}
	b := Bird{}
	h.Walk()
	b.Fly()
	b.Walk()
}