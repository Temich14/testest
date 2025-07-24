package messagies

type Create struct {
	Type     string `json:"type"`
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
