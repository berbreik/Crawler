// handler is the implementation of the service interface 
package handler

import (
	"encoding/json"
	"net/http"

	"assignment_v0.2/models"
)

// crawlerHandler implements the service interface
type crawlerHandler struct {
	crawlerService serviceInterfae
}

// NewHandler returns a new instance of the handler
func NewHandler(initialized serviceInterfae) *crawlerHandler {
	return &crawlerHandler{
		crawlerService: initialized,
	}
}

// Crawl is the handler for the /crawl endpoint, it encodes the response in json and decodes the request 
func (s *crawlerHandler) Crawl(w http.ResponseWriter, r *http.Request) {
	var requestData models.RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData, err := s.crawlerService.Crawl(requestData.Urls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
