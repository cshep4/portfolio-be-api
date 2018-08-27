package sms

import (
	"net/http"
	"fmt"
)

type SmsService struct{}

//[{"Number": "381641234567", "Content": "Sms!!!"}]
type SmsArgs struct {
	Number, Content string
}

type Response struct {
	Result string
}

func (t *SmsService) SendSMS(r *http.Request, args *SmsArgs, result *Response) error {
	*result = Response{Result: fmt.Sprintf("Sms sent to %s", args.Number)}
	return nil
}

func (t *SmsService) Path() string {
	return "/sms"
}

func (t *SmsService) Name() string {
	return "sms"
}