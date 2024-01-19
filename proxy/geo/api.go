package geo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
)

type Geo struct {
	Api *suggest.Api
}

func New() *Geo {
	creds := client.Credentials{
		ApiKeyValue:    os.Getenv("DADATA_API_KEY"),
		SecretKeyValue: os.Getenv("DADATA_SECRET_KEY"),
	}
	return &Geo{
		Api: dadata.NewSuggestApi(client.WithCredentialProvider(&creds)),
	}
}

func (g *Geo) SearchHandler(w http.ResponseWriter, r *http.Request) {
	var query RequestAddress

	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := suggest.RequestParams{
		Query: query.Query,
	}

	suggestions, err := g.Api.Address(context.Background(), &params)
	if err != nil {
		fmt.Println("search request error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ResponseAddress{
		Addresses: make([]*Address, len(suggestions)),
	}

	for i, s := range suggestions {
		response.Addresses[i] = &Address{
			Value: s.Value,
			Lat:   s.Data.GeoLat,
			Lng:   s.Data.GeoLon,
		}

		log.Printf("addr: %#v\n", response.Addresses[i])
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("search response error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (g *Geo) GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	suggestions, err := g.GetGeocode(r)
	if err != nil {
		fmt.Println("geocode request error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ResponseGeocode{
		Addresses: make([]*Address, len(suggestions)),
	}

	for i, s := range suggestions {
		response.Addresses[i] = &Address{
			Value: s.Value,
			Lat:   s.Data.Lat,
			Lng:   s.Data.Lng,
		}

		log.Printf("geocode: %#v\n", s)
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("search response error: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (g *Geo) GetGeocode(r *http.Request) ([]*SuggestionGeocode, error) {
	var query RequestGeocode
	json.NewDecoder(r.Body).Decode(&query)

	querySuggestion := RequestSuggestionGeocode{
		Lat: query.Lat,
		Lng: query.Lng,
	}

	var result = &ResponseSuggestionGeocode{}
	err := g.Api.Client.Post(context.Background(), "geolocate/address", &querySuggestion, result)
	if err != nil {
		return nil, err
	}
	
	return result.Suggestions, nil
}
