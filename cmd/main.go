package main

import (
	"github.com/josemaralves/gtfs-rio-de-janeiro-extractor-golang/internal/app/domain/models"
	"github.com/josemaralves/gtfs-rio-de-janeiro-extractor-golang/internal/app/lib/mongodb"
	"github.com/josemaralves/gtfs-rio-de-janeiro-extractor-golang/internal/app/repository"
	"github.com/josemaralves/gtfs-rio-de-janeiro-extractor-golang/internal/app/utils"
)

const path = "gtfs_rio-de-janeiro/"

var routeRepository *repository.RouteRepository
var stopRepository *repository.StopsRepository

func init() {
	mongo := mongodb.New()
	routeRepository = repository.NewRouteRepository(mongo)
	stopRepository = repository.NewStopRepository(mongo)
}

func main() {
	processLines()
	processStops()
}

func processLines() {
	routesCsv := utils.ReadCsvFile(path + "routes.txt")
	agencyCsv := utils.ReadCsvFile(path + "agency.txt")
	shapesCsv := utils.ReadCsvFile(path + "shapes.txt")
	timesCsv := utils.ReadCsvFile(path + "stop_times.txt")

	agency := models.AgencyToMap(agencyCsv)
	shapes := models.ToShapes(shapesCsv)
	times := toTimes(timesCsv)
	routes := models.ToLines(routesCsv, agency, shapes, times)

	routeRepository.InsertRoutes(routes)
}

func toTimes(raw [][]string) (times map[string][]string) {
	times = make(map[string][]string)
	for i := 0; i < len(raw); {
		var value []string
		time := raw[i]
		stopId := time[3]
		value = append(value, stopId)

		routeId := time[0][0:10]
		for i++; i < len(raw) && raw[i][0][0:10] == routeId; i++ {
			time := raw[i]
			stopId := time[3]
			value = append(value, stopId)
		}

		times[routeId] = value
	}

	return
}

func processStops() {
	stopsCsv := utils.ReadCsvFile(path + "stops.txt")
	stops := models.ToStops(stopsCsv)

	stopRepository.InsertStops(stops)
}
