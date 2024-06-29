package sort

import "github.com/ljmcclean/dsa-go/types"

func ShellSort[O types.Ordered](arr []O, gap int) {
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
