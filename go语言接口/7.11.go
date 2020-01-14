package main

import "fmt"

// go语言 类型分支， switch判断空接口中，变量类型
//type-switch 和 switch-case
//变量t 得到 areaIntf 的值和类型
//
//类型断言的 书写格式
//switch 接口变量.(type) {
//case 类型1：

//case 类型2：

//case 类型3：

//default：
//}

// 使用类型分支，判断基本类型
func printType(v interface{})  {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	// 结构体不属于基本类型，不是使用这种用法
	case struct{}:
		fmt.Println("struct")
	default:
		fmt.Println("Unknown")
	}
}

type dict struct {

}

func main()  {
	printType(1025)
	printType("ming")
	printType(false)
	printType(dict{})
	printType(123.32)
}

