package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Guessing game start!")

	// rand.Seed(time.Now().Unix())
	secret := rand.Intn(100) + 1
	var guess int
	for {
		fmt.Println("Guess the number:- ")
		fmt.Scan(&guess)

		if guess < secret {
			fmt.Println("Less than secret number")
		} else if guess > secret {
			fmt.Println("More than secret number")

		} else {
			fmt.Println("Congratulations")
			break

		}
	}
}
