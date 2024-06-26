package qsort

import (
	"github.com/ljmcclean/dsa-go/algs/isort"
	u "github.com/ljmcclean/dsa-go/utils"
)

const insertionCutoff = 30

// Helper to ensure arr isn't nil
func Sort[O u.Ordered](arr []O) {
	if arr == nil {
		return
	}
	sort(arr)
}

// quicksort with insertion sort and tail recursion optimization
func sort[O u.Ordered](arr []O) {
	for {
		if len(arr) < insertionCutoff {
			isort.Sort(arr)
			return
		}
		pil, pir := partition(arr)
		sort(arr[:pil])
		arr = arr[pir:]
	}
}

// 3 way partition for case of duplicate elements
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

// implementation of Tukey's ninther (median of medians)
func ninther[O u.Ordered](arr []O) (piv O) {
	n := len(arr)
	mid := n / 2
	deltad := n / 4
	delta := n / 8
	a := medOfThree(arr, 0, delta, deltad)
	b := medOfThree(arr, mid-delta, mid, mid+delta)
	c := medOfThree(arr, n-deltad, n-delta, n-1)
	return arr[medOfThree(arr, a, b, c)]
}

// find the median value and return its index
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
