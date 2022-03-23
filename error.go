package main

import (
	"errors"
	"fmt"
)

// func main() {
// 	var number int
// 	var err error
// 	number, err = strconv.Atoi("123GH")

// 	if err == nil {
// 		fmt.Println("Number:", number)
// 	} else {
// 		fmt.Println("Error:", err.Error())
// 	}

// 	number, err = strconv.Atoi("123")

// 	if err == nil {
// 		fmt.Println("Number:", number)
// 	} else {
// 		fmt.Println("Error:", err.Error())
// 	}
// }

func main() {
	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Valid:", valid)
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password must be at least 5 characters long")
	}

	return "Valid", nil
}
