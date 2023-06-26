package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CoprocessorBody struct {
	Query         string        `json:"query,omitempty"`
	OperationName string        `json:"operationName,omitempty"`
	Variables     interface{}   `json:"variables,omitempty"`
	Data          interface{}   `json:"data,omitempty"`
	Errors        []interface{} `json:"errors,omitempty"`
}

type CoprocessorContext struct {
	Entries interface{}
}
type CoprocessorJSON struct {
	Version int                 `json:"version"`
	Stage   string              `json:"stage"`
	Control interface{}         `json:"control"`
	ID      string              `json:"id"`
	Headers map[string][]string `json:"headers,omitempty"`
	Body    *CoprocessorBody    `json:"body,omitempty"`
	Context *CoprocessorContext `json:"context,omitempty"`
	SDL     string              `json:"sdl,omitempty"`
	Method  string              `json:"method,omitempty"`

	// subgraph* stage specific
	URI         string `json:"uri,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`

	// response stage specific
	StatusCode int `json:"status_code,omitempty"`

	// router stage specific
	Path string `json:"path,omitempty"`
}

// For JWT Client Awareness example
type JWTClaims struct {
	ClientName    string `json:"client_name"`
	ClientVersion string `json:"client_version"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("apollo")

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
	}

	http.HandleFunc("/static-subgraph", static_subgraph)
	http.HandleFunc("/guid-response", guid_response)
	http.HandleFunc("/client-awareness", client_awareness)
	log.Printf("Starting on :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func static_subgraph(w http.ResponseWriter, r *http.Request) {
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

	if payload.Headers == nil {
		payload.Headers = make(map[string][]string)
	}
	payload.Headers["source"] = []string{"coprocessor"}

	responseBody, err := json.Marshal(&payload)
	if err != nil {
		panic(err)
	}
	w.Write(responseBody)
}

func guid_response(w http.ResponseWriter, r *http.Request) {
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
	if payload.Headers == nil {
		payload.Headers = make(map[string][]string)
	}
	for i := 0; i < 10; i++ {
		payload.Headers["GUID"] = append(payload.Headers["GUID"], uuid.New().String())
	}

	responseBody, err := json.Marshal(&payload)
	if err != nil {
		panic(err)
	}
	w.Write(responseBody)
}

func client_awareness(w http.ResponseWriter, r *http.Request) {
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

	if payload.Headers["authentication"] == nil {
		responseBody := unauthorized(payload)
		w.Write(responseBody)
		return
	}
	rawToken := strings.Split(payload.Headers["authentication"][0], "Bearer ")[1]

	if rawToken == "" {
		responseBody := unauthorized(payload)
		w.Write(responseBody)
		return
	}
	claims := &JWTClaims{
		ClientName:    "coprocessor",
		ClientVersion: "loadtest",
	}

	jwtToken, err := jwt.ParseWithClaims(rawToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !jwtToken.Valid {
		responseBody := unauthorized(payload)
		w.Write(responseBody)
		return
	}

	payload.Headers["apollographql-client-name"] = []string{claims.ClientName}
	payload.Headers["apollographql-client-version"] = []string{claims.ClientVersion}
	responseBody, err := json.Marshal(&payload)
	if err != nil {
		panic(err)
	}
	w.Write(responseBody)
}

func unauthorized(payload *CoprocessorJSON) []byte {
	payload.Control = map[string]interface{}{
		"Break": 401,
	}
	responseBody, err := json.Marshal(&payload)
	if err != nil {
		panic(err)
	}
	return responseBody
}
