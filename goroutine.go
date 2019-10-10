package main
//
//import (
//	"fmt"
//	"math/rand"
//	"strconv"
//	"time"
//)
//
//// 标准库包名
///*
//bufio    带缓冲的i/o操作
//bytes	字节操作
//container	封装堆，列表和环形列表等容器
//crypto	加密算法
//database	数据库驱动和接口
//debug	调试功能
//encoding	实现算法如Json,XML, Base64等
//flag	命令行解析
//fmt 	格式化操作
//go		go语言的词法，语法树。类型等
//html	HTML转义，模版系统
//image   常见图形格式的访问及生成
//io		实现io访问接口及封装
//math	数学库
//net		网络库，支持Socket, HTTP, 邮件， RPC， SMTP
//os		操作系统平台，不依赖平台操作封装
//path
//plugin
//reflect
//regexp		正则表达式封装
//runtime
//sort		排序接口
//strings		字符串转换/解析， 实用函数
//time		时间接口
//text		文本模版
//strconv
//
//*/
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
////chan<- string 表示声明一个只能接收字符串的通道
//func productor(head string, channel chan<- string) {
//	for {
//		channel <- fmt.Sprintf("%s:%v", head, rand.Int31())
//		time.Sleep(time.Second)
//	}
//}
//
//// <-chan 表示只能用来输出的通道
//func customer(channel1 <-chan string) {
//	for {
//		//
//		message := <-channel1
//		fmt.Println(message)
//	}
//}
//
//func main() {
//	// 创建一个通道，只能传输字符串
//	channel1 := make(chan string)
//
//	// 调用生产者函数
//	go productor("cat", channel1)
//	go productor("dog", channel1)
//
//	strconv.Atoi() // string -> int
//	strconv.Itoa() // int -> string
//	strconv.ParseBool(), strconv.ParseFloat() .......
//	strconv.FormatUint()
//	strconv.AppendFloat()
//	customer(channel1)
//}