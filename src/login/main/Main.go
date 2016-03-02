package main

import (
	"flag"
	"fmt"
	"github.com/op/go-logging"
	"login"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"fox"
	"authn"
)

var log = logging.MustGetLogger("login")

func main() {
	format := logging.MustStringFormatter(
    	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	b := logging.NewLogBackend(os.Stdout, "", 0)
	bFormatter := logging.NewBackendFormatter(b, format)
	logging.SetBackend(bFormatter)
	
	var port = flag.Int("port", 8090, "Port to bind to on the localhost interface")
	flag.Parse()

	router := login.NewRouter()
	log.Debugf("Starting a server on localhost:%d", *port)
	log.Critical(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}

func init()  {
	fox.LoadConfig()
	authn.InitMint()
	sc := make(chan os.Signal, 1)
	
	signal.Notify(sc, syscall.SIGHUP)
	
	go func ()  {
		for {
			<-sc
			fox.LoadConfig()
		}		
	}()
}