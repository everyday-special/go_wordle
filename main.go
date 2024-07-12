package main

import (
	"fmt"

	"github.com/everyday-special/go_wordle/guess"
	"github.com/everyday-special/go_wordle/letterbank"
)

func main() {
	// Testing for Letterbank and Guess structs
	// TODO - main game loop
	lb := letterbank.NewLetterbank()
	lb.Print()
	var new_guess string
	fmt.Scanln(&new_guess)
	current_guess := guess.NewGuess(new_guess)
	// TODO - generate secret word from list in file
	current_guess.Check("irate")
	current_guess.Print()
	lb.Update(current_guess.Letter_colors)
	lb.Print()

	// TODO - validate guess
	fmt.Scanln(&new_guess)
	current_guess = guess.NewGuess(new_guess)
	current_guess.Check("irate")
	current_guess.Print()
	lb.Update(current_guess.Letter_colors)
	lb.Print()
}
