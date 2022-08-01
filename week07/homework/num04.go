package main

import (
	"fmt"
	"time"
)

/*
用defer优雅地打印函数的耗时
 */
func foo(i int) int {
	begin := time.Now()

	defer func(begin time.Time) {
		fmt.Printf("function use time %d ms\n", time.Since(begin).Milliseconds())
	}(begin)
	if i < 10 {
		time.Sleep(1*time.Millisecond)
		//fmt.Printf("function use time %d ms\n", time.Since(begin).Milliseconds())
		return i + 4
	} else if i < 20 {
		time.Sleep(2*time.Millisecond)
		//fmt.Printf("function use time %d ms\n", time.Since(begin).Milliseconds())
		return i * 4
	} else {
		time.Sleep(3*time.Millisecond)
		//fmt.Printf("function use time %d ms\n", time.Since(begin).Milliseconds())
		return 0
	}
}

func main()  {
	foo(115)
}