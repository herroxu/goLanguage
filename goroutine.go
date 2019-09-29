package main

import (
	"fmt"
	"time"
)

// go并发， 通过关键字开启goroutine即可
// goroutine是轻量级线程， 语法
// go 函数名（ 参数列表 ， 同一个程序中的所有的goroutine共享同一个地址空间

//使用了goroutine的函数，再次调用，传不同的参数时，相当于又调用一个goroutine


func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}
// 通道channel是一个用来传递数据的一个数据结构，可用于两个goroutine之间通过传递一个指定类型的值
// 来同步运行和通讯，操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道
// ch <- v 把 v 发送到通道 ch
// v := <- ch 从通道ch 接收数据，并把值赋给v
// ch := make(chan int)

func sum(s []int, c chan int) {
	sum := 0
	for _,v := range s {
		sum += v
	}
	c <- sum
}


func main() {
	go say("中国")
	say("中国")

	s := []int{5,34,23,45,75,23,42,25}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
	//ch 通道的缓存区的大小为2
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//v, ok := <-ch, 如果通道接收不到数据后OK就为false，这是通道就可以使用close（）函数来关闭
	c1 := make(chan int, 10)
	go fibo(cap(c1), c1)

	for i := range c1 {
		fmt.Println(i)
	}
}

func fibo(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i<n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}



