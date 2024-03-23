package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//inspecting a request's query string
		//query string: set of values added to end of url
		// starts with ? and additional vlues using &
		// used to filter
		hasFirst := r.URL.Query().Has("first")//r.url field of getroot's *http.request to access properties about url
		first := r.URL.Query().Get("first")//query method of r.url to access query string values
		hasSecond := r.URL.Query().Has("second")// has returns a bool, whether query string  has a value with key provided, eg. first
		second := r.URL.Query().Get("second")// get returns a string with values of key provided

		fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s\n",
		ctx.Value(keyServerAddr),
		hasFirst, first,
		hasSecond, second)

		io.WriteString(w, "this is my website\n")
	}

	func getHello(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
	
		fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
		io.WriteString(w, "hello, http\n")
	}

	func main(){
		mux := http.NewServeMux()
		mux.HandleFunc("/", getRoot)
		mux.HandleFunc("/hello", getHello)

		ctx := context.Background()
		server := &http.Server{
			Addr: ":3333",
			Handler: mux,
			BaseContext: func(l net.Listener) context.Context{
				ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
				return ctx
			},
		}

		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed){
			fmt.Println("server closed\n")
		}else if err != nil{
			fmt.Println("error listening for server: %s\n", err)
		}
	}
// request: curl 'http://localhost:3333?first=1&second='
// server ouput: [::]:3333: got / request. first(true)=1, second(true)=
//