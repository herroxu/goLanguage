package main

import "fmt"

// 递归函数就是自己调用自己, 但是需要设置退出条件，否则将陷入无限循环中

// 使用递归函数实现阶乘

func recursion(a uint64) uint64 {
	var result uint64 = 1
	if a != 1 {
		result *= (recursion(a-1)*a)
		return result
	} else {
		return  result
	}
}

// GO语言类型转换，int(), int64(), float64()

// GO语言提供了一种特殊的数据类型，称为接口，

func main() {
	var a uint64 = 4
	fmt.Println(recursion(a))
}