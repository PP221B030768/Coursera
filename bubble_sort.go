package main

import "fmt"

func main() {
	fmt.Println(`Enter the number of integers`)
	var n int
	if m, err := fmt.Scan(&n); m != 1 {
		panic(err)
	}
	fmt.Println(`Enter the integers`)
	nums := make([]int, n)
	ReadN(nums, 0, n)
	BubbleSort(nums)
	fmt.Println(nums)
}

func ReadN(arr []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := fmt.Scan(&arr[i]); m != 1 {
		panic(err)
	}
	ReadN(arr, i+1, n-1)
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			Swap(arr, j)
		}
	}
}

func Swap(input []int, j int) {
    if input[j] > input[j+1] {
        input[j], input[j+1] = input[j+1], input[j]
    }
}
