package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"portfolio-be-api/app/service"
	"portfolio-be-api/app/email"
	"os"
	"portfolio-be-api/app/predictor"
)

func main() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"**"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	router := mux.NewRouter()

	service.Route(router, new(email.EmailService))
	service.Route(router, new(predictor.PredictorService))

	http.ListenAndServe(":" + os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk) (router))
}