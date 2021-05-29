package authorize

import (
	"github.com/milangrahovac/geekbrains-go-level2-lesson3/example_app/authorize/token"
)

func Authorize() *token.Token {
	tok := token.New()
	// write token to base
	return tok
}
