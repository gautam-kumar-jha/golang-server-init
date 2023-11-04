package login

type Response struct {
	ShowPage  string `json:"showPage,omitempty"`
	AuthToken string `json:"authToken,omitempty"`
	UserID    string `json:"userid,omitempty"`
	Message   string `json:"message,omitempty"`
}
