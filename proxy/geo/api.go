package geo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"os"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
)

type Geo struct {
	Api *suggest.Api
}

func New() *Geo {
	// endpointURL, _ := url.Parse("http://suggestions.dadata.ru/suggestions/api/4_1/rs/")
	creds := client.Credentials{
		ApiKeyValue:    os.Getenv("DADATA_API_KEY"),
		SecretKeyValue: os.Getenv("DADATA_SECRET_KEY"),
	}
	return &Geo{
		Api: dadata.NewSuggestApi(client.WithCredentialProvider(&creds)),
		// Api: &suggest.Api{
		// Client: client.NewClient(endpointURL, client.WithCredentialProvider(&creds)),
		// },
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
	// var query RequestGeocode

	// err := json.NewDecoder(r.Body).Decode(&query)
	// if err != nil {
	// 	fmt.Println("search request error:", err.Error())
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// params := suggest.RequestParams{
	// 	Query: fmt.Sprintf("%s %s", query.Lat, query.Lng),
	// }
	// fmt.Println("Query:", query)
	// fmt.Println("Params:", params)

	// ------------------

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	geocodeURL := "http://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address"

	reqBody := strings.ReplaceAll(string(body), "lng", "lon")
	newReqBody := bytes.NewBuffer([]byte(reqBody))

	fmt.Println("body:", string(body))
	fmt.Println("req body:", reqBody)
	fmt.Println("new req body:", newReqBody)

	client := new(http.Client)
	request, err := http.NewRequest(http.MethodPost, geocodeURL, newReqBody)
	if err != nil {
		fmt.Println("meow!")
		return
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Token %s", os.Getenv("DADATA_API_KEY")))

	res, err := client.Do(request)
	if err != nil {
		log.Println("geocode request error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var suggestions ResponseGeocode
	err = json.NewDecoder(res.Body).Decode(&suggestions)
	if err != nil {
		log.Println(err)
	}

	log.Printf("GEOCODE ADDRESS:  %v", len(suggestions.Suggestions))

	// geocodeURL := "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address"

	// var result = &ResponseGeolocate{}
	// err = g.Api.Client.Post(context.Background(), "geolocate/address", params, result)
	// if err != nil {
	// 	return
	// }

	// fmt.Println("Result:", result.Suggestions)
}
