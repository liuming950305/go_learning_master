package main

import "fmt"

/**
 *
 * @author liuming-pc
 * @date 2018/2/5 12:39
 * @return
 * @version v1.0
 * 类型转换
 */
func main() {
	transform()
}

func transform() {
	var sum int = 17
	var count int = 5
	var number float32
	var f float32 = 98.88

	number = float32(sum) / float32(count)
	fmt.Println("Result:", number)
	fmt.Println("f:", float64(f))
}
