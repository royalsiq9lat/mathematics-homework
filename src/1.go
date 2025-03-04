package main

import "fmt"

func main() {
	randomNumber := rand.Intn(10)
	if randomNumber%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}
