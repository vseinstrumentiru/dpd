package dpd

import (
	"time"

	dpdSoap "github.com/vseinstrumentiru/dpd/soap"
)

//Request of calculate delivery cost
type CalculateRequest dpdSoap.ServiceCostRequest

//Address for calculate request
type CityRequest dpdSoap.CityRequest

//Dpd native, city identifier
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

//Set country code according ISO 3166-1 alpha-2 standard
//
//https://www.iso.org/obp/ui/#search
//
//If omitted, default RU
func (c *CityRequest) SetCountryCode(code string) *CityRequest {
	c.CountryCode = &code

	return c
}

func NewCalculateRequest(from, to *CityRequest, weight float64, selfPickup, selfDelivery bool) *CalculateRequest {
	req := new(CalculateRequest)

	dpdFrom := dpdSoap.CityRequest(*from)
	dpdTo := dpdSoap.CityRequest(*to)

	req.Pickup = &dpdFrom
	req.Delivery = &dpdTo
	req.Weight = &weight
	req.SelfPickup = &selfPickup
	req.SelfDelivery = &selfDelivery

	return req
}

//Set pickup address
func (r *CalculateRequest) OverridePickup(city *CityRequest) *CalculateRequest {
	dpdCityRequest := dpdSoap.CityRequest(*city)
	r.Pickup = &dpdCityRequest

	return r
}

//Set delivery address
func (r *CalculateRequest) OverrideDelivery(city *CityRequest) *CalculateRequest {
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

//List of services codes.
//If set, DPD service will return cost only for given service codes
//code - list of codes, comma separated
func (r *CalculateRequest) SetServiceCode(code string) *CalculateRequest {
	r.ServiceCode = &code

	return r
}

//If not set DPD use current date by default
func (r *CalculateRequest) SetPickupDate(time time.Time) *CalculateRequest {
	d := dpdSoap.Date(time.Format("2006-01-02"))
	r.PickupDate = &d

	return r
}

//If specific service is set up for request, call of this function with any parameter
//has not affect result of request
func (r *CalculateRequest) SetMaxDays(days int) *CalculateRequest {
	r.MaxDays = &days

	return r
}

//If specific service is set up for request, call of this function with any parameter
//has not affect result of request
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

//Set client-side delivery to DPD terminal
func (r *CalculateRequest) SetSelfPickup(flag bool) *CalculateRequest {
	r.SelfPickup = &flag

	return r
}

//Set DPD-side delivery to their terminal and customer-side pickup on DPD terminal (o_O)
func (r *CalculateRequest) SetSelfDelivery(flag bool) *CalculateRequest {
	r.SelfDelivery = &flag

	return r
}
