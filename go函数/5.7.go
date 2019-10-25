package main

import "fmt"

// 函数类型实现接口， 把函数作为接口来调用

type voker interface {
	call(interface{})
}

type s struct {

}

func (s s) call(interface{})  {
	fmt.Println("from struct")
}

func (s s) receive(interface{}) {
	fmt.Println("receive struct")
}

type funcin func(interface{})

func (f funcin) call(p interface{})  {
	f(p)
}



func main()  {
	var v, voke voker
	//stru := new(s)  //new()将结构体实例化,返回结构体的指针
	//stru := &s
	stru := s{}
	//fmt.Println("%p", *stru)
	v = stru
	v.call("hello")
	// 方法中需要的是指针，传结构体实例不行，1、实例化加地址符  2、使用new()
	// 方法中需要的是实例对象时，可以传指针，实例对象都行

	voke = funcin(func(v interface{}) {
		fmt.Println(v)
	})
	voke.call("hellp")
}

