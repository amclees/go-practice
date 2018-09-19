package fib

func Fib(n int) int {
    a := 1
    b := 1

    for ; n > 0; n-- {
        c := a + b
        a = b
        b = c
    }

    return a
}
