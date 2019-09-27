package main

import "fmt"

func main() {
	grade, mark := "B", 23
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Println(1, i)
	case int:
		fmt.Println(2)
	case float64:
		fmt.Println(3)
	case func(int) float64:
		fmt.Println(4)
	case bool, string:
		fmt.Println(5)
	default:
		fmt.Println(6)
	}

	// fallthrough 的用法, 只会强制执行下一条的case语句，不判断是否为true， 直接执行
	switch {
	case false:
		fmt.Println("1111")
		fallthrough
	case true:
		fmt.Println("2222")
		//fallthrough
	case false:
		fmt.Println("3333")
	}

	switch mark {
	case 90, 23, 43: grade = "A"
	case 80: grade = "B"
	case 70: grade = "C"
	default:
		grade = "D"
	}

	switch {
	// case后面的逗号表示或者
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B":
		fmt.Println("良好")
	case grade == "C":
		fmt.Println("及格")
	case grade == "D":
		fmt.Println("不及格")
	default:
		fmt.Println("差")
	}
}
