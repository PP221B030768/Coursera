package main

import "fmt"

type Animal struct {
	animal string
	eat    string
	move   string
	speak  string
}

func Readn(arr []string, i, n int) {
	if n == 0 {
		return
	}
	if m, err := fmt.Scan(&arr[i]); m != 1 {
		panic(err)
	}
	Readn(arr, i+1, n-1)
}

func main() {
	var cow = Animal{animal: "cow", eat: "grass", move: "walk", speak: "moo"}
	var bird = Animal{animal: "bird", eat: "worms", move: "fly", speak: "peep"}
	var snake = Animal{animal: "snake", eat: "mice", move: "slither", speak: "hsss"}
	var done bool = true
	for done == true {
		arr := make([]string, 2)
		Readn(arr, 0, 2)
		if arr[0] == "cow" {
			if arr[1] == "eat" {
				fmt.Println(cow.eat)
			} else if arr[1] == "move" {
				fmt.Println(cow.move)
			} else if arr[1] == "speak" {
				fmt.Println(cow.speak)
			}
		} else if arr[0] == "bird" {
			if arr[1] == "eat" {
				fmt.Println(bird.eat)
			} else if arr[1] == "move" {
				fmt.Println(bird.move)
			} else if arr[1] == "speak" {
				fmt.Println(bird.speak)
			}
		} else if arr[0] == "snake" {
			if arr[1] == "eat" {
				fmt.Println(snake.eat)
			} else if arr[1] == "move" {
				fmt.Println(snake.move)
			} else if arr[1] == "speak" {
				fmt.Println(snake.speak)
			}
		}
	}
}
