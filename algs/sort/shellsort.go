package sort

import "github.com/ljmcclean/dsa-go/types"

func ShellSort[O types.Ordered](arr []O, gap int) {
	if arr == nil {
		return
	} else if len(arr) == 1 {
		return
	}
	UShellSort(arr, gap)
}

// Bypasses the check made by Sort to ensure arr isn't nil
// and that the arr is longer than one element.
func UShellSort[O types.Ordered](arr []O, gap int) {
	length := len(arr)
	h := 1
	for h < length/gap {
		h = h*gap + 1
	}
	for h >= 1 {
		for i := h; i < length; i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h /= gap
	}
}
