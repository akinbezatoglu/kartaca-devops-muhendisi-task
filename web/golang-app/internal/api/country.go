package api

import (
	"net/http"

	"github.com/akinbezatoglu/kartaca-devops-muhendisi-task/web/golang-app/internal/api/response"
)

// Handles to get one random country from the elasticsearch cluster.
func (api *API) countryGetHandler(w http.ResponseWriter, r *http.Request) {

	country, err := api.ES.GetOneRandomCountry()
	if err != nil {
		response.Errorf(w, r, nil, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.Write(w, r, &country)
}
