package isort

import u "github.com/ljmcclean/dsa-go/utils"

func Sort[O u.Ordered](arr []O) {
	if arr == nil {
		return
	}
	USort(arr)
}

// *UnsafeSort* does not verify that arr isn't nil
func USort[O u.Ordered](arr []O) {
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
