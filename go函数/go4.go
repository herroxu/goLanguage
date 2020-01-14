package main

import (
	"fmt"
	"errors"
)

func tryDefer()  {
	// 实质就是一个栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println(4)
}

func tryRecover()  {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred", err)
		} else {
			panic(r)
		}
	}()
	panic(errors.New("an error"))
}

func main()  {
	tryDefer()
	tryRecover()
}



