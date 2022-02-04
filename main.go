package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)


func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func sawoSDK(w http.ResponseWriter, filename string, data interface{}) {
    t, err := template.ParseFiles(filename)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    if err := t.Execute(w, data); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
	configMap := map[string]interface{}{"apiKey": "0f6a660b-edc8-48f0-905d-359da3f18d4c", "identifier_type": "email"}
	sawoSDK(w, "assets/index.html", configMap)
}