package main

import (
	"flag"
	"fmt"
	"github.com/op/go-logging"
	"net/http"
	"os/signal"
	"syscall"
	"os"
	"fox"
	"util"
)

var log = logging.MustGetLogger("FoxService")


func main()  {
	format := logging.MustStringFormatter(
    	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	b := logging.NewLogBackend(os.Stdout, "", 0)
	bFormatter := logging.NewBackendFormatter(b, format)
	logging.SetBackend(bFormatter)
	
	var port = flag.Int("port", 8090, "Port to bind to on the localhost interface")
	flag.Parse()

	router := fox.NewRouter()
	log.Infof("Starting a server on localhost:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}

func init()  {
	util.LoadConfig()
	
	sc := make(chan os.Signal, 1)
	
	signal.Notify(sc, syscall.SIGHUP)
	
	go func ()  {
		for {
			<-sc
			util.LoadConfig()
		}		
	}()
}