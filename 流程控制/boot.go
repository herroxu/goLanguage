package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	readInput := bufio.NewReader(os.Stdin)
	fmt.Println("input your name")
	input,err := readInput.ReadString('\n')
	if err != nil {
		fmt.Printf("an error %s", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("Hello, %s", name)
	}

	for {
		input, err := readInput.ReadString('\n')
		if err != nil {
			fmt.Printf("an error %s", err)
			continue
		}
		message := input[:len(input)-1]
		message = strings.ToLower(message)

		switch message {
		case "":
			continue
		case "bye", "no":
			fmt.Println("bye")
			os.Exit(0)
		default:
			fmt.Println("sorry")
		}

	}
}



