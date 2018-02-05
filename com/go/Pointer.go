package main

import "fmt"
/**
 *
 * @author liuming-pc
 * @date 2018/2/5 10:22
 * @return
 * @version v1.0
 * Go语言的指针
 */

func main() {
	pointer()
	pointerArray()
}

func pointer() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a变量的存储地址是：%x\n", &a)
	fmt.Printf("ip变量的存储地址是：%x\n", ip)
	fmt.Printf("*ip变量的值是：%d\n", *ip)
}

func pointerArray() {
	const MAX int = 3
	var ptr [MAX] *int
	a := [3]int {1, 2, 3}
	var i int
	// 指针指向变量的值
	for i = 0; i < MAX; i ++ {
		ptr[i] = &a[i]
	}
	// 输出指针指向的变量的值
	for i = 0; i < MAX; i ++ {
		fmt.Println(*ptr[i])
	}
}