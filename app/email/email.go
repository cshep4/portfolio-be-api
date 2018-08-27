package email

import (
	"net/http"
	"fmt"
)

type EmailService struct{}

//[{"To": "demo@example.com","Subject": "Hello", "Content": "Hi!!!"}]
type EmailArgs struct {
	To, Subject, Content string
}

func (t *EmailService) SendEmail(r *http.Request, args *EmailArgs, result *string) error {
	*result = fmt.Sprintf("Email sent to %s", args.To)
	return nil
}

func (t *EmailService) Path() string {
	return "/email"
}

func (t *EmailService) Name() string {
	return "email"
}