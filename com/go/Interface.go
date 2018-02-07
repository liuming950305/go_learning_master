package main

import "fmt"

type Connector interface {
	Connect()
}

type USB interface {
	Name()
	Connector
}

type Kingston struct {
	name string
}

// 实现Connector函数
func (k Kingston) Connect() {
	fmt.Println("Connected.Device's Name:", k.name)
}
// 实现Name函数
func (k Kingston) Name() {
	fmt.Println("Hello, I'm Kingston USB device")
}

func DisConnect(usb interface{}) {
	// 判断类型
	switch v := usb.(type){
	case Kingston:
		fmt.Println("DisConnect from device:", v.name)
	default:
		fmt.Println("Unkonwn device disConnect")
	}
}

func main() {
	// 这种方式不能准确表达是否实现了接口
	usb := Kingston{name: "Kingston"}
	//var usb USB = Kingston{name: "Kingston"}

	usb.Connect()
	usb.Name()
	DisConnect(usb)
}
