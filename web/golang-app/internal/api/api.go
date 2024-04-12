package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akinbezatoglu/kartaca-devops-muhendisi-task/web/golang-app/internal/elasticsearch"
)

// Config represents the API configuration
type Config struct {
	Domain   string `mapstructure:"domain"`
	HttpPort int    `mapstructure:"http_port"`
}

// API represents the structure of the API
type API struct {
	Router *mux.Router
	ES     *elasticsearch.ES
}

// New returns the api settings
func New(router *mux.Router, es *elasticsearch.ES) *API {
	api := &API{
		Router: router,
		ES:     es,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	api.Router.HandleFunc("/", api.corsMiddleware(api.logMiddleware(api.helloHandler))).Methods("GET")
	api.Router.HandleFunc("/staj", api.corsMiddleware(api.logMiddleware(api.countryGetHandler))).Methods("GET")

	return api
}

func (a *API) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Merhaba, Go!")
}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
