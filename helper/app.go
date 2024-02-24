package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
}

// StartWebserver start simple HTTP server
func StartWebserver() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":3232", nil))
}
