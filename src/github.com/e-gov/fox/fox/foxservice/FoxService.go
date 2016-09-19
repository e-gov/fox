package main

import (
	"flag"
	"fmt"
	"github.com/e-gov/fox/fox"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/e-gov/fox/authn"
	"github.com/e-gov/fox/util"
	log "github.com/Sirupsen/logrus"
)


func main() {

	var port = flag.Int("port", 8090, "Port to bind to on the localhost interface")
	var env = flag.String("env", "DEV", "Environment the binary runs in. Accepts DEV and PROD")
	flag.Parse()

	setupLogging(env)
	initConfig()

	router := fox.NewRouter()
	log.Infof("Starting a server on localhost:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}

func initConfig() {
	util.LoadConfig()
	authn.InitValidator()

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGHUP)

	go func() {
		for {
			<-sc
			util.LoadConfig()
		}
	}()
}

//func setupLogging(slog *bool) {
//	var b logging.Backend
//
//	format := logging.MustStringFormatter(
//		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
//	)
//
//	if *slog {
//		b, _ = logging.NewSyslogBackend("Fox")
//	} else {
//		b = logging.NewLogBackend(os.Stdout, "", 0)
//	}
//
//	bFormatter := logging.NewBackendFormatter(b, format)
//	logging.SetBackend(bFormatter)
//}

func setupLogging(env *string){
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.WithFields(log.Fields{
		"env":*env,
	}).Info("Launching in " + *env + " environment")
}