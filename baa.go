package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	const size = 300

	pic := image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x<size; x++ {
		for y:=0; y<size;y++ {
			pic.SetGray(x, y, color.Gray{ 255})
		}
		pic.SetGray(x,45, color.Gray{23})
		pic.SetGray(x,65, color.Gray{23})

	}
	// 字符串可以使用索引str[1]
	// 定义多行字符串使用反引号

	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, pic)
	file.Close()
}