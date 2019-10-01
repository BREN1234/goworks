package datastructure

import (
	"math/rand"
	"time"
)

func SelectionSort(list []int) []int {
	var length int = len(list)
	for i := 0; i < length; i++ {
		var minIndex int = i
		for j := i; j < length; j++ {
			if list[j] < list[i] {
				minIndex = j
			}
		}
		list[i], list[minIndex] = list[minIndex], list[i]
	}
	return list
}
func RandomNumberGenerator(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) // - rand.Intn(999)
	}
	return slice
}
