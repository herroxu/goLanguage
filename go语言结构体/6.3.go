package main

import "fmt"

// 结构体的成员变量初始化，两种方式，1、键值对 2、多个值列表

type People struct {
	name string
	child *People
}

func min()  {
	// 使用键值对初始化结构体成员变量
	relation := People{
		name:  "father",
		child: &People{
			name:  "son",
			child: &People{
				name:  "grandson",
				child: nil,
			},
		},
	}
	fmt.Println(relation.child.child.name)
	
	// 使用多个值的列表初始化结构体
	my := People{"ming", nil}
	fmt.Println(my)
	
	//匿名结构体
	msg := &struct {
		id int
		data string
	}{
		1024,
		"hello",
	}
	printMsg(msg)
	fmt.Printf("%T\n", my)
	fmt.Printf("%T", relation)

}

func printMsg(msg *struct{
	id int
	data string}) {
	fmt.Printf("%T\n", msg)

}
