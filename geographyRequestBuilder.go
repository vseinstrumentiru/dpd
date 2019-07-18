package dpd

//NewParcelShopRequest returns pointer to ParcelShopRequest
func NewParcelShopRequest() *ParcelShopRequest {
	return new(ParcelShopRequest)
}

//SetCountryCode the code must comply with ISO 3166-1 alpha-2
//If omitted, default RU
func (r *ParcelShopRequest) SetCountryCode(code string) *ParcelShopRequest {
	r.CountryCode = &code

	return r
}

//SetRegionCode parameter must comply DPD native region code
func (r *ParcelShopRequest) SetRegionCode(code string) *ParcelShopRequest {
	r.RegionCode = &code

	return r
}

//SetCityCode parameter must comply DPD native city code
func (r *ParcelShopRequest) SetCityCode(code string) *ParcelShopRequest {
	r.CityCode = &code

	return r
}

//SetCityName ...
func (r *ParcelShopRequest) SetCityName(name string) *ParcelShopRequest {
	r.CityName = &name

	return r
}
