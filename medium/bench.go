package main

import (
	"math/rand"
)

func generateArray(size int) []int {
	var arr = make([]int, size)
	index := 0
	for index < size {
		arr[index] = rand.Intn(size)
		index++
	}
	return arr
}

func swap(arr []int, l, r int) {
	arr[l], arr[r] = arr[r], arr[l]
}

func bubleSort(arr []int) []int {
	length := len(arr)
	swapped := false
	for i := 0; i < length; i++ {
		swapped = false
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
	return arr
}

func main() {
	var sorted = bubleSort(generateArray(10000))
}
