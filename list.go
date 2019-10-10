package main

import (
	"container/list"
	"fmt"
)

type error interface {
	Error() string
}

type Error struct {
	errcode uint64
}

// 指针，函数，interface， slice，channel，map的令旨都是nil
func (e *Error) Error() string {
	switch e.errcode {
	case 1:
		return "file not found"
	case 2:
		return "time out"
	case 3:
		return "permission denied"
	default:
		return "unknown"
	}
}
func checkError(err error)  {
	if err == nil {
		panic(err)
	}
}

func main() {
	// 双链表
	a := list.New()
	var e *Error

	checkError(e)

	//var b list.List
	a.PushFront("back")
	a.PushBack(1234)
	el := a.PushBack("no")
	a.InsertAfter("high", el)
	a.InsertBefore("noon", el)
	//a.Remove(el)
	for i:=a.Front(); i != nil; i=i.Next() {
		fmt.Println(i.Value)
	}
}
/* make和new关键字的区别

make关键字的主要作用是初始化内置的数据结构，数组，切片，channel，
而当我们想要获取指向某个类型的指针时，可是使用new关键字，
make在语言中只能用于初始化语言中的基本类型
make([]int, 0, 100) -- 结构体
make(map[int]int, 10) -- 指向map结构体的指针
make(chan int, 5)  -- 指向hchan结构体的指针

new(int)  -- 返回一个指向这个类型的指针


*/