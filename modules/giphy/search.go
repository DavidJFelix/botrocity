package giphy

import (
	"log"
	"net/http"

	"github.com/hostables/botrocity/responses"
	"github.com/julienschmidt/httprouter"
	"github.com/peterhellberg/giphy"
)

var g *giphy.Client

func init() {
	g = giphy.DefaultClient
	g.Rating = "pg"
}

func SearchGiphyAPI(terms []string) ([]giphy.Data) {
	search, err := g.Search(terms)
	if err != nil {
		log.Println("Error: ", err)
		return []giphy.Data{}
	}
	return search.Data
}

// HandleGiphySearchText will respond
func HandleGiphySearchText(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	text := ""
	resp := responses.MattermostTextResponse{
		Text:     text,
		Username: "Giphy",
	}
	resp.Write(w)
}
