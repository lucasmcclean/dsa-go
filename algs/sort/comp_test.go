package sort

import (
	"testing"

	"github.com/ljmcclean/dsa-go/algs"
)

const seed int64 = 19
const size = 100000
const intSize = algs.MaxInt

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := algs.GenSliceInt(size, intSize, seed)
		b.StartTimer()
		QuickSort(slice)
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := algs.GenSliceInt(size, intSize, seed)
		b.StartTimer()
		QuickSortUnique(slice)
	}
}
