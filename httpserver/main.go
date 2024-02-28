package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) { // http.HandlerFunc
	fmt.Printf("got / request\n")             // w is used to conrol the response info being written back to the client that made reqyest(body, status code)
	io.WriteString(w, "this is my website\n") // r is used to get info about the req that came to server(eg. post req or info about client)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "hello, http\n") // using w to send some text to response body
}

// func main() {
// 	http.HandleFunc("/", getRoot)
// 	http.HandleFunc("/hello", getHello)

// 	err := http.ListenAndServe(":3333", nil) // tells global http server to listen for incooming requests on specific port with optional htt.handler
// 	// nil value for http.handler: tells listenandserve that you want to use the default server multiplexer and not the one you've set up

// 	// listenandserve is blocking call: program won't continue running until after listenandsever finishes running

// 	if errors.Is(err, http.ErrServerClosed) {
// 		fmt.Printf("server closed\n")
// 	} else if err != nil {
// 		fmt.Printf("error starting server: %s\n", err)
// 		os.Exit(1)
// 	}
// }

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
