package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"portfolio-be-api/app/sms"
	"portfolio-be-api/app/service"
	"portfolio-be-api/app/email"
	"os"
)

func main() {
	router := mux.NewRouter()

	service.Route(router, new(sms.SmsService))
	service.Route(router, new(email.EmailService))

	http.ListenAndServe(":" + os.Getenv("PORT"), router)
}