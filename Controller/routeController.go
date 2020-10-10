package controller

import (
	"net/http"

	services "github.com/Gookuruto/CodingTasks/Services"
	"github.com/gin-gonic/gin"
)

//RouteController struct for more object like structure of controller
type RouteController struct{}

//FindRoutes returns source and destinations with information about time to get to them and how far they are from source
func (r RouteController) FindRoutes(c *gin.Context) {
	routes, err := services.GetRoutes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, routes)
	}

}
