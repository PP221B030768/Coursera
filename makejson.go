package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/json"
)

func main(){

	fmt.Printf("Please, enter your name: ")
	
	input1 := bufio.NewReader(os.Stdin)
    name, err := input1.ReadString('\n')
    if err != nil {
    	fmt.Println("Ошибка ввода: ", err)
	}

	fmt.Printf("Please, enter your address: ")

	input2 := bufio.NewReader(os.Stdin)
	address, err := input2.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}

	mp := map[string]string{
		"name":    name, 
		"address": address,
	}

	js, _ := json.MarshalIndent(mp, "", " ")

	fmt.Println(string(js))
}