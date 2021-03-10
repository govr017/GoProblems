package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const COLORRED = "\033[31m"
	const COLORGREEN = "\033[32m"

	rand.Seed(time.Now().UnixNano())
	var answer int
	var x = rand.Intn(10)
	var y = rand.Intn(10)
	var method int


  fmt.Println("Hey choose minus(1) or plus(2) 1/2?")
	fmt.Scanln(&method)


  if method == 1 {
		fmt.Println("Whats", x, "", "-", "", y)
		fmt.Scanln(&answer)
		if answer == x-y {
			fmt.Println(string(COLORGREEN), "Good job!")
		} else {
			fmt.Println(string(COLORRED), "Too bad!")
		}
	} else {
		fmt.Println("Whats", x, "", "+", "", y)
		fmt.Scanln(&answer)
		if answer == x+y {
			fmt.Println(string(COLORGREEN), "Good job!")
		} else {
			fmt.Println(string(COLORRED), "Too bad!")
		}
	}

}
