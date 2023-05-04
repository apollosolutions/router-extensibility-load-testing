package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type CoprocessorBody struct {
	Query string
}
type CoprocessorJSON struct {
	Version int    `json:"version"`
	Stage   string `json:"stage"`
	Control string `json:"control"`
	ID      string `json:"id"`
	Headers map[string][]string
	Body    CoprocessorBody
	Context map[string]map[string]string
	SDL     string
	Path    string
	Method  string
}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8000
	}

	http.HandleFunc("/", coprocessor)

	log.Printf("Starting on :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func coprocessor(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var payload *CoprocessorJSON

	err = json.Unmarshal(body, &payload)
	if err != nil {
		panic(err)
	}

	log.Printf("%#v\n", payload)

	responseBody, err := json.Marshal(&payload)
	if err != nil {
		panic(err)
	}
	// do stuff
	w.Write(responseBody)
}
