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
	ctx := r.Context()// creating a new context.Context

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

// 	// listenandserve is blocking call: program won't continue running until after listenandsever finishes running or http server is told to shut down

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

	//multiple servers
	ctx, cancelCtx := context.WithCancel(context.Background())// background context is a nono-nil empty context, usually used as the starting point to create any new context => returns a context and a CancelFunc => calling CancelFunc will send the cancellation signal
	serverOne := &http.Server{
		Addr: ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context{
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	serverTwo := &http.Server{
		Addr: ":4444",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context{
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	go func(){
		err := serverOne.ListenAndServe()// parameters not provided as before because http.server has already been configured
		if errors.Is(err, http.ErrServerClosed){
			fmt.Println("server one closed\n")
		}else if err != nil{
			fmt.Println("error listening for server one: %s\n", err)
		}
		cancelCtx()// if server ends for some reason, context will end as well
	}()

	go func(){
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed){
			fmt.Println("server two closed\n")
		}else if err != nil{
			fmt.Println("error listening for server two %\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()// program will stay running until eithere of the server goroutines ends and cancelcts is called
	// once context is over, program will exit

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}


}
