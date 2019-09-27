package main

import "fmt"

// Go语言的For循环有3钟形式， 只有其中的一种使用分号

func main() {
	var b int = 15
	var a int
	numbers := [6]string{"a", "b", "c"}

	for a:=0; a<10; a++ {
		fmt.Println(a)
	}

	for a<b {
		a++
		fmt.Println(a)
	}

	// for循环的range格式可以对slice、map, 数组，字符串等进行迭代循环。格式如下
	for i,x := range numbers {
		fmt.Println(i, x)
	}

	i, j := 2,2

	for i=2; i<=100; i++ {
		for j=2; j<=(i/j); j++ {
			 if i%j==0 {
				break
			}
		}
		if j > (i/j) {
			fmt.Printf("%d  是素数\n",i)
		}
	}

	//乘法口诀表
	for i=1; i<=9; i++ {
		for j=1; j<=i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
