package main

import "fmt"

func main() {
	hashMap()
}

func hashMap() {
	// 初始化创建一个map
	var hashMap map[string]string = make(map[string]string)

	hashMap["hunan"] = "ChangSha"
	hashMap["guangdong"] = "GuangZhou"
	hashMap["anhui"] = "HeFei"
	hashMap["hubei"] = "WuHan"

	for city := range hashMap {
		fmt.Printf("%s -> %s\n", city, hashMap[city])
	}

	capital, ok := hashMap["sichuan"]
	if(ok) {
		fmt.Println("Capital of sichuan is", capital)
	} else {
		fmt.Println("Capital of sichuan is not in this map")
	}
}
