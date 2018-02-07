package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello() {
	fmt.Println("Hello, Reflection")
}
func (u User) GetName() string{
	return u.Name
}

func info(o interface{}) {
	// 取得类型
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())
	// 取得值
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	// 通过循环输出值
	for i := 0; i < v.NumField(); i ++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s %v = %v\n", t.Name(), f.Type, val)
	}
	// 输出方法
	for i := 0; i < v.NumMethod(); i ++ {
		m := t.Method(i)
		fmt.Printf("%6s %v\n", m.Name, m.Type)
	}
}

func main() {
	u := User{1, "大明", 23}
	info(u)
}


