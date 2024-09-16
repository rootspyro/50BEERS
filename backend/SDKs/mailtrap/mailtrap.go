package mailtrap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MailtrapSDK struct {
	Host        string
	APIToken    string
	DomainEmail string
}

func New(host, token, email string) *MailtrapSDK {
	return &MailtrapSDK{
		Host:        host,
		APIToken:    token,
		DomainEmail: email,
	}
}

func (m *MailtrapSDK) SendEmail(data Email) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	payload := bytes.NewBuffer([]byte(body))

	request, err := http.NewRequest(http.MethodPost, m.Host, payload)
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.APIToken))
	request.Header.Add("Content-Type", "application/json")

	client := http.Client{}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

type Email struct {
	From struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"from"`
	To []struct {
		Email string `json:"email"`
	} `json:"to"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}
