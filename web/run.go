package web

import (
	"fmt"
	"html"
	"net/http"

	"github.com/pdxjohnny/easysock"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Run() {
	mux := http.NewServeMux()
	go easysock.Hub.Run()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/ws", easysock.ServeWs)
	Start(
		mux,
		"0.0.0.0",
		"14000",
		"",
		"key",
	)
}
