package main

import "fmt"

// GO错误处理
//go语言通过 内置 的 错误接口 提供了非常简单的错误处理机制
//
//error类型 是一个 接口类型
//
// 内置的错误接口
//type error interface {
//	Error() string
//}


type DivideError struct {
	dividee int
	diveder int
}


func (de *DivideError) Error() string {
	strFormat := "cannot proceed, the divider & zero\n dividee: %d \n divider: 0 "

	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string){
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			diveder: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10  = ", result)
	}

	if _,errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is : ", errorMsg)
	}
}