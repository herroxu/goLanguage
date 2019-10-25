package main

import "fmt"

// 构造函数， 结构体初始化的过程，模拟实现
// 构造函数重载
// 基类
type Cat struct {
	Name string
	Color string
}
// 继承自基类的子类
type WhiteCat struct {
	*Cat
	Age int
}
// 构造基类
func initCat(name string, color string) *Cat {
	return &Cat{name, color}
}
// 构造子类
func whiteCat(name string,age int) *WhiteCat {
	return &WhiteCat{
		Cat: initCat(name, "White"),
		Age: 4,
	}
}



func Color(cat *Cat, color string) *Cat {
	cat.Color = color
	return cat
}

func Name(cat *Cat, name string) *Cat {
	cat.Name = name
	return cat
}


func ms()  {
	cat := &Cat{}
	// 函数重载
	Color(cat, "black")
	Name(cat, "ming")
	fmt.Println(cat.Color)
	fmt.Println(cat.Name)
	china := whiteCat("china",70)
	fmt.Println(china.Name, china.Color, china.Age)

}