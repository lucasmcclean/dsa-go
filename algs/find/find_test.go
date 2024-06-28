package find

import (
	"testing"

	"github.com/ljmcclean/dsa-go/algs"
)

func TestBinarySearch(t *testing.T) {
	sliceInt := algs.GenSliceIntAsc(100)
	index := BinarySearch(sliceInt, 42)
	if index != 42 {
		t.Error("Failed to search integers")
	}
	index = BinarySearch(sliceInt, 101)
	if index != -1 {
		t.Error("Failed to report not found")
	}
}
