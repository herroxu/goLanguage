package main

import "fmt"

// map和interface{}, 实现了一个字典
type Dictionary struct {
	Data map[interface{}]interface{}
}

// 获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.Data[key]
}

//设置键值对
func (d *Dictionary) Set(key, value interface{}) {
	d.Data[key] = value
}

// 遍历所有的键值，使用回调函数，对键值对进行处理
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {

	if callback == nil {
		return
	}
	fmt.Println(d.Data)
	for k,v := range d.Data {
		if !callback(k, v) {
			return
		}
	}
}

// 清除数据
func (d *Dictionary) Clear() {
	d.Data = make(map[interface{}]interface{})
	return
}

// 创建一个字典

func NewDictionary() *Dictionary {
	dictionary := &Dictionary{}
	//dictionary = new(Dictionary)
	dictionary.Clear()
	return  dictionary
}

func main() {
	Nick := NewDictionary()
	Nick.Set("ming", "wang")
	Nick.Set(12,34)
	Nick.Set(false,4)

	Nick.Visit(func( k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

	fmt.Println(Nick.Get("ming"))
	fmt.Println(Nick.Get(12))
	fmt.Println(Nick.Get(false))
}







