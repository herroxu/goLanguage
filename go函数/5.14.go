package main

import (
	"math"
	"fmt"
	"reflect"
	"runtime"
)
// 可变参数, 当数组用

func sum(nums...int) int {
	sum := 0
	for _,v := range nums {
		sum += v
	}
	return sum
}
// 自定义类型，引用传递, go语言都是值传递，但是可以通过指针，来间接实现 引用传递，

// 函数式编程
func apply(op func(a, b int)int, a, b int) int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println(opName)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(sum(1,2,3,4,5))
}
