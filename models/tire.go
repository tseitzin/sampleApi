package models

type Tire struct {
	TID             int    `json:"tid" gorm:"primary_key"`
	TireLocation    string `json:"tirelocation"`
	Brand           string `json:"make"`
	TireModel       string `json:"tiremodel"`
	TireSize        string `json:"tiresize"`
	MaxPressure     int    `json:"maxpressure"`
	CurrentPressure int    `json:"currentpressure"`
	TreadWear       string `json:"treadwear"`
	TireID          int    `json:"tireid"`
}
