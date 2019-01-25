package email

import (
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
	"os"
	"portfolio-be-api/app/service"
	"strings"
)

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