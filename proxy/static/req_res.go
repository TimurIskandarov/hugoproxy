package static

// swagger:route POST /api/address/geocode geo RequestGeocode
//		Получение адреса по широте и долготе
// Security:
// - Bearer: []
// Responses:
// 	 200: ResponseGeocode

// RequestGeocode security:
// -Bearer: []
// swagger:parameters RequestGeocode
type RequestGeocode struct {
	// in:body
	// required:true
	// example: {"lat":"59.7221","lon":"30.4554"}
	Coordinates string
}
