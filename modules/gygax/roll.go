package gygax

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/hostables/botrocity/responses"
	"github.com/julienschmidt/httprouter"
)

func flipCoin() string {
	if rand.Intn(2) == 0 {
		return "heads"
	}
	return "tails"
}

func rollDie(nSides int) string {
	return strconv.Itoa(rand.Intn(nSides) + 1)
}

func parseTextForDice(message string) []int {
	words := strings.Split(strings.ToLower(message), " ")
	var ret []int
	for _, untrimmedWord := range words {
		word := strings.Trim(untrimmedWord, ".?,;/!'\"{}[](&*%)")
		switch word {
		case "coin":
			ret = append(ret, 2)
		case "d4":
			ret = append(ret, 4)
		case "d6":
			ret = append(ret, 6)
		case "d8":
			ret = append(ret, 8)
		case "d10":
			ret = append(ret, 10)
		case "d12":
			ret = append(ret, 12)
		case "d16":
			ret = append(ret, 16)
		case "d20":
			ret = append(ret, 20)
		case "d32":
			ret = append(ret, 32)
		case "d64":
			ret = append(ret, 64)
		case "d100":
			ret = append(ret, 100)
		}
	}
	return ret
}

func getDiceRollMessage(dice []int) string {
	text := ""
	for i, die := range dice {
		// Add a newline before 2nd-last lines
		if i != 0 {
			text += "\n"
		}

		if die == 2 {
			text = text + "**Flipping coin:** " + flipCoin()
		} else {
			text = text + "**Rolling " + strconv.Itoa(die) + "-sided die:** " + rollDie(die)
		}
	}
	return text
}

// HandleDiceRollText provides an HTTP handler to respond to dice rolls
func HandleDiceRollText(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	text := ""
	resp := responses.MattermostTextResponse{
		Text:     text,
		Username: "Gary Gygax",
	}
	resp.Write(w)
}
