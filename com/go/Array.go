package main

import "fmt"

func main() {

	//printArray()
	muiltiArray()
}

// single array
func printArray() {
	var array = [5]float32{100.0, 25.6, 77.6, 961, 220}
	var i int
	for i = 0; i < 5; i ++ {
		fmt.Println(array[i])
	}
}

// muilti array
func muiltiArray() {
	var multi = [][]float32 {
		{2.0, 88.9},
		{4.44, 6.98},
		{5.65, 999},
	}
	var i, j int
	for i = 0; i < 3; i ++ {
		for j = 0; j < 2; j ++ {
			fmt.Println(multi[i][j])
		}
	}
}

