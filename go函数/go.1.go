package main

import "fmt"

// 函数式编程, 返回值为一个函数
func addr() func(value int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

func main()  {
	add := addr()
	fmt.Println(add(1))
	fmt.Println(add(2))
	fmt.Println(add(3))
	fmt.Println(add(4))
}
