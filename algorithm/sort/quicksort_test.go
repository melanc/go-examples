package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var a []int = []int{6, 1, 2, 5, 9, 3, 4, 7, 10, 8}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quickSort(arr []int, l int, r int){
	if l > r{
		return
	}
	i, j := l, r
	base := l
	for i != j {
		for arr[j] >= arr[base] && i < j {
			j--
		}
		for arr[i] <= arr[base] && i < j {
			i++
		}

		if i < j {
			t:= arr[i]
			arr[i] = arr[j]
			arr[j] = t
		}
	}
	t := arr[i]
	arr[i] = arr[base]
	arr[base] = t

	quickSort(arr, l, i-1)
	quickSort(arr, i+1, r)
}

