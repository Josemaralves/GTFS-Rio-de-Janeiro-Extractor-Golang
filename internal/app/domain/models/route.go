package models

type Route struct {
	IdRoute, LongName, ShortName, AgencyId, AgencyName string
}

func ToRoutes(routesCsv [][]string, agency map[string]string) (routes []Route) {
	for _, r := range routesCsv {
		route := Route{
			IdRoute:    r[5],
			LongName:   r[6],
			ShortName:  r[1],
			AgencyId:   r[0],
			AgencyName: agency[r[0]],
		}

		routes = append(routes, route)
	}

	return routes

}
