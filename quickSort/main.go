package main

import (
	"fmt"
)

func devide(a []int, low, high int) int {
	pivote := a[high]
	for j := low; j < high; j++ {
		if a[j] < pivote {
			// swap
			temp := a[j]
			a[j] = a[low]
			a[low] = temp
			low++
		}
	}
	// swap
	temp := a[low]
	a[low] = a[high]
	a[high] = temp
	return low
}

func QuickSort(a []int, low, high int) {
	if low > high {
		return
	}

	pivote := devide(a, low, high)
	QuickSort(a, low, pivote-1)
	QuickSort(a, pivote+1, high)
}

func main() {
	list := []int{7, 8, 9, 1, 2, 5, 3, 4}

	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
