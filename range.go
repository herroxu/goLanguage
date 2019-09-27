package main

import "fmt"

//range关键字用于for循环中迭代数组(array, slice, channel, mao)中的元素，在数组和切片中它返回元素的索引和索引对应的值，
//在集合中返回key-value对的key值
func main()  {
	nums := []int{1,2,3,4}
	sum := 0
	// 在切片上使用range
	for _,num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	// 在数组上使用range
	for i,num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}

	// range也可以用在map的键值对上
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k,v := range kvs {
		fmt.Printf("%s - > %s\n", k, v)
	}
	// range使用在字符串上
	for i,c := range "GO" {
		fmt.Println(i,c)
	}

}