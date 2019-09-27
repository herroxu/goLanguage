package main

import "fmt"

type HashMap struct {
	key string
	value string
	hashcode int
	next *HashMap
}

var table [16]*HashMap

func initTable() {
	for i := range table {
		table[i] = &HashMap{"", "", i, nil}
	}
}

func getInstance() [16]*HashMap {
	if table[0] == nil {
		initTable()
	}
	return table
}

func getHashcode(k string) int {
	if len(k) == 0 {
		return 0
	}

	var hashcode int = 0
	var lastIndex int = len(k)-1

	for i := range k {
		if i == lastIndex {
			hashcode += int(k[i])
			break
		}
		hashcode += (hashcode+int(k[i]))*31
	}
	return hashcode
}

func getTableIndex(hashcode int) int {
	return hashcode % 16
}

func getNodeIndex(hashcode int) int {
	return hashcode >> 4
}

func put(k string, v string) string {

	var hashcode = getHashcode(k)

	var thisNode = HashMap{k, v, hashcode,nil}

	var headPtr [16]*HashMap = getInstance()

	var tableIndex = getTableIndex(hashcode)

	fmt.Println("tableIndex", tableIndex)

	//var nodeIndex = getNodeIndex(hashcode)

	var headNodePtr *HashMap = headPtr[tableIndex]

	if (*headNodePtr).key == "" {
		*headNodePtr = thisNode
		return ""
	}

	var lastNodePtr *HashMap = headNodePtr

	var nextNodePtr *HashMap = (*headNodePtr).next

	for nextNodePtr != nil {
		lastNodePtr = nextNodePtr
		nextNodePtr = (*nextNodePtr).next
		if (*lastNodePtr).key == k {
			(*lastNodePtr).value = v
			return ""
		}
	}

	// 这种写法，只是将thisNode的地址重新赋值给nextNodePtr这个变量上
	(*lastNodePtr).next = &thisNode
	fmt.Println("=======", *((*lastNodePtr).next))
	return ""
}


func get(k string) string {
	var hashcode = getHashcode(k)

	var tableIndex = getTableIndex(hashcode)

	var headPtr [16]*HashMap = getInstance()

	var nodePtr *HashMap = headPtr[tableIndex]

	if (*nodePtr).key == k {
		return (*nodePtr).value
	}

	for (*nodePtr).next != nil {
		nodePtr = (*nodePtr).next
		fmt.Println("============", (*nodePtr).key)
		if (*nodePtr).key == k {
			return (*nodePtr).value
		}

	}
	return ""
}



func main() {
	getInstance()

	put("a", "a_put_zhong_guo")
	put("afw", "afw_put_chong")

	put("asd", "asd_put")

	fmt.Println(get("a"))
	fmt.Println(get("afw"))
	fmt.Println(get("asd"))


}

