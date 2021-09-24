package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
)

const rootCertPath = "./../local/files/cert/client.pem"
const protoPath = "./../api/proto"
const thirdPartyPath = "./../third_party"

const url = "api.dev2choiz.com:443"

func main() {
	report, err := runner.Run(
		"hello.endpoints.health.Health.Status",
		url,
		runner.WithProtoFile("health.proto", []string{thirdPartyPath, protoPath}),
		runner.WithDataFromFile("data.json"),
		runner.WithRootCertificate(rootCertPath),
		// test config
		//runner.WithConcurrency(100),
		runner.WithConcurrencySchedule("step"),
		runner.WithConcurrencyStart(15),
		runner.WithConcurrencyEnd(50),
		runner.WithConcurrencyStep(5),
		runner.WithConcurrencyStepDuration(5 * time.Second),

		runner.WithRunDuration(time.Duration(120) * time.Second),
		runner.WithStreamInterval(time.Duration(250) * time.Millisecond),
		runner.WithStreamCallCount(4),
		//runner.WithStreamCallDuration(time.Duration(1) * time.Second),
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	const logPath = "./report.html"
	_ = os.Remove(logPath)
	f, err := os.Create(logPath)
	printer := printer.ReportPrinter{Report: report }

	printer.Out = os.Stdout
	printer.Print("summary")

	printer.Out = f
	printer.Print("html")
}
