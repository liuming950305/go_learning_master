package main

import "fmt"

/**
 *
 * @author liuming-pc
 * @date 2018/2/4 22:39
 * @return
 * @version v1.0
 * DataType:
 * 1. bool类型 bool
 * 2. Number float32 float64
 * 3. string string
 * 4. 派生类型
 * 指针类型（Pointer）
 * 数组类型
 * 结构化类型(struct)
 * Channel 类型
 * 函数类型
 * 切片类型
 * 接口类型（interface）
 * Map 类型
 * var identifier type
 */

func main() {
    var a string
    var b bool
    var c string = "runoboo.com"
     fmt.Println(a, b, c)
	data()
}

func data() {
    // 同时声明多个变量
    var str1, str2, str3 string = "This", "is my", "style"
    fmt.Println(str1, str2, str3)
}
