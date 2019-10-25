package main

import (
	"fmt"
	"os"
)

type consoleWriter struct {

}

func (f *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	// os.Stdout, os.Stderr, os.Stdin, 都是*os.File类型
	_, err := os.Stdout.Write([]byte(str))
	//
	return err
}

func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}

