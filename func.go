// HTTP文件服务器， web常见的服务器之一

package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", nil)
}
/*
工作区
src目录
	以代码包的形式组织并保存Go源码文件
pkg目录
	存放通过go install命令安装后的代码包的开档文件
bin目录
	存放通过go install命令安装后，由Go命令源码文件生成的可执行文件

命令源码文件
库源码文件
测试源码文件
	1。文件名以 "_test.go"结尾
	2。文件中至少包含一个名称以Test开头或Benchmark开头，且拥有一个类型为*testing.T或者*testing.B的函数testing.T是一个结构体


*/
// 包初始化
//1。先对全局变量初始化
//2。然后init()
//3. 然后main()
func init() {
	i, j := 1, 4
	// 两值交换
	i, j = j, i
}
//静态类型语言
//var a,b *int
//var (
//	c int
//	d string
//	e []float64
//	f func() bool
//	g struct{
//		x int
//	  }
//)
//var hp int = 100
//var hp = 100

