package models

import "fmt"

//Response response structure
type Response struct {
	Source string  `json:"source"`
	Routes []Route `json:"routes"`
}

//NewResponse return nw instance of Response
func NewResponse(source Location, routes []Route) Response {
	l := new(Response)
	l.Source = fmt.Sprintf("%f,%f", source.Latitude, source.Longitude)
	l.Routes = routes
	return *l
}
