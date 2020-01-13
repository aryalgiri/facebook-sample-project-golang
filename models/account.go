package models

// Account is the data model of the sample-fb service
type Account struct {
	ID          string   `json:"id,omitempty"`
	FirstName   string   `json:"firstname,omitempty"`
	MiddleName  string   `json:"middlename,omitempty"`
	LastName    string   `json:"lastname,omitempty"`
	Email       string   `json:"email,omitempty"`
	PhoneNumber string   `json:"phonenumber,omitempty"`
	OldPassword string   `json:"oldpassword,omitempty"`
	Password    string   `json:"password,omitempty"`
	Gender      string   `json:"gender,omitempty"`
	Address     *Address `json:"address,omitempty"`
	LoginToken  string   `json:"logintoken,omitempty"`
}
