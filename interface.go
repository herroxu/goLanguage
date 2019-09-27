package main

import "fmt"

// go语言提供了另外一种数据类型， 接口， 它把所有具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现 这个接口
// 接口类型只能与结构体类型共同使用

type Phone interface {
	call() string
	receive() string
}

type phone11 struct {
	cpu string
	capacity int
}

type huawei struct {
	cpu string
	capacity int
}

func (phone11 phone11) call() string {
	return "我后台是苹果"
}

func (phone11 *phone11) receive() string {
	phone11.cpu = "gao tong"
	return "干的就是你苹果"
}

func (huawei huawei) call() string {
	return "我的后台是华为"
}

func (phone11 huawei) receive() string {

	return "干的就是你华为"
}

func main() {
	var phone Phone
	// phone11是结构体
	phone1 := phone11{"intel", 16}

	fmt.Println(phone1.call())
	fmt.Println(phone1.receive())

	fmt.Println(phone1)

	phone = new(phone11)
	fmt.Println(phone.call())
	fmt.Println(phone.receive())

	phone = new(huawei)
	fmt.Println(phone.call())
	fmt.Println(phone.receive())

}

