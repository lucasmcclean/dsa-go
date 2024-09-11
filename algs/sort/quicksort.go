package sort

import (
	"math"

	"github.com/ljmcclean/dsa-go/ds/stack"
	"github.com/ljmcclean/dsa-go/types"
)

const insertionCutoff = 17

func maxStackSize(length int) int {
	return (int(math.Log2(float64(length))) + 3) * 2
}

// QuickSort is optimized for the case of duplicate elements.
func QuickSort[O types.Ordered](arr []O) {
	stk := stack.Unsafe(maxStackSize(len(arr)), 0, len(arr))
	for stk.Len() != 0 {
		hi, lo := stk.Pop(), stk.Pop()
		if hi-lo < insertionCutoff {
			InsertionSort(arr[lo:hi])
			continue
		}
		piv := arr[ninther(arr[lo:hi])+lo]
		pil, pir := partition(arr[lo:hi], piv)
		pil += lo
		pir += lo
		if pil > lo {
			stk.Push(lo, pil)
		}
		if pir < hi {
			stk.Push(pir, hi)
		}
	}
}

// QuickSortUnique is optimized for the case of primarily unique elements.
func QuickSortUnique[O types.Ordered](arr []O) {
	stk := stack.Unsafe(maxStackSize(len(arr)), 0, len(arr))
	for stk.Len() != 0 {
		hi, lo := stk.Pop(), stk.Pop()
		if hi-lo < insertionCutoff {
			InsertionSort(arr[lo:hi])
			continue
		}
		piv := arr[ninther(arr[lo:hi])+lo]
		pi := hoarePartition(arr[lo:hi], piv)
		pi += lo
		if pi+1 > lo {
			stk.Push(lo, pi+1)
		}
		if pi < hi {
			stk.Push(pi, hi)
		}
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
			r--
			arr[i], arr[r] = arr[r], arr[i]
		default:
			i++
		}
	}
	return l, r
}

// Hoare partition for case of mostly unique elements.
func hoarePartition[O types.Ordered](arr []O, piv O) (pivIdx int) {
	l, r := 0, len(arr)-1
	for {
		for arr[l] < piv {
			l++
		}
		for arr[r] > piv {
			r--
		}
		if l >= r {
			return r
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

// Tukey's ninther (median of medians) to ensure more consistent performance.
func ninther[O types.Ordered](arr []O) (pivIdx int) {
	n := len(arr)
	var a, b, c int
	a = medOfThree(arr, 0, n/6, n/3)
	b = medOfThree(arr, n/3, 5*n/6, 2*n/3)
	c = medOfThree(arr, 2*n/3, 5*n/6, n-1)
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
