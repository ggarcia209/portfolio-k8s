package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// TmplMap maps html template paths to shortnames
// File structure:
// web
//  - html
//    - index.html
// service
//  - serve
//   - serve.go
var TmplMap = map[string]string{
	"index": "./index.html",
}

// InitHTTPServer initializes an http server at the provided address
func InitHTTPServer(addr string) *http.Server {
	srv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         addr,
		Handler:      nil,
	}
	return srv
}

func InitMuxRouter(port int) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home).Methods("GET")
	portStr := fmt.Sprintf(":%d", port)
	fmt.Printf("listening at port: '%s'...\n", portStr)
	log.Fatal(http.ListenAndServe(portStr, router))
}

// RegisterHandlers registers the http handler functions:
//  Home
//  Rankings
//  About
//  SearchResults
// and creates file servers for the css, js, and img files.
func RegisterHandlers() {
	http.HandleFunc("/", Home)
}

// Home displays home page
func Home(w http.ResponseWriter, r *http.Request) {
	if err := templExe(TmplMap["index"], w, nil); err != nil {
		fmt.Fprintf(w, "html template failed to execute: %s", err)
		fmt.Printf("html template failed to execute: %s", err)
		return
	}
	return
}

// templExe executes for the specified template for the given io.Writer and data interface
func templExe(tmpl string, w http.ResponseWriter, data interface{}) error {
	t := template.Must(template.ParseFiles(tmpl))
	if err := t.Execute(w, data); err != nil {
		fmt.Printf("template execution failed: %v", err)
		return fmt.Errorf("template execution failed: %v", err)
	}
	return nil
}
