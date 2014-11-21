package middleware

import (
	. "github.com/meatballhat/negroni-logrus"
)

func Logrus() *Middleware {
	return NewMiddleware()
}