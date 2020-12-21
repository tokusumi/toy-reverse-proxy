package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Printf("start vanila in PORT:1385")
	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(res http.ResponseWriter, req *http.Request) {
		message := map[string]string{
			"message": "Hello world, vanila Golang API server",
		}
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			panic(err.Error())
		}
		res.Write(jsonMessage)
	})
	log.Fatal(http.ListenAndServe(":1385", mux))
}
