package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
)

var path string = "signin.html"

func ourHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles(path))
	tpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	params := u.Query()
	LoginKey := params.Get("login")
	PassKey := params.Get("pass")
	/*if page == "" {
		page = "1"
	}*/

	fmt.Println("Login is: ", LoginKey)
	fmt.Println("Pass is: ", PassKey)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	//path = "signin.html"
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/signin", searchHandler)
	mux.HandleFunc("/", ourHandler)
	http.ListenAndServe(":"+port, mux)
}
