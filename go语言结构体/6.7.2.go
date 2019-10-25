package main

import "fmt"

type Actor struct {

}

func (a *Actor) OnEvent(param interface{})  {
	fmt.Println("actor event", param)
}

func GlobalEvent(param interface{})  {
	fmt.Println("global event", param)
}

func main()  {
	a := new(Actor)

	registerEvent("call", a.OnEvent)
	registerEvent("call",GlobalEvent)

	callBack("call","china")
}



