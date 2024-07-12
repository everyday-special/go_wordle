package letterbank

import (
	"fmt"

	"github.com/everyday-special/go_wordle/colors"
	"github.com/everyday-special/go_wordle/letter"
)

type letterbank struct {
	letterbank [26]letter.Letter
}

func NewLetterbank() *letterbank {
	var alphabet [26]letter.Letter
	for r := 'a'; r <= 'z'; r++ {
		alphabet[r-97] = letter.Letter{
			Ch:    r,
			Color: colors.WHITE,
		}
	}
	return &letterbank{
		letterbank: alphabet,
	}
}

func (lb *letterbank) updateColor(ch rune, color string) {
	curr_color := lb.letterbank[ch-97].Color
	if curr_color == colors.GREEN || curr_color == colors.BLACK {
		// These are "terminal" colors in the letterbank
		return
	} else if curr_color == colors.YELLOW && color != colors.GREEN {
		// only valid letterbank state change for yellow is to green
		return
	}
	// This should only allow changes from white to any color
	// or yellow to green.
	lb.letterbank[ch-97].Color = color
}

func (lb *letterbank) Print() {
	for i := 0; i < 26; i++ {
		fmt.Printf(lb.letterbank[i].Color+"%c "+colors.WHITE, lb.letterbank[i].Ch)
	}
	fmt.Println()
}

func (lb *letterbank) Update(guess_letter_colors [5]letter.Letter) {
	for _, l := range guess_letter_colors {
		lb.updateColor(l.Ch, l.Color)
	}
}
