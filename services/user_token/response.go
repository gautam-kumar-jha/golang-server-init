package token

type Response struct {
	Token   string `json:"token,omitempty"`
	UserID  string `json:"userID,omitempty"`
	Message string `json:"message,omitempty"`
}
