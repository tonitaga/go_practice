package main

import "net/http"

func main() {
	http.HandleFunc("/", serveHelloWorldFile)
	http.ListenAndServe("localhost:8888", nil)
}

func serveHelloWorldFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	http.ServeFile(w, r, "html/index.html")
}
