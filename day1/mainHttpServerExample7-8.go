package main

import (
	"net/http"
)
//7-Http server
func main() {
	//8-Handler
	http.HandleFunc("/", home)
	//7-Http server
	http.ListenAndServe(":3000", nil) // every each new connection generate a go routine(thread) with 2kb.
	//Other examples of servers
	//grpc
	//tcpAddr
}

//8-Handler
func home(w http.ResponseWriter, r*http.Request) {
	w.Write([]byte("Hello World!!!"))
}