package main
//go语言中，数组可以存储统同一类型的数据， 但在结构体中，我们可以为不同项第一不同的数据类型

//定义结构体， 需要使用type和struct语句，struct语句定义一个新的数据类型，结构体中有一个或多个成员。type语句设定了结构体的名称
//type struct_variable_name struct {
//	member definition;
//}
import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "明仔","go语言大全", 22})

	// 也可以使用 key==>value 格式, 忽略的字段为0或者空字符串
	fmt.Println(Books{title:"Go 语言", author:"明仔1", subject:" go语言大全", book_id:12321})

	// 访问结构体的成员
	book := Books{"Go 语言", "明仔","go语言大全", 23232}
	fmt.Println(book.author)

	// 结构体作为函数参数
	bookprint(book)

	// 结构体指针
	var book_ptr *Books
	book_ptr = &book
	(*book_ptr).bookchange("aa")
	fmt.Printf("==\n")
	fmt.Println(book)

}

func bookprint(book Books) {
	fmt.Println(book.author)
	fmt.Println(book.title)
	fmt.Println(book.subject)
	fmt.Println(book.book_id)
}

// 结构体的方法的定义
func (book Books) bookchange(a string) Books{
	fmt.Println(a)
	book.book_id=1234
	book.subject = "金瓶梅"
	book.author = "明仔"
	book.title = "金瓶梅大全"
	return book
}


