package main

import (
	"fmt"
)

type ParseError struct {
	Fiename string
	Line int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Fiename, e.Line)
}

func newParseError(filename string, line int) (b *ParseError) {
	b = &ParseError{"ming", 45}
	return
}



//var err = errors.New("this is an error")
//
//func div(di, dr int) (int, error)  {
//	if dr == 0 {
//		return 0,err
//	}
//	return di/dr, nil
//}

// 接口可以匹配任何类型

func main() {
	var e error
	e = newParseError("main.go", 1)
	fmt.Println(e.Error())

	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename: %s  Line: %d\n", detail.Fiename, detail.Line)
	default:
		fmt.Println("")
	}


	//fmt.Println(div(2,0))
}
