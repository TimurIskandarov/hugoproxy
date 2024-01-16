package geo

type RequestAddress struct {
	Query string `json:"query"`
}

type Address struct {
	Lat string `json:"geo_lat"`
	Lng string `json:"geo_lon"`
}

type ResponseAddress struct {
	Addresses []*Address `json:"addresses"`
}
  
type RequestAddressGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
