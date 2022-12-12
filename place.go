package gotwtr

type PlaceField string

/*
	Place field will only return
	if you've also included the expansions=geo.place_id query parameter in your request.
*/

const (
	PlaceFieldContainedWithin PlaceField = "contained_within"
	PlaceFieldCountry         PlaceField = "country"
	PlaceFieldCountryCode     PlaceField = "country_code"
	PlaceFieldFullName        PlaceField = "full_name"
	PlaceFieldGeo             PlaceField = "geo"
	PlaceFieldID              PlaceField = "id"
	PlaceFieldName            PlaceField = "name"
	PlaceFieldPlaceType       PlaceField = "place_type"
)

type Place struct {
	FullName        string    `json:"full_name"`
	ID              string    `json:"id"`
	ContainedWithin []string  `json:"contained_within,omitempty"`
	Country         string    `json:"country,omitempty"`
	CountryCode     string    `json:"country_code,omitempty"`
	Geo             *PlaceGeo `json:"geo,omitempty"`
	Name            string    `json:"name,omitempty"`
	PlaceType       string    `json:"place_type,omitempty"`
}

type PlaceGeo struct {
	Type       string                 `json:"type"`
	BBox       []float64              `json:"bbox"`
	Properties map[string]interface{} `json:"properties"`
}

func placeFieldsToString(pfs []PlaceField) []string {
	slice := make([]string, len(pfs))
	for i, pf := range pfs {
		slice[i] = string(pf)
	}
	return slice
}
