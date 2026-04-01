package main

import "fmt"

func recur(n int) {
	if n == 0 {
		return
	} else {
		fmt.Println(n)
		recur(n - 1)
	}
}

func main() {
	recur(3)
}
