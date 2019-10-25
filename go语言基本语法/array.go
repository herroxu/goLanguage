package main

import "fmt"

type intSlice []int


func (p intSlice) Len() (l int) {
	l = len(p)
	return
}

func main()  {
	arr := intSlice{1,2,3,4,5}
	fmt.Println(arr.Len())
}