package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

func GenSliceStr(size int, strLen int) []string {
	var slice []string
	for i := 0; i < size; i++ {
		slice = append(slice, RandString(strLen))
	}
	return slice
}

func GenSliceInt(size int, intSize int) []int {
	var slice []int
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Intn(intSize))
	}
	return slice
}

func GenSliceIntAsc(size int) []int {
	var slice []int
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	return slice
}

func GenSliceIntDes(size int) []int {
	var slice []int
	for i := size; i >= 0; i-- {
		slice = append(slice, i)
	}
	return slice
}

func GenSliceFloat(size int, floatSize float64) []float64 {
	var slice []float64
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Float64()*floatSize)
	}
	return slice
}

func GenSliceEmpty(size int) []int {
	return make([]int, 100)
}

func GenSliceNil() []int {
	return nil
}

func IsSorted[O Ordered](slice []O) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}

// The following was created by "icza" and modified by "moorara"
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandString(n int) (randStr string) {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
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
