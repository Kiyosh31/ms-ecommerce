package middlewares

import (
	"github.com/Kiyosh31/ms-ecommerce-common/token"
)

type Middleware struct {
	tokenCreator *token.JwtCreator
}

func NewMiddleware(secretKey string) *Middleware {
	return &Middleware{
		tokenCreator: token.NewJwtCreator(secretKey),
	}
}
