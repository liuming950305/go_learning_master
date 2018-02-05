package main

import (
	"fmt"
)

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i ++ {
		// defer 延迟调用, 会被压入defer栈中
		defer fmt.Println(i)
	}
	fmt.Println("done")
}