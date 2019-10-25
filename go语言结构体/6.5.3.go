package main

import "fmt"

func mn()  {
	p := NewPlayer(3)

	p.MoveTo(Vec2{4,4})

	for !p.IsArrived() {
		p.update()
		fmt.Println(p.Pos())
	}
}
