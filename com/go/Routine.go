package main

import (
	"time"
	"fmt"
)

func main() {
	go say("world")
	say("Hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}