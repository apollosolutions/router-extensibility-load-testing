package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type GQLResponse struct {
	Errors []GQLError
	Data   interface{}
}

type GQLError struct {
	Message    string
	Extensions interface{}
}

var endpoint = "http://router:4040"
var headers = map[string][]string{
	"Authentication": {"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjEyMzEyMzEyMzEyMywiY2xpZW50X25hbWUiOiJhcG9sbG8gY29wcm9jZXNzb3IifQ.HucZYEikOmBebtLximNoYLFPbNbBeQA_gRWcSs-dYEI"},
	"Content-Type":   {"application/json"},
}

var targets = []vegeta.Target{
	{
		Method: "POST",
		URL:    endpoint,
		Body: []byte(`{
  "query": "query getAllLocations {\n  locations {\n    description\n    id\n    overallRating\n    reviewsForLocation {\n      comment\n    }\n  }\n}",
  "variables": {},
  "operationName": "getAllLocations"
}`),
		Header: headers,
	}, {
		Method: "POST",
		URL:    endpoint,
		Body: []byte(`{
  "query": "query getLocationById($locationId: ID!) {\n  location(id: $locationId) {\n    description\n    id\n    name\n    overallRating\n    photo\n  }\n}",
  "variables": {
    "locationId": "loc-1"
  },
  "operationName": "getLocationById"
}`),
		Header: headers,
	},
}

func main() {
	rateFlag := flag.Int("rate", 100, "Rate of requests in rps")
	durationFlag := flag.String("duration", "5s", "Duration of request")
	pathFlag := flag.String("out", "", "Path to JSON output file")
	flag.Parse()

	if *rateFlag == 0 {
		log.Println("Rate must be greater than 0.")
		os.Exit(1)
	}
	if *durationFlag == "" {
		log.Println("Invalid duration format.")
		os.Exit(1)
	}
	if *pathFlag == "" {
		log.Println("Output must be set in order to run.")
		os.Exit(1)
	}
	duration, err := time.ParseDuration(*durationFlag)
	if err != nil {
		log.Println("Invalid duration format.")
		os.Exit(1)
	}
	rate := vegeta.Rate{Freq: *rateFlag, Per: time.Second}
	target := vegeta.NewStaticTargeter(targets...)

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(target, rate, duration, "GQL Test") {
		if res.Code != 200 {
			metrics.Add(res)
			continue
		}

		var response GQLResponse
		err := json.Unmarshal(res.Body, &response)
		if err != nil {
			res.Error = err.Error()
			metrics.Add(res)
			continue
		}

		if response.Errors != nil && len(response.Errors) > 0 {
			var error string
			for i, e := range response.Errors {
				if i == 0 {
					error = e.Message
					continue
				}
				error = fmt.Sprintf("%v, %v", error, e.Message)
			}
			res.Code = 418
			res.Error = error
		}
		metrics.Add(res)
	}

	metrics.Close()
	report := vegeta.NewJSONReporter(&metrics)
	report.Report(os.Stdout)
	dir := filepath.Dir(*pathFlag)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	file, err := os.OpenFile(*pathFlag, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	err = report.Report(file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
