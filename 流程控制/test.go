package main

import "fmt"

func main()  {
	mapStrInt := map[string]int{"ming":12, "zhong":23}
	mapIntStr := invertStringIntMap(mapStrInt)
	fmt.Println(mapIntStr)
	fmt.Println(len(fmt.Sprint(234343)))
}

func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key, "1")
		fmt.Println(stringsForInt)
	}
	return stringsForInt
}
