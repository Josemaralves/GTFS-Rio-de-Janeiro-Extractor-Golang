package models

func AgencyToMap(csv [][]string) map[string]string {
	agency := make(map[string]string)

	for _, raw := range csv {
		agency[raw[0]] = raw[1]
	}

	return agency
}
