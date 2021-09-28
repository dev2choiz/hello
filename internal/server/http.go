package server

import (
	"net/http"
)

// GetServeMux
func GetServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthzHandler)
	return mux
}

func healthzHandler(res http.ResponseWriter, req *http.Request) {
	_, _ = res.Write([]byte("ok"))
}
