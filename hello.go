// 包声明， 源文件中非注释的第一行指明这个文件属于哪个包， package main表示一个可独立执行的程序，每一个go的in 应用程序， 都有一个名为main的包
// 表明这个文件属于哪个包
package main

// 引入包， 告诉Go编译器， 需要使用fmt这个包， 包含栏IO函数
import (
	"fmt"
)

// go语言常量
const aa string = "italki"
const bb = "developer"






// 函数

//变量， 变量声明必须使用空格隔开
//var age int

// go语言数据类型
//布尔型 var b bool = true
//数字类型 int float32 float64
//字符串类型
//派生类型：1。指针类型 2。数组类型 3，结构化类型 4。channel类型 5。map类型 6。函数类型 7。接口类型 8。切片类型
//数字类型 uint8(0~255), uint16(0~65535), int(-128~127)
//浮点类型 float32

//一次声明多个变量, 格式
//var identifier1, identifier2 type


//语句， 表达式， 注释


// func main()是程序开始执行的函数， 一般是第一个执行的， 除非有init()函数
// { 不能单独在一行
func main() {
	var aStr string = " xu zhi "
	var bStr string
	var cStr string
	var dStr bool
	// 字符串未初始化就是空字符串
	// 数值类型未初始化为0
	// 布尔型未初始化就是false
	// 剩下的几种未初始化就是nil,
	// 指针
	var a *int
	// 数组
	var b []int
	// map
	var c map[string] int
	// channel
	var d chan int
	// 函数
	var e func(string) int
	// 接口
	var f error

	// 第二种， 根据指自行判定变量类型
	//var v_name = v_value
	//var d = true


	// 第三种
	aInt,bInt := 2, 4

	// 交换两个变量的值, 但是必须得是同类型
	_, aInt = aInt, bInt

	// 多变量的声明
	var astr, bstr, cstr string
	astr, bstr, cstr = "dong", "fang", "bu"

	// := 这种简单的赋值形式，只能在函数体重使用，不可以用于全局变量的声明和赋值， 如果局部变量未被使用，编译会报错。
	// 但是全局变量是可以声明但是不使用的
	aint, bint := 4,6

	var eint, fstr = 45, "bai"

	// 值类型和引用类型

	// 这个内存地址被称为指针， 当r1 = r2的时候，只是引用地址被复制

	bStr = "dong"
	fmt.Println("Hello, World")
	// go语言的字符串可以使用 + 来连接
	fmt.Println("google" + aStr +"ming")
	fmt.Println(bStr+cStr)
	fmt.Println(dStr)
	fmt.Println(a, b, c, d, e, f, aInt, bInt)
	fmt.Println(astr, bstr, cstr, fstr)
	fmt.Println(aint, bint, eint)
	fmt.Println(&aint)
	//fmt.Println(mathclass.ADD(1,1))
	//fmt.Println(mathclass.Sub(2,1))
	a_int, b_int, _ := numbers()
	fmt.Println(a_int, b_int)
}

func numbers()(int, int, string){
	a, b, c := 1,2,"ming"
	return a,b,c
}

// fmt.Println() 将字符串输入到控制台，自动增加换行符， fmt.Print(), 没有换行符， 也支持使用变量

// 当标识符以一个大写字母开头，那么这个标识符就可以被外部对象所使用， 如果是小写，就不可以被外部导入使用， 但是在包内都可以使用

//go 语言中的25个关键字或者保留字， 还有36个预定义标识符


