package middleware

import (
	"homework/database"
	"homework/infra/jwt"
)

type Middleware struct {
	cacher database.Cacher
	Jwt    jwt.JWT
}

func NewMiddleware(cacher database.Cacher, jwt jwt.JWT) Middleware {
	return Middleware{cacher, jwt}
}
