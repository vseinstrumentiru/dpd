package dpd_sdk

import (
	dpdSoap "git.vseinstrumenti.net/golang-sandbox/dpd-soap"
	"time"
)

type calculateRequest dpdSoap.ServiceCostRequest
type city dpdSoap.CityRequest

type CityRequest interface {
	SetIndex(index string) *city
	SetCityName(name string) *city
	SetRegionCode(code int) *city
	SetCountryCode(code string) *city
}

type CalculateRequest interface {
	SetPickup(city *city) *calculateRequest
	SetDelivery(city *city) *calculateRequest
	SetWeight(weight float64) *calculateRequest
	SetVolume(volume float64) *calculateRequest
	SetServiceCode(code string) *calculateRequest
	SetPickupDate(time time.Time) *calculateRequest
	SetMaxDays(days int) *calculateRequest
	SetMaxCost(cost float64) *calculateRequest
	SetDeclaredValue(declaredValue float64) *calculateRequest
	SetSelfPickup(flag bool) *calculateRequest
	SetSelfDelivery(flag bool) *calculateRequest

	toDpdRequest() *dpdSoap.ServiceCostRequest
}

func NewCity(cityId int64) *city {
	return &city{
		CityId: &cityId,
	}
}

func (c *city) SetIndex(index string) *city {
	c.Index = &index

	return c
}

func (c *city) SetCityName(name string) *city {
	c.CityName = &name

	return c
}

func (c *city) SetRegionCode(code int) *city {
	c.RegionCode = &code

	return c
}

func (c *city) SetCountryCode(code string) *city {
	c.CountryCode = &code

	return c
}

func NewCalculateRequest() *calculateRequest {
	return &calculateRequest{}
}

func (r *calculateRequest) Request() *calculateRequest {
	return r
}

func (r *calculateRequest) SetPickup(city *city) *calculateRequest {
	dpdCityRequest := dpdSoap.CityRequest(*city)
	r.Pickup = &dpdCityRequest

	return r
}

func (r *calculateRequest) SetDelivery(city *city) *calculateRequest {
	delivery := dpdSoap.CityRequest(*city)
	r.Delivery = &delivery

	return r
}

func (r *calculateRequest) SetWeight(weight float64) *calculateRequest {
	r.Weight = &weight

	return r
}

func (r *calculateRequest) SetVolume(volume float64) *calculateRequest {
	r.Volume = &volume

	return r
}

func (r *calculateRequest) SetServiceCode(code string) *calculateRequest {
	r.ServiceCode = &code

	return r
}

func (r *calculateRequest) SetPickupDate(time time.Time) *calculateRequest {
	d := dpdSoap.Date(time.Format("2006-01-02"))
	r.PickupDate = &d

	return r
}

func (r *calculateRequest) SetMaxDays(days int) *calculateRequest {
	r.MaxDays = &days

	return r
}

func (r *calculateRequest) SetMaxCost(cost float64) *calculateRequest {
	r.MaxCost = &cost

	return r
}

func (r *calculateRequest) SetDeclaredValue(declaredValue float64) *calculateRequest {
	r.DeclaredValue = &declaredValue

	return r
}

func (r *calculateRequest) toDpdRequest() *dpdSoap.ServiceCostRequest {
	dpdReq := dpdSoap.ServiceCostRequest(*r)

	return &dpdReq
}

func (r *calculateRequest) SetSelfPickup(flag bool) *calculateRequest {
	r.SelfPickup = &flag

	return r
}

func (r *calculateRequest) SetSelfDelivery(flag bool) *calculateRequest {
	r.SelfDelivery = &flag

	return r
}
