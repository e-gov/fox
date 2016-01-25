package main

import (
	"log"
	"net/http"
	"os"
)

func main()  {
	nodeName = os.Args[1]

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8090", router))
}