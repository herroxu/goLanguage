package main

import "fmt"

//切片是动态数组，切片的长度是不确定的，可以追加元素
// 切片不需要说明长度
//  未指定大小的数组为切片
//  var identifier [] type
//或者使用make函数来创建切片, 内置类型 切片

//  var slice1 []type = make([]type, len)
//  间写: capacity为可选参数
//  slice1 := make([]type, len, capacity)

// 切片初始化
var arr = [] int {11,22,33,44,55,66,77,88}


func main() {
	s := [] int {1,2,3}
	s1 := arr[2: 5]
	s2 := arr[2: ]
	s3 := arr[ :5]
	s4 := arr[ : ]

	// len() 和 cap() 函数
	//切片是可以索引的，可以有len()方法获取长度
	//切片提供了计算容量的方法cap(),可以计算切片最长可以达到多少
	var numbers = make([]int, 3,5)
	printSlice(numbers)

	fmt.Println(s, s1, s2, s3, s4)

	//append() , copy() 函数
	//拷贝切片的代码copy()，向切片追加新元素的append()方法
	var num []int
	printSlice(num)

	num = append(num, 0)
	printSlice(num)

	num = append(num, 11,22,33)
	printSlice(num)

	num1 := make([]int, len(num), (cap(num))*2)
	printSlice(num1)

	//把num的内容拷贝至num1中
	copy(num1, num)
	printSlice(num1)

}


func printSlice(x []int) {
	fmt.Printf("len: %d, cap: %d, array:%d \n\n",len(x), cap(x), x)
}














