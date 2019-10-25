package main

import "fmt"

type Data struct {
	complax []int
	instance InnerData
	ptr *InnerData
}


type InnerData struct {
	a int
}

func PassByValue(inner Data) Data {
	fmt.Printf("inner data, %+v\n", inner)
	fmt.Printf("inner data, %p\n", &inner)

	//inner.ptr = 3

	return inner
}

func main()  {
	in := Data{complax:[]int{1,2,3}, instance:InnerData{a:2}, ptr:&InnerData{a:23}}
	fmt.Println(in.instance.a)

	fmt.Printf("inner data, %+v\n", in)
	fmt.Printf("inner data, %p\n", &in)

	out := PassByValue(in)

	fmt.Println(*out.ptr)
	fmt.Println(*in.ptr)


	fmt.Printf("inner data, %+v\n", out)
	fmt.Printf("inner data, %p\n", &out)
}

