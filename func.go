package main

// 函数声明告诉了编译器函数的名称，返回类型，参数
// go语言的标准库提供了多种可动用的内置函数，len()
import (
	"fmt"
	"math"
)

// 函数声明, 函数的参数默认使用的是值传递，这样不会影响实参的值
func max(num1 string, num2 int) int {
	result := num1
	if len(result) >= num2 {
		return len(result)
	} else {
		return num2
	}
}

// 定义一个结构体类型和该类型的方法
type Circle struct {
	radius float64
}

// 定义该结构体的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}


// 声明一个函数类型
type cb func(int) int

// 写一个函数的闭包，闭包的目的是可以共享某些变量
func getsecquence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}


func main() {
	var a string="abcd"
	var b int=3
	var ret int
	var c1 Circle
	c1.radius = 10.00

	fmt.Println(c1.getArea())

	nextNumber := getsecquence()
	for i:=1; i<10; i++ {
		fmt.Println(nextNumber())
	}

	ret = max(a, b)
	fmt.Println(ret)

	getsqurtroot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getsqurtroot(19))


	//
	var dong, ge int
	dong = testCallBack(3, callback)
	ge = testCallBack(6, func(x int) int {
		fmt.Println("我不是潘金莲")
		return x
	})
	fmt.Println(dong, ge)

	add_fun := add(1,3)
	fmt.Println(add_fun())
	fmt.Println(add_fun())
	fmt.Println(add_fun())

	add_f := add_p(3,4)
	fmt.Println(add_f(5,6))
	fmt.Println(add_f(3,4))
	fmt.Println(add_f(1,2))

}

func testCallBack(x int, f cb) int {
	return f(x)
}

func callback(x int) int {
	fmt.Println("我是回调函数")
	return x
}

// 闭包不带参数
func add(x, y int) func() (int, int) {
	i := 0
	return func() (int, int) {
			i++
			return i,x+y
	}
}

// 闭包带参数补充
func add_p(x, y int) func(i, j int) (int, int, int){
	c := 0
	return func(i, j int) (int, int, int){
		c++
		return c, x+y, i+j
	}
}

// go语言函数方法
// go语言函数作为实参
// go
