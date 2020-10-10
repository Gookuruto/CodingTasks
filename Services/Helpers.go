package services

import (
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"

	models "github.com/Gookuruto/CodingTasks/Models"
	osrm "github.com/karmadon/gosrm"
	geo "github.com/paulmach/go.geo"
)

//ParseStringLocation parse string in format latitude,longitude to Location structure
func ParseStringLocation(location string) models.Location {
	destinationArray := strings.Split(location, ",")
	latitude, err := strconv.ParseFloat(destinationArray[0], 64)
	if err != nil {
		panic(err)
	}
	longitude, err := strconv.ParseFloat(destinationArray[1], 64)
	if err != nil {
		panic(err)
	}
	return models.NewLocation(latitude, longitude)
}

//SetOsrmClient Setup default osrm client for this task
func SetOsrmClient() *osrm.OsrmClient {
	options := &osrm.Options{
		Url:            url.URL{Host: "https://router.project-osrm.org/"},
		Service:        osrm.ServiceRoute,
		Version:        osrm.VersionFirst,
		Profile:        osrm.ProfileDriving,
		RequestTimeout: 5,
	}
	client := osrm.NewClient(options)

	return client
}

type Result struct {
	Message models.Route
	Error   error
}

//MakeOsrmRequest make a request based on source and destination points
func MakeOsrmRequest(client *osrm.OsrmClient, source, destination models.Location) <-chan Result {
	r := make(chan Result)
	var err2 error = nil
	go func() {
		request := &osrm.RouteRequest{
			Coordinates: geo.PointSet{{source.Latitude, source.Longitude}, {destination.Latitude, destination.Longitude}},
		}
		resp, err := client.Route(request)

		if err != nil {
			log.Print("route failed: %v", err)
		}
		if resp.Code != "Ok" {
			err2 = errors.New(resp.Code)
		}
		r <- Result{models.NewRoute(resp.Routes[0].Distance, resp.Routes[0].Duration, destination), err2}
	}()
	return r

}
