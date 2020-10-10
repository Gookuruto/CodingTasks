package tests

import (
	"sort"
	"testing"

	models "github.com/Gookuruto/CodingTasks/Models"
)

//TestSort Testing sort on slice of Route strutures
func TestSort(t *testing.T) {
	testArray := []models.Route{
		{100, 230, ""},
		{200, 23, ""},
		{100, 23, ""},
	}
	expectedResult := []models.Route{
		{100, 23, ""},
		{200, 23, ""},
		{100, 230, ""},
	}
	sort.Sort(models.ByTimeAndDistance(testArray))
	for index, route := range testArray {
		if route.Time != expectedResult[index].Time && route.Distance != expectedResult[index].Distance {
			t.Errorf("route is not sorted as expeted")
		}
	}
}
