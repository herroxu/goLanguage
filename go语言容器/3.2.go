package main

import "fmt"

func main()  {
	var s []int // 声明一个slice, s == nil

	for  i:=0; i<100; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2,4,6,8}
	s2 := make([]int, 10)
	s3 := make([]int, 16, 32)
	copy(s2, s1)
	// delete
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s1,s2,s3)
}
