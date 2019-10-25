package main

import "fmt"

// go方法是作用在接收器上的一个函数，接收器是某种类型的变量，方法是一种特殊类型的函数
// 除了接口类型， 接口是一个抽象的定义，而方法确是具体实现, invalid receiver


type str string
// 指针接收器 和 非指针接收器
func (s str) call() str  {
	st := s + "欢迎来到英雄联盟"
	s = st
	return st
}


func man()  {
	var s str
	s = "明仔 "
	s1 := s.call()
	fmt.Println(s1)
	fmt.Println(s)
}