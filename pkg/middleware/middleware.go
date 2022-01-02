package middleware

import (
	"context"
	"eduapp/pkg/handler"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func Middleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return handler.SECRET_KEY, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				next(w, r.WithContext(ctx), ps)
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	}
}
