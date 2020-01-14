package main

import (
	"fmt"
	"mooc"
	"time"
)

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("https://www.iyi.com/")
}

// 接口的组合
type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) string {
	poster.Post("http://www.imooc.com", map[string]string{
		"name": "ming",
		"course": "golang",
	})
}

// 接口的组合
type RetriverPost interface {
	Retriver
	Poster
}

// go语言标准接口三个


func main() {
	// 由使用者定义接口
	var r Retriver
	// 接口变量自带指针， 几乎不使用接口的指针
	// head为interface类型，取其中的值 head.(int)
	// r还没有定义
	// r开始定义
	r = mooc.Retriever{"FireFox", time.Minute}
	// 接口的 值 和 类型
	fmt.Printf("%T, %v\n", r, r)
	//
	re := mooc.Retri{"ming"}
	fmt.Printf("%T, %v\n", re, re)
	fmt.Println("=============")

	// Type assertion
	s1, ok := r.(mooc.Retri)
	fmt.Println(s1, ok)

	fmt.Println("=============")
	inspect(r)
	inspect(re)
	//fmt.Println(download(r))
}

func inspect(v interface{})  {
	switch v := v.(type) {
	case mooc.Retriever:
		fmt.Println("Retriever", v.UserAgent)
	case mooc.Retri:
		fmt.Println("Retri", v.Contents)
	}
}