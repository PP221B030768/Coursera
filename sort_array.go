package main

import (
	"fmt"
	"sort"
)

func readArray() []int {
	var n int
	fmt.Println("Enter the number of inputs:")
	_, _ = fmt.Scanln(&n)

	fmt.Println("Enter the inputs:")
	myarr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanln(&myarr[i])
	}
	return myarr
}

func partition(a []int, size int) [][]int {
	var arr [][]int

	for size < len(a) {
		arr = append(arr, a[0:size])
		a = a[size:]
	}
	return append(arr, a)
}

func sortArray(a []int, c chan []int) {
	sort.Ints(a)
	c <- a
}

func mapSorting(arrayCount int, array [][]int, ch chan []int) {
	for i := 0; i < arrayCount; i++ {
		go sortArray(array[i], ch)
	}
}

func reduceSorting(arrayCount int, ch chan []int) []int {
	reducedArray := make([]int, 0)
	for i := 0; i < arrayCount; i++ {
		var sorted = <-ch
		reducedArray = append(reducedArray, sorted...)
	}
	return reducedArray
}

func main() {
	myarr := readArray()

	var size int

	if len(myarr)%4 != 0 {
		size = len(myarr)/4 + 1
	} else {
		size = len(myarr)/4
	}
	array := partition(myarr, size)
	arrayCount := len(array)
	fmt.Printf("array: %v\n", array)

	var ch = make(chan []int, size)
	mapSorting(arrayCount, array, ch)
	reducedArray := reduceSorting(arrayCount, ch)

	sortArray(reducedArray, ch)
	fullySorted := <-ch

	fmt.Printf("Result: %v", fullySorted)
}