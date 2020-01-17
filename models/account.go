package models

// Account is the data model of the sample-fb service
type Account struct {
	ID          string   `json:"id,omitempty" db:"id"`
	Firstname   string   `json:"firstname,omitempty" db:"firstname"`
	MiddleName  string   `json:"middlename,omitempty" db:"middlename"`
	Lastname    string   `json:"lastname,omitempty" db:"lastname"`
	Email       string   `json:"email,omitempty" db:"email"`
	PhoneNumber string   `json:"phonenumber,omitempty" db:"phonenumber"`
	Password    string   `json:"password,omitempty" db:"password"`
	Gender      string   `json:"gender,omitempty" db:"gender"`
	Address     *Address `json:"address,omitempty" db:"address"`
	Status      string   `json:"status,omitempty" db:"status"`
	OldPassword string   `json:"oldpassword,omitempty" db:"oldpassword"`
}
