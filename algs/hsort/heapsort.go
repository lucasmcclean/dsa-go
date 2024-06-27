package hsort

import (
	u "github.com/ljmcclean/dsa-go/utils"
)

func Sort[O u.Ordered](arr []O) {
	if arr == nil {
		return
	} else if len(arr) == 1 {
		return
	}
	USort(arr)
}

// Bypasses the check made by Sort to ensure arr isn't nil
// and that the arr is longer than one element.
func USort[O u.Ordered](arr []O) {
	length := len(arr)
	heapify(arr, length)
	for i := length - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i)
	}
}

func heapify[O u.Ordered](arr []O, length int) {
	for i := (length/2 - 1); i >= 0; i-- {
		siftDown(arr, length, i)
	}
}

func siftDown[O u.Ordered](arr []O, length int, root int) {
	for {
		child := root*2 + 1
		if child >= length {
			return
		}
		if child+1 < length && arr[child+1] > arr[child] {
			child++
		}
		if arr[root] >= arr[child] {
			return
		}
		arr[root], arr[child] = arr[child], arr[root]
		root = child
	}
}
