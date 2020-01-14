package main

import "fmt"

func main()  {
	var arr1 [5]int
	arr2 := [6]int{11,22,33,44,55,66} // 使用:=定义时，一定要实例化
	arr3 := [...]int{1,2,3} // 使用...表示不定长数组，由后面的个数决定， 切片
	var grid [4][5]int
	// 数组在函数中作为参数传递时， 传递是值传递， 数组是值类型， 数组指针当数组实例用

	fmt.Println(arr1, arr2, arr3, grid)
	for i,v := range arr3 {
		fmt.Println(i ,v)
	}
	// Slice切片
	arr4 := [...]int{1,2,3,4,5,6}
	s1 := arr4[2:4]
	s2 := arr2[2:4]
	s1[1] = 1000
	fmt.Println(s1, s2, arr4)

}
// 切片从数组中切出来，然后修改切片中的值，数组的值也会改变, 切片是引用类型
// 没有指定长度，就是切片类型，切片是引用类型， 把数组变成切片，arr1[:]
// slice 本身没有数据，是对底层array的一个view，且是一个引用类型
// slice添加数组时， 如果没有超过array的cap，则会修改array，如果超过，系统重新分配更大的底层数组
func printArray(arr []int) {

}
