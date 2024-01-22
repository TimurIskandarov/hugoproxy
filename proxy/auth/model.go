package auth

import "github.com/go-chi/jwtauth"

var (
	TokenAuth *jwtauth.JWTAuth
	users     = make(map[string]string)
)

type User struct {
	Name     string
	Password string
}

type ResponseRegister struct {
	Token string
}

type ResponseLogin struct {
	Token   string
}
