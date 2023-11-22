package models

type Customer struct {
	ID          int       `json:"id" gorm:"primary_key"`
	LastName    string    `json:"lastname"`
	FirstName   string    `json:"firstname"`
	PhoneNumber string    `json:"phonenumber"`
	Address1    string    `json:"address1"`
	Address2    string    `json:"address2"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Zip         string    `json:"zip"`
	Vehicles    []Vehicle `gorm:"ForeignKey:VehId"`
}
