package register

type RequestBody struct {
	MobileNo string `json:"mobileNo"`
	Password string `json:"password"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
}

type Request struct {
	RequestBody RequestBody `json:"requestBody,omitempty"`
}

func (req Request) validateRequest() (string, bool) {
	if req.RequestBody.MobileNo == "" || req.RequestBody.Password == "" {
		return "invalid request", false
	}
	return "", true
}
