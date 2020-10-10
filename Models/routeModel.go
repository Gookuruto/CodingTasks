package models

import "fmt"

//Route route struture
type Route struct {
	Distance    float64 `json:"distance"`
	Time        float64 `json:"duration"`
	Destination string
}

//NewRoute Return new instance of Route
func NewRoute(distance, time float64, destination Location) Route {
	l := new(Route)
	l.Distance = distance
	l.Time = time
	l.Destination = fmt.Sprintf("%f,%f", destination.Latitude, destination.Longitude)
	return *l
}

//ByTimeAndDistance interface for purpose of sorting arrays of Route
type ByTimeAndDistance []Route

func (a ByTimeAndDistance) Len() int { return len(a) }
func (a ByTimeAndDistance) Less(i, j int) bool {
	if a[i].Time == a[j].Time {
		return a[i].Distance < a[j].Distance
	}
	return a[i].Time < a[j].Time
}
func (a ByTimeAndDistance) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
