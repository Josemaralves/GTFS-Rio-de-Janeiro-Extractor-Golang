package models

import "strconv"

type Itinerary struct {
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
	Sequence  int     `bson:"sequence"`
}

func ToShapes(raw [][]string) (shapes map[string][]Itinerary) {
	shapes = make(map[string][]Itinerary)

	for i := 0; i < len(raw); {
		var itineraries []Itinerary
		it := raw[i]
		itinerary := getItinerary(it)
		itineraries = append(itineraries, itinerary)

		routeId := getRouteId(it)
		for i++; i < len(raw) && getRouteId(raw[i]) == routeId; i++ {
			itineraries = append(itineraries, getItinerary(raw[i]))
		}

		shapes[routeId] = itineraries
	}
	return
}

func getRouteId(v []string) string {
	return v[0][0:10]
}

func getItinerary(v []string) Itinerary {
	sequence, _ := strconv.Atoi(v[3])
	lat, _ := strconv.ParseFloat(v[1], 64)
	lon, _ := strconv.ParseFloat(v[2], 64)

	return Itinerary{
		Latitude:  lat,
		Longitude: lon,
		Sequence:  sequence,
	}
}
