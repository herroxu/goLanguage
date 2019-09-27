package main

import (
	"errors"
	"fmt"
)

// GO错误处理
//go语言通过 内置 的 错误接口 提供了非常简单的错误处理机制
//
//error类型 是一个 接口类型
//
// 内置的错误接口
//type error interface {
//	Error() string
//}

func sqrt( n float64) (float64, error) {
	if n <0 {
		return 0, errors.New("math is not 0")
	} else {
		return 23,errors.New("")
	}
}

type DivideError struct {
	devidee int
	devider int
}


func (de *DivideError) Error() string {
	strFormat := "cannot proceed, %d"
	return fmt.Sprintf(strFormat, de.devidee)
}


func main() {
	a := 2.0123
	result,err := sqrt(a)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)



}
