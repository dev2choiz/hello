package app

import (
	"github.com/dev2choiz/hello/pkg/handlers"
	"github.com/dev2choiz/hello/pkg/helloworld"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExecuteGin() error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, helloworld.Say())
	})

	r.POST("/notify/function1", handlers.PubSubNotifyGin)

	f := func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	}
	r.GET("/healthz", f)
	r.GET("/ready", f)

	return r.Run(":8080")
}
