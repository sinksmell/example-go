package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", fooHandler)
	http.ListenAndServe(":23333", nil)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	go func (){
		fmt.Fprintf(w, "hello ,%s", r.Method)
	}()
}
