package qsort

import (
	"sync"

	"github.com/ljmcclean/dsa-go/algs/isort"
	u "github.com/ljmcclean/dsa-go/utils"
)

func PSort[O u.Ordered](arr []O) {
	var wg sync.WaitGroup
	wg.Add(1)
	psort(arr, &wg)
	wg.Wait()
}

func psort[O u.Ordered](arr []O, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(arr) < 15 {
		isort.Sort(arr)
		return
	}
	pivIdx := partition(arr)
	wg.Add(2)
	go psort(arr[:pivIdx], wg)
	go psort(arr[pivIdx:], wg)
}
