package main

import "fmt"

func main() {
	var array = [] int{1, 2, 3, 4, 5, 6, 7, 8}
	c := make(chan int)
	go sum(array[:len(array)/3], c)
	go sum(array[len(array)/2:], c)
	x, y := <-c , <-c
	fmt.Println("x + y = ", x + y)
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	// 将值送入c
	c <- sum
}
