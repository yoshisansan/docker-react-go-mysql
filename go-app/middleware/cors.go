package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

// func loggingHandler(next http.Handler) http.Handler {
//   fn := func(w http.ResponseWriter, r *http.Request) {
//     t1 := time.Now()
//     next.ServeHTTP(w, r)
//     t2 := time.Now()
//     log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
//   }

//   return http.HandlerFunc(fn)
// }

// func recoverHandler(next http.Handler) http.Handler {
//   fn := func(w http.ResponseWriter, r *http.Request) {
//     defer func() {
//       if err := recover(); err != nil {
//         log.Printf("panic: %+v", err)
//         http.Error(w, http.StatusText(500), 500)
//       }
//     }()

//     next.ServeHTTP(w, r)
//   }

//   return http.HandlerFunc(fn)
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
//   fmt.Fprintf(w, "About!")
// }

// func indexHandler(w http.ResponseWriter, r *http.Request) {
//   fmt.Fprintf(w, "Welcome!")
// }

// func call() {
// 	setupResponse(&w, req)
// 	commonHandlers := alice.New(loggingHandler, recoverHandler)
// 	http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
// 	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
// 	http.ListenAndServe(":8080", nil)
// }

type Cors struct {
	AllowOrigin string
	AllowCredentials bool
	AllowMethod []string
	AllowHeaders []string
}

func MapString(x []string) string {

	var joinnedStr bytes.Buffer
	for _, e := range x {
			str := fmt.Sprintf("%s,",e)
			joinnedStr.WriteString(str)
	}
	return joinnedStr.String()
}

func JoinStringArr(x []string) string{
	return MapString(x)
}

// func SetupResponse(config Cors)(w http.ResponseWriter) {
func SetupResponse(w http.ResponseWriter, r *http.Request, config Cors) error {
		(w).WriteHeader(200)
		(w).Header().Set("Access-Control-Allow-Origin", config.AllowOrigin)
		(w).Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(config.AllowCredentials))
    (w).Header().Set("Access-Control-Allow-Methods", JoinStringArr(config.AllowMethod))
    (w).Header().Set("Access-Control-Allow-Headers", JoinStringArr(config.AllowHeaders)) // ex. ["Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"]
		fmt.Print("hoge")

		return nil
}

// PUTやDELETEを使用する場合はpreflightに対してstatus okを出す
// そうしないとPUTやDELETEのような非単純リクエストが実行されない
func PreflightCheck(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	return nil
}

func CorsHandler(config Cors) (mw func(http.Handler) http.Handler) {
	mw = func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					PreflightCheck(w, r)
					SetupResponse(w, r, config)
					h.ServeHTTP(w, r)
			})
	}
	return mw
}

