package qsort

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
	for len(arr) > 0 {
		pil, pir := partition(arr)
		USort(arr[:pil])
		arr = arr[pir:]
	}
}

// Three-way partition for case of duplicate elements.
func partition[O u.Ordered](arr []O) (lPivIdx, rPivIdx int) {
	piv := ninther(arr)
	l, r := 0, len(arr)
	for i := 0; i < r; {
		switch {
		case arr[i] < piv:
			arr[i], arr[l] = arr[l], arr[i]
			i++
			l++
		case arr[i] > piv:
			arr[i], arr[r-1] = arr[r-1], arr[i]
			r--
		default:
			i++
		}
	}
	return l, r
}

// Implementation of Tukey's ninther (median of medians) to
// ensure more consistent performance.
func ninther[O u.Ordered](arr []O) (piv O) {
	n := len(arr)
	a := medOfThree(arr, 0, n/6, n/3)
	b := medOfThree(arr, n/3, 5*n/6, 2*n/3)
	c := medOfThree(arr, 2*n/3, 5*n/6, n-1)
	return arr[medOfThree(arr, a, b, c)]
}

// Find the median value of arr between three indexes and return
// that index.
func medOfThree[O u.Ordered](arr []O, lo, med, hi int) (pivIdx int) {
	a, b, c := arr[lo], arr[med], arr[hi]
	if a < b {
		switch {
		case b <= c:
			return med
		case a < c:
			return hi
		default:
			return lo
		}
	}
	switch {
	case a <= c:
		return lo
	case b < c:
		return hi
	default:
		return med
	}
}
