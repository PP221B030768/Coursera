package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	GName() string
}

type cow struct {
	name string
}

type snake struct {
	name string
}

type bird struct {
	name string
}

func (c cow) Eat() {
	fmt.Printf("%s eats grass\n", c.name)
}

func (s snake) Eat() {
	fmt.Printf("%s eats mice\n", s.name)
}

func (b bird) Eat() {
	fmt.Printf("%s eats worms\n", b.name)
}

func (c cow) Move() {
	fmt.Printf("%s walks\n", c.name)
}

func (s snake) Move() {
	fmt.Printf("%s slithers\n", s.name)
}

func (b bird) Move() {
	fmt.Printf("%s flies\n", b.name)
}

func (c cow) Speak() {
	fmt.Printf("%s moos\n", c.name)
}

func (s snake) Speak() {
	fmt.Printf("%s hssses\n", s.name)
}

func (b bird) Speak() {
	fmt.Printf("%s peeps\n", b.name)
}

func (c cow) GName() string {
	return c.name
}

func (s snake) GName() string {
	return s.name
}

func (b bird) GName() string {
	return b.name
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

	sl := []Animal{}

	fmt.Print(">")

	var done bool = true
	for done == true {
		arr := make([]string, 3)
		Readn(arr, 0, 3)
		if arr[0] == "newanimal" {
			if arr[2] == "cow" {
				sl = append(sl, cow{name: arr[1]})
				fmt.Println("Created it!")
			} else if arr[2] == "snake" {
				sl = append(sl, snake{name: arr[1]})
				fmt.Println("Created it!")
			} else if arr[2] == "bird" {
				sl = append(sl, bird{name: arr[1]})
				fmt.Println("Created it!")
			}
		} else if arr[0] == "query" {
			for _, anim := range sl {
				if anim.GName() == arr[1] {
					if arr[2] == "eat" {
						anim.Eat()
					} else if arr[2] == "move" {
						anim.Move()
					} else if arr[2] == "speak" {
						anim.Speak()
					}
				}
			}
		}
	}
}
