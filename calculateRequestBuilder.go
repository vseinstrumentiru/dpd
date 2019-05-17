package dpd_sdk

import (
	dpdSoap "git.vseinstrumenti.net/golang-sandbox/dpd-sdk/dpd-soap"
	"time"
)

//Запрос на подсчет стоимости доставки
type CalculateRequest dpdSoap.ServiceCostRequest

//Адрес для подсчета стомиости доставки
type CityRequest dpdSoap.CityRequest

func NewCity(cityId int64) *CityRequest {
	return &CityRequest{
		CityId: &cityId,
	}
}

func (c *CityRequest) SetIndex(index string) *CityRequest {
	c.Index = &index

	return c
}

func (c *CityRequest) SetCityName(name string) *CityRequest {
	c.CityName = &name

	return c
}

func (c *CityRequest) SetRegionCode(code int) *CityRequest {
	c.RegionCode = &code

	return c
}

func (c *CityRequest) SetCountryCode(code string) *CityRequest {
	c.CountryCode = &code

	return c
}

func NewCalculateRequest() *CalculateRequest {
	return &CalculateRequest{}
}

func (r *CalculateRequest) SetPickup(city *CityRequest) *CalculateRequest {
	dpdCityRequest := dpdSoap.CityRequest(*city)
	r.Pickup = &dpdCityRequest

	return r
}

func (r *CalculateRequest) SetDelivery(city *CityRequest) *CalculateRequest {
	delivery := dpdSoap.CityRequest(*city)
	r.Delivery = &delivery

	return r
}

func (r *CalculateRequest) SetWeight(weight float64) *CalculateRequest {
	r.Weight = &weight

	return r
}

func (r *CalculateRequest) SetVolume(volume float64) *CalculateRequest {
	r.Volume = &volume

	return r
}

func (r *CalculateRequest) SetServiceCode(code string) *CalculateRequest {
	r.ServiceCode = &code

	return r
}

func (r *CalculateRequest) SetPickupDate(time time.Time) *CalculateRequest {
	d := dpdSoap.Date(time.Format("2006-01-02"))
	r.PickupDate = &d

	return r
}

func (r *CalculateRequest) SetMaxDays(days int) *CalculateRequest {
	r.MaxDays = &days

	return r
}

func (r *CalculateRequest) SetMaxCost(cost float64) *CalculateRequest {
	r.MaxCost = &cost

	return r
}

func (r *CalculateRequest) SetDeclaredValue(declaredValue float64) *CalculateRequest {
	r.DeclaredValue = &declaredValue

	return r
}

func (r *CalculateRequest) toDpdRequest() *dpdSoap.ServiceCostRequest {
	dpdReq := dpdSoap.ServiceCostRequest(*r)

	return &dpdReq
}

func (r *CalculateRequest) SetSelfPickup(flag bool) *CalculateRequest {
	r.SelfPickup = &flag

	return r
}

func (r *CalculateRequest) SetSelfDelivery(flag bool) *CalculateRequest {
	r.SelfDelivery = &flag

	return r
}
