package middleware

import (
	"log"
	"net/http"
)

func accessMiddleWare(handler http.Handler) http.Handler {
	log.Println("~router.accessMiddleWare ~~~~~")
	trustedRoutes := map[string]int{
		"/api/login":       0,
		"/api/register":    0,
		"/api/getmessages": 0,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* some checks? */
		route := r.URL.Path
		_, trustedRoutExist := trustedRoutes[route]
		if trustedRoutExist {
			handler.ServeHTTP(w, r)
		} else {
			//checkJWT
			authTokenHeader := r.Header.Get("Authorization")
			if len(authTokenHeader) > 0 {
				checkRes := checkToken(authTokenHeader, signingKey)
				if checkRes == nil {
					handler.ServeHTTP(w, r)
				} else {
					log.Println("ckeck tocken rison: ", checkRes)
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("TokenExpired"))
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("NoToken"))
			}
		}
	})
}
