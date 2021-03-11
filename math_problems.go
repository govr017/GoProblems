package main

import (
	"fmt"
	"math/rand"
	"time"
)

const COLOR_DEFAULT = "\033[0m"
const COLOR_RED = "\033[31m"
const COLOR_GREEN = "\033[32m"
const COLOR_YELLOW = "\033[33m"
const COLOR_BLUE = "\033[34m"

func main() {
	rand.Seed(time.Now().UnixNano())
	var answer int
	var x = rand.Intn(10)
	var y = rand.Intn(10)
	var method int

	fmt.Print("Choose Math-method: minus (1) or plus (2): ", COLOR_YELLOW)
	fmt.Scanln(&method)
	fmt.Print(COLOR_DEFAULT)

	if method == 1 {
		fmt.Printf("Whats: %d - %d = %s", x, y, COLOR_YELLOW)
		fmt.Scanln(&answer)
		if answer == x-y {
			fmt.Print(COLOR_GREEN, "Good job!", COLOR_DEFAULT, "\n")
		} else {
			fmt.Print(COLOR_RED, "Too bad!", COLOR_DEFAULT, "\n")
		}
	} else {
		fmt.Printf("Compute: %d + %d = %s", x, y, COLOR_YELLOW)
		fmt.Scanln(&answer)
		if answer == x+y {
			fmt.Print(COLOR_GREEN, "Good job!", COLOR_DEFAULT, "\n")
		} else {
			fmt.Print(COLOR_RED, "Too bad!", COLOR_DEFAULT, "\n")
		}
	}
}
