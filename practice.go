package main

import "fmt"

// 常量表达式中，函数必须是内置函数，否则编译不过
const (
	//a = "abc"
	//b = len(a)
	//c = unsafe.Sizeof(a)
	a = iota
	b = 3<<iota
	c
	j
)

//iota, 特殊常量

// go语言内置运算符：
// 算术运算符: +， -， *， /， %， ++， --
// 关系运算符：==， !=, >, <, >=, <=
// 逻辑运算符: &&, ||, !
// 位运算符: &, |, ^, <<, >>
// 赋值运算符:  =, +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, !=
// 其他运算符 &, *

func main(){
	println(a, b, c, j)
	//var d,f int
	T, F := true, false
	var d int = 34
	var ptr *int

	ptr = &d
	fmt.Println(*ptr)
	fmt.Println(d)

	if( a <= b ) {
		fmt.Println("第一行")
	} else {
		fmt.Println("第二行")
	}

	if( a>=b ) {
		fmt.Println("天下无敌")
	} else {
		fmt.Println("一统江湖")
	}
	// =================
	if( !(T&&F) ) {
		fmt.Println("天下无敌")
	} else {
		fmt.Println("一统江湖")
	}
	// ================
	if( c>a ){
		fmt.Println("我错了")
		if( c>b ) {
			fmt.Println("好")
		} else {
			fmt.Println("错")
		}
	}

	fmt.Println("大结局")

	var a_int int
	var b_int int
	fmt.Printf("请输入密码： \n")
	fmt.Scan(&a_int)

	if( a_int!=123456 ) {
		fmt.Printf("请再次输入密码：   \n")
		fmt.Scan(&b_int)

		if( b_int==123456 ) {
			fmt.Println("芝麻开门")
		} else {
			fmt.Println("芝麻中毒")
		}

	} else {
		fmt.Println("芝麻开门")
	}
}

