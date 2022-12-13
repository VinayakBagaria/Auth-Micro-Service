package middlewares

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/VinayakBagaria/auth-micro-service/api/restutil"
	"github.com/VinayakBagaria/auth-micro-service/security"
)

func LogRequests(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next(w, r)
		log.Printf(`{"proto": "%s", "method": "%s", "route": "%s%s", "request_time": "%v"}`,
			r.Proto, r.Method, r.Host, r.URL.Path, time.Since(t))
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		splitted := strings.Split(header, " ")
		if len(splitted) != 2 {
			restutil.WriteError(w, http.StatusUnauthorized, restutil.ErrUnauthorized)
			return
		}

		tokenString := splitted[1]
		token, err := security.ParseToken(tokenString)
		if err != nil || !token.Valid {
			restutil.WriteError(w, http.StatusUnauthorized, restutil.ErrUnauthorized)
			return
		}

		next(w, r)
	}
}
