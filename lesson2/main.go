package main

import "fmt";

func main() {
	fmt.Println(sum(6, 7));
	fmt.Scanln();
}

func sum(a int, b int) int {
	return a + b;
}