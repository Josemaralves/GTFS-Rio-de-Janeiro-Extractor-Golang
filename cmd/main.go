package main

import (
	"github.com/josemaralves/gtfs-rio-de-janeiro-extractor-golang/internal/app/domain/models"
	repository2 "github.com/josemaralves/gtfs-rio-de-janeiro-extractor-golang/internal/app/repository"
	"github.com/josemaralves/gtfs-rio-de-janeiro-extractor-golang/internal/app/utils"
)

const path = "gtfs_rio-de-janeiro/"

var repository *repository2.RouteRepository

func init() {
	repository = repository2.New()
}

func main() {
	routesCsv := utils.ReadCsvFile(path + "routes.txt")
	agencyCsv := utils.ReadCsvFile(path + "agency.txt")

	agency := models.AgencyToMap(agencyCsv)
	routes := models.ToRoutes(routesCsv, agency)

	repository.InsertRoutes(routes)
}
