package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	cdx "github.com/CycloneDX/cyclonedx-go"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received a request")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "Mae"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)
}

func main() {
	log.Print("helloworld: starting server...")
	bom := cdx.NewBOM()
	log.Print(bom)
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
