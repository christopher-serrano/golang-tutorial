package main

import "fmt"

func main() {
	var username string = "wagslane"
	var password_int int = 123123123123

	// parsing to string() doesn't work
	fmt.Println("Authorization Basic: ", username+":"+string(password_int))

	var password_str string = "123123123123"

	// setting the correct type from the beginning DOES work (duh)
	fmt.Println("Authorization Basic: ", username+":"+password_str)
}
