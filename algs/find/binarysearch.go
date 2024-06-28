package find

import "github.com/ljmcclean/dsa-go/types"

func BinarySearch[O types.Ordered](arr []O, target O) (index int) {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (r-l)/2 + l
		if arr[mid] > target {
			r = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
