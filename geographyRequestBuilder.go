package dpd_sdk

import dpdSoap "dpd-soap"

//Запрос на получение списка точек самовывоза с ограничениями
type ParcelShopRequest dpdSoap.DpdParcelShopRequest

func NewParcelShopRequest() *ParcelShopRequest {
	return new(ParcelShopRequest)
}

func (r *ParcelShopRequest) SetCountryCode(code string) *ParcelShopRequest {
	r.CountryCode = &code

	return r
}

func (r *ParcelShopRequest) SetRegionCode(code string) *ParcelShopRequest {
	r.RegionCode = &code

	return r
}

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
