package util

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"os/signal"
	"syscall"
)

func SetupSvcLogging(env *string){
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.WithFields(log.Fields{
		"env":*env,
	}).Info("Launching in " + *env + " environment")
}

func InitConfig() {
	LoadConfig()

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGHUP)

	go func() {
		for {
			<-sc
			LoadConfig()
		}
	}()
}

