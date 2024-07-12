package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"

	"github.com/everyday-special/go_wordle/guess"
	"github.com/everyday-special/go_wordle/letterbank"
)

//go:embed word_list.txt
var SECRET_WORD_LIST string

//go:embed allowed_guesses.txt
var ALLOWED_GUESS_LIST string
var ALLOWED_GUESSES map[string]int = map[string]int{}

func main() {
	// Initialize
	for _, str := range strings.Split(ALLOWED_GUESS_LIST, "\n") {
		ALLOWED_GUESSES[str] = 0
	}
	lb := letterbank.NewLetterbank()
	secret_word := getSecretWord()
	round := 0
	var current_guess *guess.Guess
	var win bool = false
	var history [6]guess.Guess

	for round < 6 && !win {
		// Main game loop
		fmt.Println("\033[2J")
		lb.Print()
		for i := range round {
			history[i].Print()
		}
		current_guess = getUserGuess()
		win = current_guess.Check(secret_word)
		lb.Update(current_guess.Letter_colors)
		history[round] = *current_guess
		round++
	}

	fmt.Println("\033[2J")
	lb.Print()
	for i := range round {
		history[i].Print()
	}
	if win {
		fmt.Println("You win!")
	} else {
		fmt.Printf("You lose. The secret word was %s.", secret_word)
	}
}

func getSecretWord() string {
	secret_word_list := strings.Split(SECRET_WORD_LIST, "\n")
	return secret_word_list[rand.Intn(len(secret_word_list))]
}

func getUserGuess() *guess.Guess {
	var user_guess string
	fmt.Scan(&user_guess)
	_, is_valid_guess := ALLOWED_GUESSES[user_guess]
	for !is_valid_guess {
		fmt.Println("Please enter a valid 5 letter word:")
		fmt.Scan(&user_guess)
		_, is_valid_guess = ALLOWED_GUESSES[user_guess]
	}
	// TODO - confirm guess is a valid word
	return guess.NewGuess(user_guess)
}
