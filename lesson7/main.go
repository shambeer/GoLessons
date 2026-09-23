package main

import (
	"fmt"
);

func endProgram() {
	fmt.Println("\n\nНажмите Enter для продолжения...");
	fmt.Scanln();
}

type User struct {
	name string
	age int
}

func inputUser() User {
	var user User;

	fmt.Println("Введите имя пользователя: ")
	fmt.Scanln(&user.name)
	fmt.Println("Введите возраст пользователя: ")
	fmt.Scanln(&user.age)
	fmt.Print("\n\n")

	return user
}

func outUser(user User) {
	fmt.Println("Имя:", user.name)
	fmt.Println("Возраст:", user.age)
}

func main() {
	var user1 User
	
	user1 = inputUser()
	outUser(user1)
	endProgram()
}