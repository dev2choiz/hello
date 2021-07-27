package main

import (
	"encoding/json"
	"fmt"
	"github.com/dev2choiz/hello/pkg/handlers"
	"github.com/dev2choiz/hello/pkg/helloworld"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := helloworld.Say()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(data)
	})

	http.HandleFunc("/notify/function1", handlers.PubSubNotify)

	f := func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		if err != nil {
			panic(err)
		}
	}
	http.HandleFunc("/healthz", f)
	http.HandleFunc("/ready", f)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	fmt.Println("listen port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
