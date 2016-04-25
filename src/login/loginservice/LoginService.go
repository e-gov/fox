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
	"authn"
	"util"
)

var log = logging.MustGetLogger("login")

func main() {
	var port = flag.Int("port", 8091, "Port to bind to on the localhost interface")
	var slog = flag.Bool("syslog", false, "If present, logs are sent to syslog")

	flag.Parse()
	
	initConfig()
	setupLogging(slog)
	
	router := login.NewRouter()
	log.Infof("Starting a server on localhost:%d", *port)
	log.Critical(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}

func initConfig()  {
	util.LoadConfig()
	
	authn.InitMint()
	sc := make(chan os.Signal, 1)
	
	signal.Notify(sc, syscall.SIGHUP)
	
	go func ()  {
		for {
			<-sc
			util.LoadConfig()
		}		
	}()
}

func setupLogging(slog *bool)  {
	var b logging.Backend
	
	format := logging.MustStringFormatter(
    	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)

	if *slog  {
		b, _ = logging.NewSyslogBackend("Login")	
	}else{
		b = logging.NewLogBackend(os.Stdout, "", 0)	
	}
	
	bFormatter := logging.NewBackendFormatter(b, format)
	logging.SetBackend(bFormatter)	
}