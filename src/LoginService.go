package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"login"
	"os/signal"
	"syscall"
)

func main()  {
	var port = flag.Int("port", 8090, "Port to bind to on the localhost interface")
	flag.Parse()

	router := login.NewRouter()
	log.Printf("Starting a server on localhost:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}

func init()  {
	login.LoadKey("key.fernet")
	
	sc := make(chan os.Signal, 1)
	
	signal.Notify(sc, syscall.SIGHUP)
	
	go func ()  {
		for {
			<-sc
			login.LoadKey("key.fernet")
		}		
	}()
}