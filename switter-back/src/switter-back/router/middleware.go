package router

import "net/http"

func accessMiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* some checks? */
		handler.ServeHTTP(w, r)
	})
}
