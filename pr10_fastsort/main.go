package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var inArr []int
	str := "[5, 2, 6, 4, 1, 3]"
	json.Unmarshal([]byte(str), &inArr)

	sortArr := quickSort(inArr, 0, len(inArr) - 1)
	fmt.Println(sortArr)
}

func quickSort(arr []int, s int, e int) []int {
	if e-s+1 <= 1 {
		return arr
	}

	pivot := arr[e] // Значение крайне правое
	left := s // Указатель на ячейку (формально больше pivot)

	for i := s; i <= e; i++ {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left += 1
		}
	}

	arr[e] = arr[left]
	arr[left] = pivot

	quickSort(arr, s, left - 1)
	quickSort(arr, left + 1, e)

	return arr
}
