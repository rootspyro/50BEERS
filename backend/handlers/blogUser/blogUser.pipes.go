package bloguser

type LoginDTO struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
