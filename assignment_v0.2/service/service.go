// package service implements the buisness logic and is used by the handler
package service

import (
	"errors"
	"io/ioutil"
	"net/http"
	"sync"

	"assignment_v0.2/models"
)

type ServiceCrawl struct{}

// NewService returns a new instance of the service
func NewService() *ServiceCrawl {
	return &ServiceCrawl{}
}

// Crawl iterates over the urls and fetches the response from the url
func (h *ServiceCrawl) Crawl(urls []string) (*models.ResponseData, error) {
	var wg sync.WaitGroup
	var result []models.Body

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			body, err := getResp(url)
			if err != nil {
				return
			}
			result = append(result, body)
		}(url)
	}

	wg.Wait()
	return &models.ResponseData{
		Result: result,
		Error:  nil,
	}, nil

}

// getResp fetches response from the url returns the body of the response
func getResp(url string) (models.Body, error) {

	resp, err := http.Get(url)
	if err != nil {
		return models.Body{}, errors.New("error in fetching data")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Body{}, errors.New("error in reading the body")
	}

	str := string(body)

	return models.Body{
		Url:  url,
		Data: str,
	}, nil
}
