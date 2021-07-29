package handlers

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)



func PubSubNotify(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		log.Printf("Error parsing body: %v", err)
		http.Error(w, "can't parse body", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "samyn-project4")
	if err != nil {
		http.Error(w, "can't create pubsub client:" + err.Error(), http.StatusBadRequest)
		return
	}
	defer client.Close()

	id := ""
	t := client.Topic("hello-function1")
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(data["name"].(string)),
	})

	<- result.Ready()
	id, err = result.Get(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := fmt.Sprintf("From PubSubNotify. id=%s body=%v", id, data)
	_, _ = w.Write([]byte(res))
	return
}

func PubSubNotifyGin(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		c.String(http.StatusBadRequest, "can't read body")
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		log.Printf("Error parsing body: %v", err)
		c.String(http.StatusBadRequest, "can't parse body")
		return
	}

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "samyn-project4")
	if err != nil {
		c.String(http.StatusBadRequest, "can't create pubsub client:" + err.Error())
		return
	}
	defer client.Close()

	id := ""
	t := client.Topic("hello-function1")
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(data["name"].(string)),
	})

	<- result.Ready()
	id, err = result.Get(ctx)
	if err != nil {
		c.String(http.StatusBadRequest, "can't create pubsub client:" + err.Error())
		return
	}

	res := fmt.Sprintf("From PubSubNotify. id=%s body=%v", id, data)
	c.String(http.StatusOK, res)
	return
}
