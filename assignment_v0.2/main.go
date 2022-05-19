// package main initializes the server and starts listening to the port 8080
package main


import (
	"net/http"

	"assignment_v0.2/handler"
	"assignment_v0.2/service"
	"github.com/gorilla/mux"
)

// main is the entry point of the application
func main() {

	handlermain := handler.NewHandler(service.NewService())

	route := mux.NewRouter()
	route.HandleFunc("/api/crawl", handlermain.Crawl).Methods("POST")

	err := http.ListenAndServe(":8080", route)
	if err != nil {
		panic(err)
	}
}
