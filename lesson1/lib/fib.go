package lib

func Fib(n int) int{
    if (n <= 1) {
        return 1;
    } else {
        return n * Fib(n - 1);
    }
}