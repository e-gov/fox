package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"os"
)

func main()  {
	var nodeName string 
	var port = flag.Int("port", 8090, "Port to bind to on the localhost interface")
	flag.StringVar(&nodeName,"name", "my", "Name of the running instance")
	flag.Parse()

	router := NewRouter(nodeName)
	log.Printf("Starting a server on localhost:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}

func init()  {
	LoadConfig()
	
	sc := make(chan os.Signal, 1)
	
	signal.Notify(sc, syscall.SIGHUP)
	
	go func ()  {
		for {
			<-sc
			fox.LoadConfig()
		}		
	}()
}