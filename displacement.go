package main

import (
	"fmt"
	"math"
)

func main(){
	var n int = 3
	nums := make([]float64, n)
	fmt.Println("Please, enter values for acceleration, initial velocity, and initial displacement:")
	Read(nums, 0, n)
	acceleration := nums[0]
	init_velocity := nums[1]
	init_displacement := nums[2]
	var time float64
	fmt.Println("Please, enter value for time:")
	fmt.Scan(&time)
	fn := GenDisplaceFn(float64(acceleration), float64(init_velocity), float64(init_displacement))
	dis := fn(time)
	fmt.Println(dis)
}

func Read(arr []float64, i, n int){
	if n == 0{
		return
	}
	if m, err := fmt.Scan(&arr[i]); m != 1{
		panic(err)
	}
	Read(arr, i+1, n-1)
}

func GenDisplaceFn(acceleration, initial_velocity, initial_displacement float64) func (float64) float64{
	function := func(time float64) float64{
		return initial_displacement + initial_velocity * time + acceleration * 0.5 * math.Pow(time, 2)
	}
	return function
}