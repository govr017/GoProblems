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
const VAR_MAX_DEFAULT = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	var answer int
	var var_max int // maximum value of each number
	var x int
	var y int
	var method int

	fmt.Printf("Set MAXIMUM number for values (default = %d): %s", VAR_MAX_DEFAULT, COLOR_YELLOW)
	_, err := fmt.Scanln(&var_max)
	if err != nil {
		var_max = VAR_MAX_DEFAULT
		fmt.Print("var_max = ", var_max)
	}
	fmt.Println(COLOR_DEFAULT)
	x = rand.Intn(var_max)
	y = rand.Intn(var_max)

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
