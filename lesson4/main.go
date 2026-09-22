package main

import (
	"fmt"
);

func endProgram() {
	fmt.Println("\n\nНажмите Enter для выхода...");
	fmt.Scanln();
}

func outArr(arr [5]int) {
	fmt.Println("Результат: ")
	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 == 0 {
			fmt.Print(arr[i], " ");
		} else {
			fmt.Print("  ");
		}
	}
}

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i;
	}

	outArr(arr);
	endProgram();
}