package server

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/helper"
	"net/http"
)

func ExecuteSvcHttp(conf *Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helper.Json(w, http.StatusOK, map[string]interface{}{
			"response": conf.Name + " ok",
		})
	})

	http.HandleFunc("/healthz", helper.OkHandler)

	fmt.Println("listen :", conf.HttpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", conf.HttpPort), nil)
	if err != nil {
		panic(err)
	}
}
