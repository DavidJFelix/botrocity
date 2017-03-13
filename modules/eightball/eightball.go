package eightball

import (
	"math/rand"
	"net/http"

	"github.com/hostables/botrocity/responses"
	"github.com/julienschmidt/httprouter"
)

// getFortune fetches a random fortune from the standard eightball
// response set
func getFortune() string {
	fortunes := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes, definitely",
		"You may rely on it",
		"As I see it, yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy, try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	// Get a random index and return the value at that index
	return fortunes[rand.Intn(len(fortunes))]
}

// HandleMagicEightballText provides an HTTP handler to respond
// with Magic eight ball fortunes via a TextResponse
func HandleMagicEightballText(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := responses.MattermostTextResponse{
		Text:     getFortune(),
		Username: "Magic 8 ball",
		IconURL:  "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/8_ball_icon.svg/120px-8_ball_icon.svg.png",
	}
	resp.Write(w)
}
