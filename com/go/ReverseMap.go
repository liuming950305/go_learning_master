package main

import "fmt"

func main() {
	map1 := map[int]string{1:"OK", 2: "Bad", 3: "Good", 4: "Great", 5: "Perfect"}
	fmt.Println(reverse(map1))
}

func reverse(source map[int]string) (target map[string]int) {
	temp := make(map[string]int)
	for k,v := range source {
		// 根据key取得value
		temp[v] = k
	}
	return temp
}
