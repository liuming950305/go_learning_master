package main

import (
	"fmt"
	"math"
)
/**
 * 
 * @author liuming-pc
 * @date 2018/2/5 14:27
 * @return 
 * @version v1.0
 * 函数也可以代表值
 */

// 定义结构体
type Vertex struct {
	X, Y float64
}

// 使用receiver绑定函数
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	temp := func(x, y int) int {
		return x + y
	}
	fmt.Println(temp(1, 2))

	// 使用结构体
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}
