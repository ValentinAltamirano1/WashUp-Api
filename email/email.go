package email

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailClient interface {
	ResetPassword(email, uniqueID string) error
}
type SendGridClient struct {
	APIKey string
}

func NewSendGridClient(apiKey string) EmailClient {
	return &SendGridClient{APIKey: apiKey}
}

func (s *SendGridClient) ResetPassword(email, uniqueID string) error {
	from := mail.NewEmail("WashUp", "valealta28@gmail.com")
	subject := "Restablecer contraseña"
	to := mail.NewEmail("", email)

	link := "http://localhost:3000/reset-password/" + uniqueID
	content := "Haz clic en el siguiente enlace para restablecer tu contraseña: " + link
	plainTextContent := content
	htmlContent := "<p>" + content + "</p>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(s.APIKey)

	_, err := client.Send(message)
	fmt.Println(message)
	fmt.Println(err)
	return err	
}
