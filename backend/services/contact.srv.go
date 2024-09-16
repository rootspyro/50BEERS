package services

import (
	"fmt"

	"github.com/rootspyro/50BEERS/SDKs/mailtrap"
)

type ContactSrv struct{
	email string
	sdk   *mailtrap.MailtrapSDK
}

func NewContactSrv(contactEmail string, sdk *mailtrap.MailtrapSDK) *ContactSrv {
	return &ContactSrv{
		email: contactEmail,
		sdk: sdk,
	}
}

func (srv *ContactSrv) SendContactEmail(data ContactDTO) error {

	var email mailtrap.Email

	email.From.Name = data.Name
	email.From.Email = data.Email
	email.To = append(email.To, struct{Email string "json:\"email\""}{
		Email: srv.email,
	})
	email.Subject = fmt.Sprintf("Contact from blog - %s", data.Name)
	email.Text = data.Message

	return srv.sdk.SendEmail(email)
}

type ContactDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
