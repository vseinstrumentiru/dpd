package dpd

import dpdSoap "github.com/vseinstrumentiru/dpd/soap"

//Request list of self delivery points with their restrictions
type ParcelShopRequest dpdSoap.DpdParcelShopRequest

func NewParcelShopRequest() *ParcelShopRequest {
	return new(ParcelShopRequest)
}

//Set country code according ISO 3166-1 alpha-2 standard
//https://www.iso.org/obp/ui/#search
//If omitted, default RU
func (r *ParcelShopRequest) SetCountryCode(code string) *ParcelShopRequest {
	r.CountryCode = &code

	return r
}

func (r *ParcelShopRequest) SetRegionCode(code string) *ParcelShopRequest {
	r.RegionCode = &code

	return r
}

//Dpd native, city identifier
func (r *ParcelShopRequest) SetCityCode(code string) *ParcelShopRequest {
	r.CityCode = &code

	return r
}

func (r *ParcelShopRequest) SetCityName(name string) *ParcelShopRequest {
	r.CityName = &name

	return r
}

func (r *ParcelShopRequest) toDPDRequest() *dpdSoap.DpdParcelShopRequest {
	dpdReq := dpdSoap.DpdParcelShopRequest(*r)

	return &dpdReq
}
