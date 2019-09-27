package main

import "fmt"

// go语言中变量可以在三个地方声明
// 函数内定义的变量称为局部变量
//函数外定义的变量称为全局变量
//函数定义中的变量称为形式参数
//go语言中全局变量可以和局部变量同名，但是局部变量的优先级更高

//形式参数会作为局部变量来使用
//for循环中，初始化的a, a:=0 与定义的a不同, 如果初始化时是a=0, 则都是一样的


// go语言数组：相同的唯一类型，且长度固定， 数组中的元素可以通过索引来修改或读取，索引从0开始
//数组的声明
var aArray [10] int

//如果不声明数组的大小，[] 默认为元素的个数来设置数组的大小
//数组的初始化
var balance = [] int {1,2,3,4,5,6,7,8,9,10}
//变量初始化赋值与数组初始化赋值 = 号的位置不一样
var salary int = balance[5]

//多维数组，声明方式
//var variable_name [size][size][size] variable_type
var triple = [1][2][3] int {
	{
		{1,2,3},
		{4,5,6}}}

func main() {
	a := aArray[0]
	b := len(balance)

	balance[5] = 12
	fmt.Println(a, b)
	fmt.Println(balance, salary)
	//访问数组元素
	for i:=0; i<=9; i++{
		aArray[i]=i
	}
	fmt.Println(aArray)
	fmt.Println(triple[0][1][2])
	// fmt.Printf(), 格式化打印

	ba := balance[2:5]
	fmt.Println(ba)
	ba = append(ba, 2222,)
	fmt.Println(234>>4)
	fmt.Println(235>>4)
	fmt.Println(236>>4)
	fmt.Println(237>>4)
	fmt.Println(230>>4)

}
