package main

import "fmt"

func fire(a, _ int) (c string) {
	fmt.Println("fire", a)
	c = "ming"
	return c
}

func main() {
	// f 为函数类型变量
	var f func(int, int) string
	f = fire
	a := f(2, 3)
	fmt.Println(a)
}