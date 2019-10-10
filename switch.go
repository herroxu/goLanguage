package main

import "fmt"

var a []int = []int {1,2,3,4}

// make() 构造切片
// make([]type, size, cap)
// append()  内建函数，为切片动态添加元素
// copy() 内建函数， range关键词

func main()  {
	b := []int{7,8,9}
	var c [3]int
	c = [3]int{2,3,4}
	a = append(a, 5)
	a = append(a, 6, 7, 8, 9)
	a = append(a, []int{10, 11, 12, 13}...) //...使用来将切片解包

	fmt.Println(a, b, c)
}