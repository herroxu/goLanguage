package main

import "io"

// go语言中，接口与接口之间也是可以通过嵌套创造出新的接口

type device struct {

}

func (d *device) Write(p []byte) (n int,err error){
	return 0, nil
}

func (d *device) Close() error {
	return nil
}

func main()  {
	// 接口类型，被赋予实例
	var wc io.WriteCloser = new(device)

	wc.Write(nil)

	wc.Close()

	var writeOnly io.Writer = new(device)
	writeOnly.Write(nil)
}
