package logger

import (
	"cloud.google.com/go/logging"
	"context"
	"fmt"
	"google.golang.org/genproto/googleapis/api/monitoredres"
	"log"
	"os"
)

var sdInst *logging.Logger

var gcpResource = &monitoredres.MonitoredResource{Type: "k8s_container"}

func initOff() {
	ctx := context.Background()
	projectID := os.Getenv("GCP_PROJECT_ID")
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	sdInst = client.Logger("hello-api")
}

func CInfof(msg string, args ...interface{}) {
	cLogf(logging.Info, msg, args)
}

func CWarnf(msg string, args ...interface{}) {
	cLogf(logging.Warning, msg, args)
}

func cLogf(s logging.Severity, msg string, args ...interface{}) {
	sdInst.Log(logging.Entry{
		Severity: s,
		Payload:  fmt.Sprintf(msg, args...),
		Resource: gcpResource,
	})
}
