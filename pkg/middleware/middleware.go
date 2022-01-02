package middleware

import (
	Types "eduapp/CommonTypes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var jwtKey = []byte("my_secret_key")

func Middleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tknStr := r.Header.Get("authtoken")
		if len(tknStr) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		claims := &Types.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r.WithContext(r.Context()), ps)
	}
}
