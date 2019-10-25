package main

import (
	"runtime"
	"fmt"
)

type panicContext struct {
	function string
}

func ProtectRun(entry func())  {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		case *panicContext:
			fmt.Println("error:", err)
		}
	}()

	entry()
}

func main()  {
	fmt.Println("运行前")
	ProtectRun(func() {
		fmt.Println("front")

		panic(&panicContext{"手动错误"})

		fmt.Println("back")
	})

	ProtectRun(func() {
		fmt.Println("front 1")
		var a int
		a = 1
		fmt.Println("back 1", a)
	})
	fmt.Println("运行后")
}



