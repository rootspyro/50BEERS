package services

type ContactSrv struct{}

func NewContactSrv() *ContactSrv {
	return &ContactSrv{}
}

func (srv *ContactSrv) SendEmail() error {
	return nil
}

type EmailDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
