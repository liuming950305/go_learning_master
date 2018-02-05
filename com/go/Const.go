package main

import "fmt"

/**
 * 
 * @author liuming-pc
 * @date 2018/2/5 10:00
 * @return 
 * @version v1.0
 * Const use
 */

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "常量"
	area = WIDTH * LENGTH
	fmt.Printf("面积为 : %d", area)
	fmt.Println()
	fmt.Println(a, b, c)

	useIota()
}


// iota

func useIota() {
	const(
		// i = 1 << 0
		i = 1 << iota
		// j == 3 << 1
		j = 3 << iota
		// k = 3 << 2
		k
		// k = 3 << 3
		l
	)
	fmt.Println(i, j, k, l)
}
