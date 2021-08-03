package main

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/helper"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helper.Json(w, http.StatusOK, map[string]interface{}{
			"response": "svc1 ok",
		})
	})

	http.HandleFunc("/healthz", helper.OkHandler)

	p := "8081"
	fmt.Println("listen :", p)
	err := http.ListenAndServe(fmt.Sprintf(":%s", p), nil)
	if err != nil {
		panic(err)
	}
}
