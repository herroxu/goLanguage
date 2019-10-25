package main

import "fmt"

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct {

}

func (DummyWriter) Write(buf []byte) (int, error)  {
	return len(buf), nil
}

//func ( ) Writer(buf []byte) (int, error) {
//	return len(buf), nil
//}

func min()  {
	var x interface{} = DummyWriter{}
	//
	var y interface{} = "abc"
	//
	var w Writer
	//
	var ok bool

	//DummyWriter既实现来Writer接口，也实现interface{}
	w2,ok := x.(interface{})
	fmt.Println(w2, ok)

	w, ok = x.(Writer)
	fmt.Println(w, ok)

	// y的动态类型为string，string类型并没有实现Writer接口，所以错误
	y1,ok := y.(Writer)
	fmt.Println(y1, ok)

}