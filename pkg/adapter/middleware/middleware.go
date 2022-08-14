package middleware

import (
	"fmt"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		//jaeger.LogRequest(r.Header)

		next.ServeHTTP(w, r)
	})
}
