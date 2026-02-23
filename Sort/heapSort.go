package main

import (
	"fmt"
	"math/rand/v2"
)

func restore(arr []int, root, n int) {
	var child int = 2 * root
	var temp = arr[root]

	for child < n {
		if (child+1) < n && arr[child] < arr[child+1] {
			child = child + 1
		}

		if temp < arr[child] {
			arr[root] = arr[child]
			root = child
			child = 2 * child
		} else {
			break
		}
	}

	arr[root] = temp
}

func heapSort(arr []int) []int {
	arr = append([]int{0}, arr...)

	n := len(arr)

	// restore
	for i := (n - 1) / 2; i >= 1; i-- {
		restore(arr, i, n)
	}

	// sort
	for i := n - 1; i >= 2; i-- {
		arr[1], arr[i] = arr[i], arr[1]

		restore(arr, 1, i)

	}

	arr = arr[1:]
	return arr
}

func main() {
	const size = 10
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.IntN(100)
	}

	fmt.Println("Before sorting:", arr)

	arr = heapSort(arr)

	fmt.Println("After sorting:", arr)
}
