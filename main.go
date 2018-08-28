package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"go-tester/app/sms"
	"go-tester/app/service"
	"go-tester/app/email"
	"os"
)

func main() {
	router := mux.NewRouter()

	service.Route(router, new(sms.SmsService))
	service.Route(router, new(email.EmailService))

	http.ListenAndServe(":" + os.Getenv("PORT"), router)
}