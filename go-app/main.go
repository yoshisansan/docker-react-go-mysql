package main

import (
	"net/http"

	"github.com/justinas/alice"
	"local.pkg/middleware"
	// "local.pkg/sample"
)

func myApp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}


func main() {
	corsOption := middleware.Cors{
		AllowOrigin: "*",
		AllowCredentials: true,
		AllowMethod: []string{"POST, GET, OPTIONS, DELETE, PUT"},
		AllowHeaders: []string{"Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Cookie"},
	}
	oh := http.HandlerFunc(myApp)
	commonHandlers := alice.New(middleware.CorsHandler(corsOption))
	http.Handle("/", commonHandlers.Then(oh))
	http.ListenAndServe(":8080", nil)
}