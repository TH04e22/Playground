package main

import (
	"fmt"
	"math/rand/v2"
)

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	var pivot int = arr[left]
	var i, j int = left + 1, right - 1

	for i < j {
		for arr[i] < pivot && i < right {
			i++
		}

		for arr[j] > pivot && j > left {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	arr[left], arr[j] = arr[j], arr[left]
	quickSort(arr, left, j)
	quickSort(arr, j+1, right)
}

func main() {
	const size = 10
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.IntN(100)
	}

	fmt.Println("Before sorting:", arr)

	quickSort(arr, 0, size)

	fmt.Println("After sorting:", arr)
}
