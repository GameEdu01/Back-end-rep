package middleware

import (
	Types "eduapp/CommonTypes"
	myerrors "eduapp/pkg/errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var jwtKey = []byte("my_secret_key")

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cookie, err := r.Cookie("authToken")

		if err != nil {
			fmt.Println("Error occured while reading cookie")
			w.WriteHeader(http.StatusUnauthorized)
			myerrors.Handle401(w, r)
			return
		}
		if len(cookie.Value) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			myerrors.Handle401(w, r)
			return
		}
		claims := &Types.Claims{}
		tkn, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				myerrors.Handle401(w, r)
				return
			}
			fmt.Println(err.Error())
			fmt.Println(tkn)
			w.WriteHeader(http.StatusBadRequest)
			myerrors.Handle400(w, r)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			myerrors.Handle401(w, r)
			return
		}
		next(w, r.WithContext(r.Context()), ps)
	}
}
