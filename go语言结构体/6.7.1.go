package main

// 事件系统实现 事件的响应和处理

var eventByName  = make(map[string][]func(interface{}))

func registerEvent(name string, callback func(interface{})) {
	list := eventByName[name]

	list = append(list, callback)

	eventByName[name] = list
}

func callBack(name string, param interface{})  {
	list := eventByName[name]

	for _, callback := range list {
		callback(param)
	}
}
