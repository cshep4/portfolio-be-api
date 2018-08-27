package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"tester/app/sms"
	"tester/app/service"
	"tester/app/email"
)

func main() {
	router := mux.NewRouter()

	service.Route(router, new(sms.SmsService))
	service.Route(router, new(email.EmailService))

	http.ListenAndServe(":1337", router)
}