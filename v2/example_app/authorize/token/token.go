package token

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	Value      string
	Expiration time.Time
}

func New() *Token {
	return &Token{
		Value:      uuid.NewString(),
		Expiration: time.Now().Add(time.Hour),
	}
}
