package qsort

import (
	"sync"

	u "github.com/ljmcclean/dsa-go/utils"
)

func PSort[O u.Ordered](arr []O) {
	if arr == nil {
		return
	} else if len(arr) == 1 {
		return
	}
	prepareSort(arr)
}

// Bypasses the check made by Sort to ensure arr isn't nil
// and that the arr is longer than one element.
func UPSort[O u.Ordered](arr []O) {
	prepareSort(arr)
}

// Setup and cleanup wait group for parallel execution.
func prepareSort[O u.Ordered](arr []O) {
	var wg sync.WaitGroup
	wg.Add(1)
	psort(arr, &wg)
	wg.Wait()
}

// Recursive parallel quicksort.
func psort[O u.Ordered](arr []O, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		pil, pir := partition(arr)
		wg.Add(1)
		go psort(arr[:pil], wg)
		arr = arr[pir:]
	}
}
