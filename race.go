package main

import "fmt"

func by3(x *uint64) {
	*x = *x * 3
}

func by7(x *uint64) {
	*x = *x + 7
}

func main() {
	var x uint64 = 1
	for i := 0; i < 10; i++ {
		go by3(&x)
		go by7(&x)
	}
	fmt.Println(x)
}

// A race condition occurs when multiple threads try to access and modify the same data (memory address).

// This code multiplies the integer by 3 and increments by 7.
// The code is non-deterministic (at the go instruction level) because it is unknown
// whether go will first increment or first double. Therefore, the answer could be
// Scenario 1: (1 * 3) + 7 = 10
// Scenario 2: (1 + 7) * 3 = 24

// Also code could be non-deterministic at the machine level when the value of the
// variable is already loaded by one goroutine but the result has not been written back
// before the second go routine starts. In this case the number is not incremented or
// not multiplied.
