package main

import "fmt"

// map是一种无序的键值对的集合。Map最重要的一点是通过key值可以快速的实现数据的检索，

//定义map，可以通过make函数，也可以通过map关键字, 默认的map是nil
//第一种
//var map_variable map[key_data_type]value_data_type
//第二钟
//map_variable := make(map[key_data_type]value_data_type)


func main() {

	var country map[string]string
	country = make(map[string]string)

	country1 := make(map[string]string)

	country["china"] = "beijing"
	country["japan"] = "toyto"
	country["english"] = "london"

	for cou,_ := range country {
		fmt.Printf("%s----->%s\n\n", cou, country[cou])
	}

	fmt.Println(country1)

	// 查看元素在集合中是否存在
	capt,ok := country["china"]
	if (ok) {
		// beijing， true
		fmt.Println(capt, ok)
	} else {
		// map[],  false
		fmt.Println(capt, ok)
	}

	//delete() 用于删除map中的元素，参数为map和其对应的key值
	delete(country,"china")
	fmt.Println(country)

	var a string="cas"
	var b int=234567
	fmt.Println(a[2], b>>4)
}











