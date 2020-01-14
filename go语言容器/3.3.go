package main

import "fmt"

func main()  {
	// 建立map
	m := map[string]string{
		"python": "1",
		"go": "2",
		"c++": "3",
	}

	m2 := make(map[string]int) // 实例
	// map 是无序的
	fmt.Println(m, m2)
	// 遍历
	for k,v := range m {
		fmt.Println(k, v)
	}
	// 取值， 取不到则取默认值，不报错
	order := m["python"]
	or, ok := m["c++"]
	fmt.Println(order, or, ok)
	// delete key-value
	delete(m, "go")
	fmt.Println(order, or, ok, m)

	// map的key，必须可以比较相等
	// 除了slice， map，function 的内建类型都可以作为key
	// struct类型不包括上述字段，也可以作为key

}
