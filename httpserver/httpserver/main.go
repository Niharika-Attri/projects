package main

import (
	//"fmt"
	"fmt"
	"html/template"
	"net/http"
)

// function to handle get and post, return error if unsupported
func handleRequest(w http.ResponseWriter, r *http.Request){
	// switch r.Method{
	// case http.MethodGet:
	// 	fmt.Fprint(w, "GET request")
	// case http.MethodPost:
	// 	fmt.Fprint(w, "POST request")
	// default:
	// 	fmt.Fprintf(w, "unsupported method: %s", r.Method)
	// }

	switch r.Method{
	case http.MethodGet:
		tmpl, err := template.ParseFiles("templates/index.html")
		if err!= nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	default:
		fmt.Fprintf(w, "unsupported method: %s", r.Method)
	}
}

func main(){
	http.HandleFunc("/", handleRequest)

	http.ListenAndServe(":8800", nil)
}