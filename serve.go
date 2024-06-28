package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		fmt.Println("./serve <directory>")
		log.Fatal("Please provide a directory to serve")
	}
	fs := http.FileServer(http.Dir(argsWithoutProg[0]))
	http.Handle("/", fs)

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
