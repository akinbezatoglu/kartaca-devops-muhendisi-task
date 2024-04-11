package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/akinbezatoglu/kartaca-devops-muhendisi-task/web/golang-app/internal/server"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func main() {
	// Create server instance
	instance := server.NewInstance()

	// Interrupt handler
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		log.Infof("Received %s signal", <-c)
		instance.Shutdown()
	}()

	// Start server
	instance.Start()
}
