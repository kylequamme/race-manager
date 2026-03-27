package main

import webview "github.com/webview/webview_go"

func startWebView() {
	// Create and launch UI
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Race Manager")
	w.SetSize(1280, 720, webview.HintNone)
	w.Navigate("http://127.0.0.1:9090")
	// For testing webview with local dev
	// w.Navigate("http://127.0.0.1:3000")
	w.Run()
}
