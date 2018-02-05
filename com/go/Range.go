package main

import "fmt"

func main() {
	rangeArray()
	rangeMap()
	rangeUnicode()
}

func rangeArray() {
	var arrays = []int {1, 2, 3, 4}
	sum := 0
	for _, array := range arrays {
		sum += array
	}
	fmt.Printf("The slice's reuslt: %d\n", sum);
}

func rangeMap() {
	kvs := map[string] string {"a": "apple", "b": "banana"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}
}

func rangeUnicode() {
	// index, Unicode
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
