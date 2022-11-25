package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"github.com/julienschmidt/httprouter"
)

type CountHandler struct{}

func count(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "0 pin bookmarks currently stored")
}

func singlePin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "presenting pin, %s\n", p.ByName("url"))
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main() {
	mux := httprouter.New()

	mux.GET("/pins/:url", singlePin)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	// http.HandleFunc("/pins/count", log(count))

	server.ListenAndServe()
}
