package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/akinbezatoglu/kartaca-devops-muhendisi-task/web/golang-app/internal/api"
	"github.com/akinbezatoglu/kartaca-devops-muhendisi-task/web/golang-app/internal/elasticsearch"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

// Config represents the server configuration
type Config struct {
	API *api.Config           `mapstructure:"api"`
	ES  *elasticsearch.Config `mapstructure:"elasticsearch"`
}

// Instance represents an instance of the server
type Instance struct {
	API    *api.API
	server *http.Server
}

// NewInstance returns an new instance of our server
func NewInstance() *Instance {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Fatal("Could not load configuration")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.WithError(err).Fatal("Could not load configuration")
	}

	// Initialize elasticsearch connection
	es, err := elasticsearch.New(config.ES)
	if err != nil {
		log.WithError(err).Fatal("Could not create elasticsearch instance")
	}

	// Initialize API
	router := mux.NewRouter()
	api := api.New(router, es)

	// Startup the HTTP Server in a way that we can gracefully shut it down again
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.API.Domain, config.API.HttpPort),
		Handler: api.Router,
	}

	return &Instance{
		API:    api,
		server: server,
	}
}

// Start starts the server
func (i *Instance) Start() {
	if err := i.server.ListenAndServe(); err != nil {
		log.WithError(err).Error("HTTP Server cannot started")
		i.Shutdown()
	}
}

// Shutdown stops the server
func (i *Instance) Shutdown() {

	// Shutdown HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := i.server.Shutdown(ctx); err != nil {
		log.WithError(err).Error("Failed to shutdown HTTP server gracefully")
		os.Exit(1)
	}

	log.Info("Shutdown HTTP server gracefully...")
	os.Exit(0)
}
