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
	prepareSort(arr)
}

// *UnsafeParallelSort* does not verify that arr isn't nil
func UPSort[O u.Ordered](arr []O) {
	prepareSort(arr)
}

func prepareSort[O u.Ordered](arr []O) {
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
