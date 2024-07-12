package guess

import (
	"fmt"

	"github.com/everyday-special/go_wordle/colors"
	"github.com/everyday-special/go_wordle/letter"
)

type guess struct {
	word          string
	Letter_colors [5]letter.Letter
}

func NewGuess(new_guess string) *guess {
	var letter_colors [5]letter.Letter
	for i, ch := range new_guess {
		letter_colors[i] = letter.Letter{
			Ch:    ch,
			Color: colors.WHITE,
		}
	}

	return &guess{
		Letter_colors: letter_colors,
		word:          new_guess,
	}
}

func (guess guess) Print() {
	for _, l := range guess.Letter_colors {
		fmt.Printf(l.Color+"%c"+colors.WHITE, l.Ch)
	}
	fmt.Println()
}

func (guess *guess) Check(secret_word string) bool {
	if guess.word == secret_word {
		for i := range guess.Letter_colors {
			guess.Letter_colors[i].Color = colors.GREEN
		}
		return true
	}

	var secret_word_ch_counts map[rune]int = map[rune]int{}
	// Get letter counts from secret_word
	for _, ch := range secret_word {
		count := secret_word_ch_counts[ch]
		secret_word_ch_counts[ch] = count + 1
	}

	// Check for exact matches
	for i, ch := range secret_word {
		if rune(guess.word[i]) == ch {
			secret_word_ch_counts[ch] -= 1
			guess.Letter_colors[i].Color = colors.GREEN
		}
	}

	// Check for partial matches
	for i, letter := range guess.Letter_colors {
		if secret_word_ch_counts[letter.Ch] > 0 && letter.Color != colors.GREEN {
			secret_word_ch_counts[letter.Ch] -= 1
			guess.Letter_colors[i].Color = colors.YELLOW
		}
	}

	return false
}
