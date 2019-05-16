package dpd_sdk

import dpdSoap "git.vseinstrumenti.net/golang-sandbox/dpd-soap"

type parcelShopRequest dpdSoap.DpdParcelShopRequest

type ParcelShopRequest interface {
	SetCountryCode(code string) *parcelShopRequest
	SetRegionCode(code string) *parcelShopRequest
	SetCityCode(code string) *parcelShopRequest
	SetCityName(name string) *parcelShopRequest

	toDPDRequest() *dpdSoap.DpdParcelShopRequest
}

func NewParcelShopRequest() *parcelShopRequest {
	return new(parcelShopRequest)
}

func (r *parcelShopRequest) SetCountryCode(code string) *parcelShopRequest {
	r.CountryCode = &code

	return r
}

func (r *parcelShopRequest) SetRegionCode(code string) *parcelShopRequest {
	r.RegionCode = &code

	return r
}

func (r *parcelShopRequest) SetCityCode(code string) *parcelShopRequest {
	r.CityCode = &code

	return r
}

func (r *parcelShopRequest) SetCityName(name string) *parcelShopRequest {
	r.CityName = &name

	return r
}

func (r *parcelShopRequest) toDPDRequest() *dpdSoap.DpdParcelShopRequest {
	dpdReq := dpdSoap.DpdParcelShopRequest(*r)

	return &dpdReq
}
