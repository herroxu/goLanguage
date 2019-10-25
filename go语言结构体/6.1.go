package main

import "fmt"

// 使用new或者&，产生的都是类型指针

type Point struct {
	x int
	y int
}

func m()  {
	point := new(Point) // point为指针类型
	point.y = 10
	(*point).x = 20
	fmt.Println((*point).x)
	fmt.Println(point.x)

	var p Point
	p.x = 1
	p.y = 2
	fmt.Println(p.x)

	// 返回结构体指针类型, new的参数只能是类型，不能进行初始化
	p1 := new(Point)
	fmt.Println(p1)

	p2 := &Point{x:2,y:5}
	fmt.Println(p2)

	//
	a := new([3]int)
	fmt.Println(a)

	b := &[]int{1,2,3}
	fmt.Println(b)
	i := 3

	cmd := newCommand("mng", &i, "dong")
	fmt.Println(cmd)

}

type command struct {
	Name string
	Var *int
	version string
}

func newCommand(name string, varlef *int, version string) *command {
	return &command{
		Name:   name,
		Var:    varlef,
		version: version,
	}
}


