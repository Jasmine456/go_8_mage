package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("逻辑处理器数目：%d\n", runtime.NumCPU())
}
