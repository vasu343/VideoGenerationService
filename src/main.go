package main

import (
	"net/http"
	"videogeneration"
	"github.com/gorilla/mux"
)

var dao config.DatabaseDao

func init() {
	dao.Connect()
}


func authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// decorate http handler with authentication to secure the backend apis
		// Here we can write some custom function to validate the token or api-key
		// After that we can serve the request
		h.ServeHTTP(w,r)
	})
}

func main() {

	// using gorilla mux 
	r := mux.NewRouter()
	var service videogeneration.VideoGenerateService
	serviceDao := videogeneration.Dao{dao}
	service = videogeneration.VideoGenerateService{Dao: serviceDao}

	generateVideoHandler := httptransport.NewServer(
		videogeneration.GenerateVideoEndpoint(service),
		videogeneration.DecodeService,
		videogeneration.EncodeResponse)

	r.Methods("POST").Path("/api/GenerateVideo").Handler(authenticate(generateVideoHandler))