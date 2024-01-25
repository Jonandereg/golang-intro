package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type RequestBody struct {
	Name string `json:"name"`
}
type ResponseBody struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	name := reqBody.Name

	if name == "" {
		name = "World"
	}
	region := os.Getenv("FLY_REGION")
	if region == "" {
		region = "unknown"
	}
	var respBody ResponseBody
	respBody.Message = fmt.Sprintf("Hello %s I am responding from fly edge computing on region %s!", name, region)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)

}

func main() {
	http.HandleFunc("/hello", helloHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
