package main

import "fmt"
import "os"

func main(){
	var num float64
	fmt.Fscan(os.Stdin, &num)
	fmt.Println(int(num))
}