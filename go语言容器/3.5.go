package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	s := "我是中国人"// UTF-8 可变长字节编码
	fmt.Printf("%s\n", s)
	fmt.Println(len(s), []byte(s)) // 15个字节
	for _,b := range []byte(s) {
		fmt.Printf("%X  ",b)

	}

	for i,ch := range s { // ch is rune,为四字节类型的unicode
		fmt.Printf("(%d, %X) ",i, ch)
	}
	fmt.Println()
	fmt.Println(utf8.RuneCountInString(s))

	b := []byte(s)
	for len(b) > 0 {
		ch, size := utf8.DecodeRune(b)
		fmt.Printf("%c",ch)
		fmt.Println(ch, size)

		b = b[size:]
	}
}

