package main

import (
	"flag"
	"fmt"
	"github.com/e-gov/fox/fox"
	"net/http"

	"github.com/e-gov/fox/util"
	log "github.com/Sirupsen/logrus"
	"github.com/e-gov/fox/authn"
)


func main() {

	var port = flag.Int("port", 8090, "Port to bind to on the localhost interface")
	var env = flag.String("env", "DEV", "Environment the binary runs in. Accepts DEV and PROD")
	flag.Parse()

	util.SetupSvcLogging(env)
	util.InitConfig()
	authn.InitValidator()

	router := fox.NewRouter()
	log.Infof("Starting a server on localhost:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}

