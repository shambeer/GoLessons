package main

import "fmt"

func fib(n int) int{
    if (n <= 1) {
        return 1;
    } else {
        return n * fib(n - 1);
    }
}

func main() {
    var(
        n int = 5;
    )
    fmt.Println(fib(n));
}