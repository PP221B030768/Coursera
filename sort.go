package main

import (
	"fmt"
	"sort"
)

func readArray() []int {
	var n, i int
	fmt.Println("Enter the number of inputs")
	_, _ = fmt.Scanln(&n)

	fmt.Println("Enter the inputs")
	a := make([]int, n)
	for i = 0; i < n; i++ {
		_, _ = fmt.Scanln(&a[i])
	}

	return a
}

func partition(a []int, arraySize int) [][]int {
	var array [][]int

	for arraySize < len(a) {
		a, array = a[arraySize:], append(array, a[0:arraySize:arraySize])
	}
	return append(array, a)
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
	a := readArray()
	var arraySize int

	if len(a)%4 > 0 {
		arraySize = len(a)/ 4 + 1
	} else {
		arraySize = len(a)/ 4
	}

	array := partition(a, arraySize)
	arrayCount := len(array)
	fmt.Printf("array: %v\n", array)

	var ch = make(chan []int, arraySize)
	mapSorting(arrayCount, array, ch)
	reducedArray := reduceSorting(arrayCount, ch)

	sortArray(reducedArray, ch)
	fullySorted := <-ch

	fmt.Printf("Result: %v", fullySorted)
}