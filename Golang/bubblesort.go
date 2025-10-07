// bubble_sort.go
// Run with: go run bubble_sort.go

package main

import (
	"fmt"
)

// bubbleSort sorts the slice in ascending order using Bubble Sort algorithm.
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // swap
				swapped = true
			}
		}
		// If no two elements were swapped in the inner loop, break early
		if !swapped {
			break
		}
	}
}

func main() {
	var n int
	fmt.Print("Enter number of elements: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter", n, "elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	bubbleSort(arr)

	fmt.Println("Sorted array:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
