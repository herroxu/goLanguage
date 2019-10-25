package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "skill usage")

func main()  {
	flag.Parse()
	fu := map[string]func(){
		"fly": func() {
			fmt.Println("fly")
		},
		"blue": func() {
			fmt.Println("blue")
		},
		"red": func() {
			fmt.Println("red")
		},
	}
	if f,ok := fu[*skillParam]; ok {
		f()
	} else {
		fmt.Println("not found")
	}
}



