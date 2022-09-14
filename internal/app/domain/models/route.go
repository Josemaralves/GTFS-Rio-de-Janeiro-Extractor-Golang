package models

type Route struct {
	Id, LongName, ShortName, AgencyId, AgencyName string
}

func ToRoutes(routesCsv [][]string, agency map[string]string) (routes []Route) {
	for _, r := range routesCsv {
		route := Route{
			Id:         r[5],
			LongName:   r[6],
			ShortName:  r[1],
			AgencyId:   r[0],
			AgencyName: agency[r[0]],
		}

		routes = append(routes, route)
	}

	return routes

}
