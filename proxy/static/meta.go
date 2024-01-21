// Package classification Geo API.
//
// Документация Geo API.
//
//	Schemes: http, https
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//	- multipart/form-data
//
//	Produces:
//	- application/json
//
//	Security:
//	- basic
//
//
//	SecurityDefinitions:
//	  Bearer:
//	    type: apiKey
//	    name: Authorization
//	    in: header
//
// swagger:meta
package static

import "test/geo"

//go:generate `swagger generate spec -o /go/src/proxy/public/swagger.json --scan-models`

// swagger:route POST /api/address/geocode geo RequestGeocode
//		Получение адреса по долготе и широте
// Security:
// - Bearer: []
// Responses:
//   200: body:ResponseGeocode

// RequestGeocode security:
// - Bearer: []

// swagger:parameters RequestGeocode
type RequestGeocode struct {
	// Координаты
	// in:body
	// required:true
	// example: {"lat":"59.7221","lng":"30.4554"}
	Coordinates string
}

// swagger:model ResponseGeocode
type ResponseGeocode struct {
	// in: body
	// Addresses содержит информацию о адресах
	Addresses geo.Address
}

// swagger:route POST /api/address/search geo RequestAddress
//		Получение координат по адресу
// Security:
// - Bearer: []
// Responses:
//	 200: body:ResponseAddress

// RequestAddress security:
// - Bearer: []

// swagger:parameters RequestAddress
type RequestAddress struct {
	// Адрес
	// in:body
	// required:true
	// example: {"query": "москва тверская"}
	Query string
}

// swagger:model ResponseAddress
type ResponseAddress struct {
	// in:body
	// Addresses содержит информацию о координатах
	Addresses geo.Address
}
