package main

import "fmt"
/**
 * 
 * @author liuming-pc
 * @date 2018/2/4 22:32
 * @return 
 * @version v1.0
 * If go file exist the init function. It will be executed first.
 * Go do not need semicolon, every line is a sentence.
 */

func init() {
	fmt.Println("This function will be executed first")
}
func main() {
    fmt.Println("Hello,Go")
}
