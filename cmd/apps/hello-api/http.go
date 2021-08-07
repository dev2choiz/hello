package main

import (
	"encoding/json"
	"fmt"
	"github.com/dev2choiz/hello/pkg/handlers"
	"github.com/dev2choiz/hello/pkg/helper"
	"github.com/dev2choiz/hello/pkg/server"
	"github.com/dev2choiz/hello/pkg/version"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func executeApiHttp(conf *server.Config) {
	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{}

		res, err := getSvcData(os.Getenv("SVC1_BASE_URL"))
		if err == nil {
			data["svc1"] = res
		}
		res, err = getSvcData(os.Getenv("SVC2_BASE_URL"))
		if err == nil {
			log.Println(err)
			data["svc2"] = res
		}

		log.Println(data)
		helper.Json(w, http.StatusOK, map[string]interface{}{
			"status": "ok",
			"version": version.Get(),
			"data": data,
		})
	})

	http.HandleFunc("/", helper.OkHandler)
	http.HandleFunc("/notify/function1", handlers.PubSubNotify)

	http.HandleFunc("/healthz", helper.OkHandler)
	http.HandleFunc("/ready", helper.OkHandler)
	http.HandleFunc("/secured", helper.OkHandler)

	fmt.Println("listen :", conf.HttpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", conf.HttpPort), nil)
	if err != nil {
		panic(err)
	}
}

func getSvcData(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var data map[string]string
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		return "", err
	}

	return data["response"], nil
}
