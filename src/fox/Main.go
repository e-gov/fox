package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main()  {
	var port = flag.Int("port", 8090, "Port to bind to on the localhost interface")
	flag.StringVar(&nodeName,"name", "my", "Name of the running instance")
	flag.Parse()

	router := NewRouter()
	log.Printf("Starting a server on localhost:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
