package main
//
//import (
//	"encoding/base64"
//	"fmt"
//)
//
////一个指针变量指向一个值的内存地址，声明格式
////var var_name *var_type
//var ip *int
//var sp *string
//
////go语言指针数组
//const MAX int = 3
//
////有一种情况，我们可能需要保存数组，这样我们就需要使用到指针
//var ptr [MAX]*int
//
//// go指向指针的指针, 指向指针的指针变量的声明格式如下
//var ptr1 [MAX]**int
//
////三重指针
//var ptr2 [MAX]***int
//
////指针作为函数参数，只需要设置参数类型为指针形式
//
////new() 函数，创建指针的另一种方法
//str := new(string)
//*str = "aaa"
//// make() 用来生成实例对象, 生成数组，结构体，接口， map
//// map元素的删除delete()内建函数 函数从map中删除键值对， delete(map, key)
//
//
//func main() {
//	var s string="ss"
//	var i int=23
//	var n *string
//	ip = &i
//	sp = &s
//	a := [] int {11, 22, 33}
//	var a_int int = 23
//	var b_int int = 34
//
//	fmt.Println(ip, sp)
//	fmt.Println(*ip, *sp)
//	if n==nil {
//		fmt.Println(n)
//	}
//	if n != nil {
//		fmt.Println(n, a)
//	}
//	for i:=0; i<MAX; i++ {
//		ptr[i] = &a[i]
//		ptr1[i] = &ptr[i]
//		ptr2[i] = &ptr1[i]
//	}
//	for i:=0; i<MAX; i++ {
//		fmt.Println(&ptr[i], ptr[i], *ptr[i], **ptr1[i], ***ptr2[i])
//	}
//
//
//	fmt.Println(a_int, b_int)
//	swap(&a_int, &b_int)
//	fmt.Println(a_int, b_int)
//	message := "中国"
//	encodeMessage := base64.StdEncoding.EncodeToString([]byte (message))
//	fmt.Println(encodeMessage)
//	me,_ := base64.StdEncoding.DecodeString(encodeMessage)
//	fmt.Println(string(me))
//
//
//}
//
//// 参数的值为指针类型，就是引用传递
//func swap(x *int, y *int) {
//	//var temp int
//	//temp = *x
//	//*x = *y
//	//*y = temp
//	*x, *y = *y, *x
//	//x, y = y, x
//}
//
//
//
