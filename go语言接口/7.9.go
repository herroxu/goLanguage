package main

import "fmt"

// go语言 空接口类型 interface{}

//空接口 是 接口类型 的 特殊形式，空接口没有任何方法，因此任何类型，都无需实现空接口，从实现的角度看，任何值都满足这个接口的需求，
//因此空接口类型，可以保存任何值，也可以从空接口中取出 原值
func main()  {
	var any interface{}
	any = 1
	fmt.Println(any)
	any = "hello"
	// any内部保存来一个字符串，但类型依然是interface{}
	fmt.Println(any)
	any = false
	fmt.Println(any)

	var a int = 1
	var i interface{} = a
	var b int = i.(int) // i.(Type)
	fmt.Println(b)
	// 存在空接口中的值 比较
	// 类型不同 的 空接口 间的 比较 结果 不相同， 类型不同必定不相等
	//  不能比较空接口中的动态值
	// 当接口中保存有动态类型的值时， 运行时触发错误
	// map和切片 是动态类型， 不可比较


}





