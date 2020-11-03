package locations

type Country struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	TimeZone       string         `json:"time_zone"`
	GeoInformation GeoInformation `json:"geo_information"`
	States         []State        `json:"states"`
}

type GeoInformation struct {
	Location GeoLocation `json:"location"`
}

// 緯度経度
type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitute float64 `json:"longitude"`
}

type State struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
