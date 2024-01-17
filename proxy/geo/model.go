package geo

// API: подсказки по адресам https://dadata.ru/api/suggest/address/
type RequestAddress struct {
	Query string `json:"query"`
}


type ResponseAddress struct {
	Addresses []*Address `json:"addresses"`
}

type Address struct {
	Value string `json:"value"`
	Lat   string `json:"lat"`
	Lng   string `json:"lon"`
}


type RequestGeocode struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}


type ResponseGeocode struct {
	Suggestions []*Geolocate `json:"suggestions"`
}

type Geolocate struct {
	Value             string `json:"value"`
	UnrestrictedValue string `json:"unrestricted_value"`
	Data              Data   `json:"data"`
}

type Data struct {
}