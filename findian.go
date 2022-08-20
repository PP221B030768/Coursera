package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	var str string
	fmt.Fscan(os.Stdin, &str)
	if (string(str[0]) == string('i') || string(str[0]) == string('I')) && (string(str[len(str)-1]) == string('n') || string(str[len(str)-1]) == string('N')) && (strings.Contains(str, string('a')) || strings.Contains(str, string('A'))) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
	i, _ := strconv.Atoi("10")
  y := i * 2
  fmt.Println(y)
  x:=7
  switch {
    case x>3:
      fmt.Printf("1")
    case x>5:
      fmt.Printf("2")
    case x==7:
      fmt.Printf("3")
    default: 
      fmt.Printf("4")
  }
  var xtemp int
  x1 := 0
  x2 := 1
  for x:=0; x<5; x++ {
    xtemp = x2
    x2 = x2 + x1
    x1 = xtemp
  }
  fmt.Println(x2)
}
