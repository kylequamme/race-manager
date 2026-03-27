package main

import (
	"net/http"
)

func main() {
	//startApi()
	// For testing webview local dev; comment out line above
	go startApi()
	go startFileServer()
	startWebView()
}

func startFileServer() {
	// Create fileServer
	http.Handle("/", http.FileServer(http.Dir("./frontend/out")))
	http.ListenAndServe(":9090", nil)
}
