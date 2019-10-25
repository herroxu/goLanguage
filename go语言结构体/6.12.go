package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"reflect"
)

type Screen struct {
	Size float32
	Rex, Rey int
}

type Battery struct {
	Capacity int
}

func genJsonData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchBar bool
	}{
		Screen:Screen{
			Size:32.2,
			Rex:34,
			Rey:23},
		Battery:Battery{Capacity:344},
		HasTouchBar: true,
	}

	v1 := reflect.TypeOf(*raw)
	fmt.Println(v1.Kind())

	fmt.Println(raw)
	// 使用json.Marshal()进行json序列化，将raw变量序列化为[]byte格式的JSON数据
	jsonData,_ := json.Marshal(raw)

	return jsonData
}

func main()  {
	jsonData := genJsonData()
	fmt.Println(jsonData)
	fmt.Println(string(jsonData))

	//
	screenAndTouch := struct {
		HasTouchBar bool
		Screen
	}{}
	runtime.GC()
	//fmt.Printf("%d\n", runtime.MemStats.Alloc/1024)
	// 在一个对象obj被移除前执行一些特殊操作，
	//runtime.SetFinalizer(obj, func(obj *types.Object))
	//
	json.Unmarshal(jsonData, &screenAndTouch)

	fmt.Println(screenAndTouch)


}

// GC 垃圾回收器
// 通过runtime包防问GC进程




