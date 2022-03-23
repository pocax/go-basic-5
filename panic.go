package main

import (
	"errors"
	"fmt"
)

// func main() {
// 	var password string
// 	fmt.Scanln(&password)

// 	if valid, err := validPassword(password); err != nil {
// 		panic(err.Error())
// 	} else {
// 		fmt.Println(valid)
// 	}
// }

// func validPassword(password string) (string, error) {
// 	pl := len(password)

// 	if pl < 5 {
// 		return "", errors.New("Password must be at least 5 characters long")
// 	}

// 	return "Valid", nil
// }

func main() {
	defer catchError()
	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchError() {
	if r := recover(); r != nil {
		fmt.Println("Error:", r)
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password must be at least 5 characters long")
	}

	return "Valid", nil
}
