package sort

import (
	"sync"

	"github.com/ljmcclean/dsa-go/types"
)

func PQuickSort[O types.Ordered](arr []O) {
	if arr == nil {
		return
	} else if len(arr) == 1 {
		return
	}
	preparePQuickSort(arr)
}

// Bypasses the check made by Sort to ensure arr isn't nil
// and that the arr is longer than one element.
func UPQuickSort[O types.Ordered](arr []O) {
	preparePQuickSort(arr)
}

// Setup and cleanup wait group for parallel execution.
func preparePQuickSort[O types.Ordered](arr []O) {
	var wg sync.WaitGroup
	wg.Add(1)
	pQuickSort(arr, &wg)
	wg.Wait()
}

// Recursive parallel quicksort.
func pQuickSort[O types.Ordered](arr []O, wg *sync.WaitGroup) {
	defer wg.Done()
	for len(arr) > 0 {
		pil, pir := partition(arr)
		wg.Add(1)
		go pQuickSort(arr[:pil], wg)
		arr = arr[pir:]
	}
}
