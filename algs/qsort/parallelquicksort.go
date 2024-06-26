package qsort

import (
	"sync"

	"github.com/ljmcclean/dsa-go/algs/isort"
	u "github.com/ljmcclean/dsa-go/utils"
)

func PSort[O u.Ordered](arr []O) {
	if arr == nil {
		return
	}
	pSort(arr)
}

func pSort[O u.Ordered](arr []O) {
	var wg sync.WaitGroup
	wg.Add(1)
	psort(arr, &wg)
	wg.Wait()
}

func psort[O u.Ordered](arr []O, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if len(arr) < insertionCutoff {
			isort.Sort(arr)
			return
		}
		pil, pir := partition(arr)
		wg.Add(1)
		go psort(arr[:pil], wg)
		arr = arr[pir:]
	}
}
