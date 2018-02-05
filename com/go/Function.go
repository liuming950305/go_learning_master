package main

import "fmt"

func main() {
    fmt.Println("The max result：", max(10,20))
    // It uses in function.
    a,b := swap("David", "Parker")
    fmt.Println("Swap result:", a, b)
    // 闭包
    result := getSequence()
    fmt.Println(result())
    fmt.Println(result())
    fmt.Println(result())
    result1 := getSequence()
    fmt.Println(result1())
    fmt.Println(result1())
    fmt.Println(result1())
}

func max(source, target int) int {
    if(source > target) {
        return source
    }
    return target
}

func swap(source, target string) (string,string) {
    return target,source
}

// 函数的闭包
func getSequence() func() int {
    i := 0
    return func() int {
        i += 1
        return i;
    }
}