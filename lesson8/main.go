package main

import (
	"fmt"
)

func endProgram() {
	fmt.Print("\n\nНажмите Enter для продолжения...")
	fmt.Scanln()
}

type User struct {
	Name string
	Age  int
}

func main() {
	var user User
	user.Name = "User name"
	user.Age = 10

	fmt.Println(user.Name, user.Age)
	endProgram()
}
