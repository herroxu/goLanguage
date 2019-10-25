package main

import (
	"os"
	"errors"
	"fmt"
)

type fileWriter struct {
	file *os.File
}

func (f *fileWriter) SetFile(filename string) (err error)  {
	//
	if f.file != nil {
		f.file.Close()
	}
	// 创建一个文件，并保存文件句柄
	f.file, err = os.Create(filename)
	//
	return nil
}

func (f *fileWriter) Write(data interface{}) error {
	//
	if f.file == nil {
		return errors.New("file not created")
	}

	str := fmt.Sprintf("%v\n", data)

	// 将数据以字节数组写入文件中
	_,err := f.file.Write([]byte(str))

	return err
}

func newFileWriter() *fileWriter  {
	return &fileWriter{}
}




