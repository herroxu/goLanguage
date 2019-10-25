package main

import (
	"fmt"
	"sync"
	"os"
)

var (
	valueByKey = make(map[string]int)
	valueByKeyGuard sync.Mutex
)

func readValue(key string) int {
	valueByKeyGuard.Lock()
	//v := valueByKey[key]
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

func main()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	// sync.Mutex
	valueByKey["rmb"] = 23
	fmt.Println(readValue("rmb"))
}

func fileSize(filename string) int64  {
	// 文件句柄
	f, err := os.Open(filename)

	if err != nil {
		return 0
	}
	//文件状态信息
	info, err := f.Stat()

	if err != nil {
		return 0
	}
	// 文件大小
	size := info.Size()

	f.Close()
	return size
}

func simple(filename string) int64 {
	f, err := os.Open(filename)

	if err != nil {
		return 0
	}
	defer f.Close()

	//文件状态信息
	info, err := f.Stat()

	if err != nil {
		return 0
	}
	// 文件大小
	size := info.Size()

	// return之前，defer机制触发，关闭文件
	return size

}
