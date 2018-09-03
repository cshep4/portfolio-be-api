package email

import (
	"net/http"
	"os"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"errors"
	"portfolio-be-api/app/service"
	"strings"
)

type EmailService struct{}

//[{"To": "demo@example.com","From": "demo@example.com", "Sender": "Hello", "Recipient": "Hello", "Subject": "Hello", "Content": "Hi!!!"}]
type EmailArgs struct {
	To, From, Sender, Recipient, Subject, Content string
}

func (t *EmailService) SendEmail(r *http.Request, args *EmailArgs, result *service.Response) error {
	from := mail.NewEmail(args.Sender, args.From)
	subject := args.Subject

	to := mail.NewEmail(args.Recipient, args.To)

	plainTextContent := args.Content

	htmlContent := strings.Replace(args.Content,"\n","<br>",-1)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return errors.New(response.Body)
	}

	*result = service.Response{StatusCode:response.StatusCode, Body:response.Body}

	return nil
}

func (t *EmailService) Path() string {
	return "/email"
}

func (t *EmailService) Name() string {
	return "email"
}