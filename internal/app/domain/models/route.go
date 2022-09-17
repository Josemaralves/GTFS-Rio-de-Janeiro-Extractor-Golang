package models

type Lines struct {
	IdRoute     string      `bson:"idRoute"`
	LongName    string      `bson:"longName"`
	ShortName   string      `bson:"shortName"`
	AgencyId    string      `bson:"agencyId"`
	AgencyName  string      `bson:"agencyName"`
	Itineraries []Itinerary `bson:"itineraries"`
	Stops       []string    `bson:"stops"`
}

func ToLines(routesCsv [][]string, agency map[string]string, shapes map[string][]Itinerary, times map[string][]string) (routes []Lines) {
	for _, r := range routesCsv {
		route := Lines{
			IdRoute:     r[5],
			LongName:    r[6],
			ShortName:   r[1],
			AgencyId:    r[0],
			AgencyName:  agency[r[0]],
			Itineraries: shapes[r[5]],
			Stops:       times[r[5]],
		}

		routes = append(routes, route)
	}

	return routes

}
