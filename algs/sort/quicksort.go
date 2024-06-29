package sort

import (
	"github.com/ljmcclean/dsa-go/types"
)

const insertionCutoff = 25

// Optimized for case of duplicate elements.
func QuickSort[O types.Ordered](arr []O) {
	for {
		if len(arr) < insertionCutoff {
			InsertionSort(arr)
			return
		}
		piv := arr[ninther(arr)]
		l, r := partition(arr, piv)
		QuickSort(arr[:l])
		arr = arr[r:]
	}
}

// Optimized for case of mostly unique elements.
func QuickSortUnique[O types.Ordered](arr []O) {
	if len(arr) < insertionCutoff {
		InsertionSort(arr)
		return
	}
	piv := arr[ninther(arr)]
	l, r := hoarePartition(arr, piv)
	if r > 0 {
		QuickSortUnique(arr[:r+1])
	}
	if l < len(arr)-1 {
		QuickSortUnique(arr[l:])
	}
}

// Three-way partition for case of duplicate elements.
func partition[O types.Ordered](arr []O, piv O) (lPivIdx, rPivIdx int) {
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

// Hoare partition for case of mostly unique elements.
func hoarePartition[O types.Ordered](arr []O, piv O) (lPivIdx, rPivIdx int) {
	l, r := 0, len(arr)-1
	for l <= r {
		for arr[l] < piv {
			l++
		}
		for arr[r] > piv {
			r--
		}
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	return l, r
}

// Implementation of Tukey's ninther (median of medians) to
// ensure more consistent performance.
func ninther[O types.Ordered](arr []O) (pivIdx int) {
	n := len(arr)
	a := medOfThree(arr, 0, n/6, n/3)
	b := medOfThree(arr, n/3, 5*n/6, 2*n/3)
	c := medOfThree(arr, 2*n/3, 5*n/6, n-1)
	return medOfThree(arr, a, b, c)
}

// Find the median value of arr between three indexes and return
// that index.
func medOfThree[O types.Ordered](arr []O, lo, med, hi int) (pivIdx int) {
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
