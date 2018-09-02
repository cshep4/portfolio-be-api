package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"portfolio-be-api/app/service"
	"portfolio-be-api/app/email"
	"os"
	"portfolio-be-api/app/predictor"
)

func main() {
	router := mux.NewRouter()

	service.Route(router, new(email.EmailService))
	service.Route(router, new(predictor.PredictorService))

	http.ListenAndServe(":" + os.Getenv("PORT"), router)
}