package demo

import (
	"time"
	"fmt"
)

func timeCount() time.Duration {
	start := time.Now()
	sum := 0
	for i:=0; i<1000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	return elapsed
}

//func main()  {
//	test()
//}

