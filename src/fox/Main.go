package main

import (
    "fmt"
    "log"
	"net/http"
	"os"
)

func main()  {
	nodeName = os.Args[1]

	router := NewRouter()
	var port int = 8090
	log.Printf("Starting a server on localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
