package main

import "fmt"

func accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func player(name string,) func() (string, int)  {
	hp := 150
	// 返回一个闭包
	return func() (string,  int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

func main()  {

	player1 := player("zhong")
	name, hp := player1()
	fmt.Println(name, hp)

	accumulator := accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Printf("%p\n",accumulator)

	fmt.Println("=====")
	accumulator2 := accumulate(1111)

	fmt.Println(accumulator2())
	fmt.Println(accumulator2())
	fmt.Println(accumulator2())
	fmt.Println(accumulator2())
	fmt.Println(accumulator2())
	fmt.Printf("%p\n",accumulator2)
	fmt.Println(accumulator())

}

