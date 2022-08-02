package middleware

import "net/http"


func enforceXMLHandler(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    next.ServeHTTP(w, r)
  })
}

func Handler() (mw func(http.Handler) http.Handler) {

	mw = func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					h.ServeHTTP(w, r)
			})
	}
	return
}