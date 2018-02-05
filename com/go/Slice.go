package main

import "fmt"

func main() {
	slice()
}

func slice() {

	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	numbers = append(numbers, 1, 2, 3, 4)
	printSlice(numbers)

	// 通过make函数创建slice, 并将容量扩容2倍
	numbers1 := make([]int, len(numbers) ,cap(numbers) * 2)
	printSlice(numbers1)

}

func printSlice(numbers []int) {
	fmt.Printf("len= %d, cap=%d, slice=%v\n", len(numbers), cap(numbers), numbers)
}
