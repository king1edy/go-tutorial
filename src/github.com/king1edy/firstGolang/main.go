package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	f := fib()

	fmt.Println(f(), f(), f(), f(), f(), f())
}

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
