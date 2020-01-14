package main

import (
	"fmt"
)

// 使用类型分支，判断接口类型
// 结构体和接口 混合使用
type Alipay struct {

}

type Cash struct {

}

func (a *Alipay) Safe() {

}

func (a *Alipay) Unsafe() {

}

func (c *Cash) Unsafe() {

}

type SafePay interface {
	Safe()
}

type UnsafePay interface {
	Unsafe()
}

func printPay(v interface{}) {
	switch v.(type) {
	case UnsafePay:
		fmt.Println("unsafe pay")
	case SafePay:
		fmt.Println("safe pay")
	}
}

func main()  {
	printPay(&Alipay{})
	printPay(&Cash{})
}
// interface类型，error是接口
// type error interface {
//		Error() string
// }

// package fmt
// import error
//func Errorf(format string, args ... interface{}) error {
//	return errors.New(fmt.Sprintf((format, args)))
//}


