package field

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
