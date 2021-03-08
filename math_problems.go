package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	rand.Seed(time.Now().UnixNano())
	var answer int
	var x = rand.Intn(10)
	var y = rand.Intn(10)

	fmt.Println("Whats", x, "", "+", "", y)
	fmt.Scanln(&answer)
	if answer == x+y {
		fmt.Println(string(colorGreen), "Good job!")
	} else {
		fmt.Println(string(colorRed), "Too bad!")
	}

}
