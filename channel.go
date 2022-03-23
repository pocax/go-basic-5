package main

import (
	"fmt"
	"time"
)

// func main() {
// 	c := make(chan int)

// 	c <- value

// 	result := <-c
// }

// func main() {
// 	c := make(chan string)

// 	msg1 := <-c
// 	fmt.Println(msg1)

// 	msg2 := <-c
// 	fmt.Println(msg2)

// 	msg3 := <-c
// 	fmt.Println(msg3)

// 	close(c)
// }

// func introduce(student string, c chan string) {
// 	result := fmt.Sprintf("Hi, I'm %s", student)
// 	c <- result
// }

//channel with anonymous function

// func main() {
// 	c := make(chan string)

// 	student := []string{"John", "Mary", "Mike"}

// 	for _, v := range student {
// 		go func(student string) {
// 			fmt.Println("Student", student)
// 			result := fmt.Sprintf("Hi, I'm %s", student)
// 			c <- result
// 		}(v)
// 	}

// 	for i := 1; i <= 3; i++ {
// 		print(c)
// 	}
// 	close(c)
// }

// func print(c chan string) {
// 	fmt.Println(<-c)
// }

//implement channel direction

// func main() {
// 	c := make(chan string)

// 	student := []string{"John", "Mary", "Mike"}

// 	for _, v := range student {
// 		go introduce(v, c)
// 	}

// 	for i := 1; i <= 3; i++ {
// 		print(c)
// 	}
// 	close(c)
// }

// func print(c chan string) {
// 	fmt.Println(<-c)
// }

// func introduce(student string, c chan string) {
// 	result := fmt.Sprintf("Hi, I'm %s", student)
// 	c <- result
// }

// buffer v unbuffered channel

// func main() {
// 	c1 := make(chan int)

// 	go func(chan int) {
// 		fmt.Println("func go routine start sending data into channel")
// 		c1 <- 10
// 		fmt.Println("func go routine generate after sending data into channel")
// 	}(c1)

// 	fmt.Println("main go routine sleep for 2 seconds")
// 	time.Sleep(2 * time.Second)

// 	fmt.Println("main go routine start receiving data from channel")
// 	d := <-c1
// 	fmt.Println("main go routine receive data from channel", d)

// 	close(c1)
// 	time.Sleep(time.Second)

// }

// unbuffered channel

// func main() {
// 	c1 := make(chan int, 3)

// 	go func(c chan int) {
// 		for i := 1; i < 5; i++ {
// 			fmt.Println("func go routine start sending data into channel", i)
// 			c <- i
// 			fmt.Println("func go routine generate after sending data into channel", i)
// 		}
// 		close(c)
// 	}(c1)

// 	fmt.Println("main go routine sleep for 2 seconds")
// 	time.Sleep(2 * time.Second)

// 	for v := range c1 {
// 		fmt.Println("main go routine receive data from channel", v)
// 	}
// }

// select channel
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "Hello"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		c2 <- "World"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
