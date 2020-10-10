package models

//Location representation of geographic location
type Location struct {
	Latitude  float64
	Longitude float64
}

//NewLocation return new pointer to Location
func NewLocation(latitude, longitude float64) Location {
	l := new(Location)
	l.Latitude = latitude
	l.Longitude = longitude
	return *l
}
