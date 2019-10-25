package main

import "fmt"

// go语言中 使用 接口断言 将
// 		1、接口转换成 另外一个接口，
// 		2、也可以将接口类型 转换为 另外的类型（非接口类型）
//类型断言 格式：
//i.(T)  i表示一个接口的类型，T表示一个类型
// 一个类型断言 检测 它操作对象的动态类型是否和断言的类型匹配
// t := i.(T)
//i 表示：接口变量， t表示： 转换后的变量

// 实现某个接口的类型同时实现来另外一个接口，此时可以在两个接口间转换
type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct {

}

func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

type pig struct {

}

func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

func main()  {
	animals := map[string]interface{} {
		"bird": new(bird),
		"pig": new(pig),
	}
	//
	for name, obj := range animals {
		// 接口断言
		f, isFly := obj.(Flyer)
		w, isWalk := obj.(Walker)

		fmt.Printf("%s, %v fly, %v walk\n", name, isFly, isWalk)
		if isFly {
			f.Fly()
		}
		if isWalk {
			w.Walk()
		}

	}
}










