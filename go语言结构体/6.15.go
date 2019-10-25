package main

import (
	"fmt"
	"bytes"
	"reflect"
	"strconv"
	"errors"
)

// GO语言中，结构体 可以通过系统提供的 json.Marshal() 函数进行 序列化
// 通过 反射 获取结构体成员 及 各种值 的过程， 下面使用反射 将结构体序列化为 文本数据

func main()  {

	type Skill struct {
		Name string
		Level int
	}

	type Actor struct {
		Name string
		Age int

		skills []Skill
	}

	//
	a := Actor{
		Name:   "ming",
		Age:    20,
		skills: []Skill{
			{Name:"roll and roll", Level:1},
			{Name:"flash and flash", Level:2},
			{Name:"time to time", Level:3},
		},
	}

	if jsonData,eil := MarshJson(a); eil==nil {
		fmt.Println(string(jsonData))
		fmt.Println(jsonData)

	} else {
		fmt.Println(eil)
	}
}

func MarshJson(v interface{}) (string, error) {
	// v为一个结构体实例

	// 准备一个缓冲池
	var b bytes.Buffer

	// reflect包， Typeof(), Valueof()

	fmt.Println(reflect.ValueOf(v))
	// 反射值对象
	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {
		return b.String(),err
	} else {
		return "", nil
	}
	//
}

func writeAny(buff *bytes.Buffer, value reflect.Value) error {
	fmt.Println(value.Kind())
	fmt.Println(value)
	switch value.Kind() {
	case reflect.String:
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buff, value)
	case reflect.Struct:
		return writeStruct(buff, value)
	default:
		return errors.New("un support kind"+ value.Kind().String())
	}
	return nil
}

func writeSlice(buff *bytes.Buffer, value reflect.Value) error {
	buff.WriteString("[")

	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)

		writeAny(buff, sliceValue)

		if s < value.Len()-1 {
			buff.WriteString(",")
		}
	}
	buff.WriteString("]")
	return nil
}

func writeStruct(buff *bytes.Buffer, value reflect.Value) error {
	valueType := value.Type()

	buff.WriteString("{")

	for i:=0; i<value.NumField(); i++ {
		//
		fieldValue := value.Field(i)
		//
		fieldType := valueType.Field(i)
		//
		buff.WriteString("\"")
		//
		buff.WriteString(fieldType.Name)
		//
		buff.WriteString("\":")
		//
		writeAny(buff, fieldValue)
		//
		if i<value.NumField()-1 {
			buff.WriteString(",")
		}
	}

	buff.WriteString("}")
	return nil
}

// json语法规则

// JSON语法是 Javascript对象表示法语法 的 子集
// 1、数据在名称（值对）中
// 2、数据由逗号分割
// 3、花括号保存对象
// 4、方括号保存数组

// JSON名称（值对）
// 名称（值对）：字段名称（在双引号）中，后面冒号，然后是值
// 值可以是以下几种类型
//  1、数字（整数，浮点）
//  2、字符串（双引号中）
//  3、逻辑值 （true或false）
//  4、数组 （方括号中）
//  5、对象 （花括号中）
//  6、null



// json值
// javascript对象表示法， 存储和交换文本信息的语法  类似xml
// 更小，更快，更易解析
// JSON 轻量级 文本数据 交换格式
// json文本格式 在 语法上 与创建 javascript 对象的代码相同
// eval()函数， 用JSON数据 生成 原生的 Javascript对象
// JSON数据可以通过Ajax进行传输




