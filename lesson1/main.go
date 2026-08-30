package main

import (
    "fmt"
    "bufio"
    "os"
    "lesson1/lib"
)

func main() {
    var(
        n int = 5;
    )
    fmt.Println(Fib(n));
    bufio.NewReader(os.Stdin).ReadBytes('\n');
}