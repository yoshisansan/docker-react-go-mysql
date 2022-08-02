package sample

import (
	"errors"
	"log"
	"net/http"
	"time"

	"local.pkg/api"
)

var euro = map[string]string{
	"2008": "Spain",
	"2012": "Spain",
	"2016": "Portugal",
}

type winner struct {

}

func SampleView(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()
	year := query.Get("year")

	if winner, ok := euro[year]; ok {
		api.JsonResponse(w, map[string]interface{}{"winner": winner})
		return nil
	} else {
		w.WriteHeader(http.StatusNotFound)
		return errors.New("Not Found")
	}
}

func SampleViewHandler(w http.ResponseWriter, r *http.Request)(mw func(http.Handler) http.Handler) {
	SampleView(w, r)
	// mw = func(h http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 			SampleView(w, r)
	// 			h.ServeHTTP(w, r)
	// 	})
	// }
	return
}

// func SampleViewHandler(next http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		SampleView(w,r)
//     next.ServeHTTP(w, r)
//   })
// }

func loggingHandler(next http.Handler) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()
    next.ServeHTTP(w, r)
    t2 := time.Now()
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
  }

  return http.HandlerFunc(fn)
}