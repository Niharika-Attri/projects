package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

const keyServerAddr = "serverAddr" // acts as the key for http server's address value in the http.request context

func getRoot(w http.ResponseWriter, r *http.Request) { // http.HandlerFunc
	ctx := r.Context()

	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))

	//fmt.Printf("got / request\n")             // w (interface) is used to conrol the response info being written back to the client that made reqyest(body, status code)
	io.WriteString(w, "this is my website\n") // r is used to get info about the req that came to server(eg. post req or info about client)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
	//fmt.Printf("got /hello request\n")
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

	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr: ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context{
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
