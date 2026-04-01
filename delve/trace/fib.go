package main

import "fmt"

func fib(n int, ch chan<- int) {
	first := 1
	second := 2
	for i := 0; i < n; i++ {
		first, second = second, first+second
	}

	ch <- second
}

func main() {
	ch := make(chan int)

	go fib(10, ch)

	fmt.Println(<-ch)
}
