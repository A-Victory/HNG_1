package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Info struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

func index(w http.ResponseWriter, r *http.Request) {
	my_info := Info{}

	my_info.SlackUsername = "victoryotaghogho"
	my_info.Age = 22
	my_info.Backend = true
	my_info.Bio = "I'm physicist and an aspiring web developer(backend). I'm currently engaging in an internship programming initiated by HNG where I hope to further improve in working with other developers as well as fostering long last relationships. "

	json.NewEncoder(w).Encode(my_info)
}

func main() {

	http.HandleFunc("/", index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(port, nil))
}
