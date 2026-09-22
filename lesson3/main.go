package main

import (
	"fmt"
	"lesson3/math_shambeer"
);

func main() {
	const (
		a int = 6
		b int = 7
	)
	
	fmt.Println(math_shambeer.Sum(a, b));
	fmt.Println(math_shambeer.Diff(a, b));
	fmt.Scanln();
}