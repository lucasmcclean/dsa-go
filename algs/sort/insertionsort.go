package sort

import "github.com/ljmcclean/dsa-go/types"

func InsertionSort[O types.Ordered](arr []O) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
