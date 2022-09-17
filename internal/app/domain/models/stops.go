package models

import "strconv"

type Stop struct {
	StopId    string  `bson:"stopId"`
	Name      string  `bson:"name"`
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
}

func ToStops(raw [][]string) (stops []Stop) {
	for _, s := range raw {
		lat, _ := strconv.ParseFloat(s[2], 64)
		lon, _ := strconv.ParseFloat(s[3], 64)

		stop := Stop{
			StopId:    s[0],
			Name:      s[1],
			Latitude:  lat,
			Longitude: lon,
		}

		stops = append(stops, stop)
	}

	return
}
