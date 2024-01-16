package geo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"os"

	dadata "github.com/ekomobile/dadata/v2"
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

	for _, s := range suggestions {
		fmt.Printf("Suggestion: %v\n", s)
		fmt.Printf("Value: %v\n", s.Value)
		fmt.Printf("Result: %v\n", s.Data.Result)
	}

	// TODO

}
