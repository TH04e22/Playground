package main

import "fmt"

func callPanic() {
	defer func() {
		fmt.Println("d1")
	}()

	defer func() {
		fmt.Println("d2")
	}()

	defer func() {
		fmt.Println("d3")
	}()

	panic("Very dangerous!")
}

func main() {
	callPanic()
}
