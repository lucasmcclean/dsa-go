package qsort

import (
	"math/rand"

	"github.com/ljmcclean/dsa-go/algs/isort"
	u "github.com/ljmcclean/dsa-go/utils"
)

func Sort[O u.Ordered](arr []O) {
	if len(arr) < 15 {
		isort.Sort(arr)
		return
	}
	pi := partition(arr)
	Sort(arr[:pi])
	Sort(arr[pi:])
}

func partition[O u.Ordered](arr []O) (pivotIndex int) {
	piv := pivot(arr)
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
	}
}

func pivot[O u.Ordered](arr []O) (piv O) {
	n := len(arr)
	a, b, c := arr[rand.Intn(n)], arr[rand.Intn(n)], arr[rand.Intn(n)]
	if a < b {
		if b <= c {
			return b
		} else if a < c {
			return c
		} else {
			return a
		}
	} else {
		if a <= c {
			return a
		} else if b < c {
			return c
		} else {
			return b
		}
	}
}
