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

// API: обратное геокодирование https://dadata.ru/api/geolocate/
type ResponseGeocode struct {
	Addresses []*Address `json:"addresses"`
}

// GetGeocode
type ResponseSuggestionGeocode struct {
	Suggestions []*SuggestionGeocode `json:"suggestions"`
}

type SuggestionGeocode struct {
	Value             string `json:"value"`
	UnrestrictedValue string `json:"unrestricted_value"`
	Data              Data   `json:"data"`
}

type Data struct {
	Lat string `json:"geo_lat"`
	Lng string `json:"geo_lon"`
}
