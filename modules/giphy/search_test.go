package giphy

import (
	"testing"
	"log"
)

func TestSearchGiphyApi(t *testing.T) {
	log.Println(searchGiphyAPI([]string{"george","costanza"})[0].Images.Original.URL)
}
