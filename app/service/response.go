package service

type ResponseEnvelope struct {
	IsSucess     bool        `json:"isSucess"`
	Message      string      `json:"message,omitempty"`
	ResponseBody interface{} `json:"responseDetails,omitempty"`
}
