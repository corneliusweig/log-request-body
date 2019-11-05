package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type requestHandler struct{}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	go func() { log.Fatal(http.ListenAndServe(":"+port, &requestHandler{})) }()

	for {
		time.Sleep(60 * time.Second)
		log.Print("still listening..")
	}
}

func (h *requestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request body: ")
	io.Copy(os.Stdout, r.Body)
	defer r.Body.Close()
	fmt.Print("\n")
}
