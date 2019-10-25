package main

import "fmt"

// 内嵌结构体成员名字冲突

type Wheel struct {
	size float32
}

type Brand struct {
	name string
}

type Car struct {
	Wheel
	Brand
	Engine struct{
		Power int
		Type string
	}
}

func main()  {
	car := Car{
		Wheel{size:32},
		Brand{"ben chi"},
		struct {
			Power int
			Type  string
		}{Power: 2000 , Type: "de guo" },
	}
	fmt.Println(car)
}