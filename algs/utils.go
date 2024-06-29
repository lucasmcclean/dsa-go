package algs

import (
	"math/rand"
	"unsafe"

	"github.com/ljmcclean/dsa-go/types"
)

const MaxInt = 2147483647

// Generates a slice of strings.
func GenSliceStr(size int, strLen int, seed int64) []string {
	var slice []string
	for i := 0; i < size; i++ {
		slice = append(slice, RandString(strLen, seed))
	}
	return slice
}

// Generates a slice of positive integers.
func GenSliceInt(size int, intSize int, seed int64) []int {
	var slice []int
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < size; i++ {
		slice = append(slice, r.Intn(intSize))
	}
	return slice
}

// Generates a slice of ascending integers from 0 to 'size'.
func GenSliceIntAsc(size int) []int {
	var slice []int
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	return slice
}

// Generates a slice of descending integers from 'size' to 0.
func GenSliceIntDes(size int) []int {
	var slice []int
	for i := size; i >= 0; i-- {
		slice = append(slice, i)
	}
	return slice
}

// Generates a slice of positive float64s.
func GenSliceFloat(size int, floatSize float64, seed int64) []float64 {
	var slice []float64
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < size; i++ {
		slice = append(slice, r.Float64()*floatSize)
	}
	return slice
}

// Generates an empty slice.
func GenSliceEmpty(size int) []int {
	return make([]int, size)
}

// Returns nil.
func GenSliceNil() []int {
	return nil
}

// Checks if slice is sorted in ascending order.
func IsSorted[O types.Ordered](slice []O) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}

// Created by "icza" and modified by "moorara" original can be found at:
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandString(n int, seed int64) (randStr string) {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits := 6
	var letterIdxMask int64 = 1<<letterIdxBits - 1
	letterIdxMax := 63 / letterIdxBits

	r := rand.New(rand.NewSource(seed))
	b := make([]byte, n)
	for i, cache, remain := n-1, r.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
