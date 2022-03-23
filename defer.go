package main

import (
	"fmt"
	"os"
)

// func main() {
// 	defer fmt.Println("defer function start to execute")
// 	fmt.Println("Hai everyone")
// 	fmt.Println("Welcome back to Go")
// }

// func main() {
// 	callDeferfunc(false)
// 	fmt.Println("Hai everyone")
// }

// func callDeferfunc(condition bool) {
// 	if condition == false {
// 		fmt.Println("Dont call defer function")
// 		return
// 	}
// 	defer deferFunc()
// }

// func deferFunc() {
// 	fmt.Println("defer function start to execute")
// }

func main() {
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before exiting")
	os.Exit(1)
}
