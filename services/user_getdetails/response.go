package getdetails

// Response ...
type Response struct {
	UserID   string  `json:"userID,omitempty"`
	Name     string  `json:"name,omitempty"`
	MobileNo string  `json:"mobileNo,omitempty"`
	Email    string  `json:"email,omitempty"`
	Message  string  `json:"message,omitempty"`
	Address  Address `json:"address,omitempty"`
}

// Address ...
type Address struct {
	HouseNo    string `json:"houseNo/wardNo,omitempty"`
	LandMark   string `json:"landMark,omitempty"`
	Area       string `json:"area/village,omitempty"`
	Thana      string `json:"thana,omitempty"`
	PostOffice string `json:"postOffice,omitempty"`
	Town       string `json:"town/city,omitempty"`
	State      string `json:"state,omitempty"`
}
