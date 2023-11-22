package models

type Vehicle struct {
	VID      int    `json:"vid" gorm:"primary_key"`
	Make     string `json:"make"`
	CarModel string `json:"carmodel"`
	VIN      string `json:"vin"`
	Year     int    `json:"year"`
	VehId    int    `json:"vehid"`
	Tires    []Tire `gorm:"ForeignKey:TireID"`
}
