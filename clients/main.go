package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	ADDR := os.Getenv("ADDR")
	if len(ADDR) == 0 {
		ADDR = ":443"
	}

	TLSCERT := os.Getenv("TLSCERT")
	if len(TLSCERT) == 0 {
		log.Fatal("No TLSCERT environment variable found")
	}

	TLSKEY := os.Getenv("TLSKEY")
	if len(TLSKEY) == 0 {
		log.Fatal("No TLSKEY environment variable found")
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./dist")))

	log.Printf("server is listening at %s", ADDR)
	log.Fatal(http.ListenAndServeTLS(ADDR, TLSCERT, TLSKEY, mux))
}
