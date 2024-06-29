package sort

import "github.com/ljmcclean/dsa-go/types"

func HeapSort[O types.Ordered](arr []O) {
	length := len(arr)
	heapify(arr, length)
	for i := length - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i)
	}
}

func heapify[O types.Ordered](arr []O, length int) {
	for i := (length/2 - 1); i >= 0; i-- {
		siftDown(arr, length, i)
	}
}

func siftDown[O types.Ordered](arr []O, length int, root int) {
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
