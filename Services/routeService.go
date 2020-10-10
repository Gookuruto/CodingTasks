package services

import (
	"errors"
	"sort"

	models "github.com/Gookuruto/CodingTasks/Models"
	"github.com/gin-gonic/gin"
)

//GetRoutes return routes to all destinations
func GetRoutes(c *gin.Context) (models.Response, error) {
	client := SetOsrmClient()
	destinations, _ := c.Request.URL.Query()["dst"]
	var allRoutes []models.Route
	source := c.Query("src")
	sourceLocation := ParseStringLocation(source)
	result := make(chan Result)
	var err error

	for _, destination := range destinations {
		location := ParseStringLocation(destination)
		go MakeOsrmRequest(client, sourceLocation, location, result)
		route := <-result
		if route.Error != nil {
			err = errors.New("some routes are not returned")
		} else {
			allRoutes = append(allRoutes, route.Message)
		}
	}

	sort.Sort(models.ByTimeAndDistance(allRoutes))
	return models.NewResponse(sourceLocation, allRoutes), err

}
