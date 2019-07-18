package dpd

import (
	"time"
)

//NewCity creates new pointer to CityRequest
//use by CalculateRequest
func NewCity(cityID int64) *CityRequest {
	return &CityRequest{
		CityID: &cityID,
	}
}

//SetIndex set zip code of city
func (c *CityRequest) SetIndex(index string) *CityRequest {
	c.Index = &index

	return c
}

//SetCityName set city name for city request
func (c *CityRequest) SetCityName(name string) *CityRequest {
	c.CityName = &name

	return c
}

//SetRegionCode set region code for city request
func (c *CityRequest) SetRegionCode(code int) *CityRequest {
	c.RegionCode = &code

	return c
}

//SetCountryCode the code must comply with ISO 3166-1 alpha-2
//If omitted, default RU
func (c *CityRequest) SetCountryCode(code string) *CityRequest {
	c.CountryCode = &code

	return c
}

//NewCalculateRequest creates new pointer to CalculateRequest with required parameters
func NewCalculateRequest(from, to *CityRequest, weight float64, selfPickup, selfDelivery bool) *CalculateRequest {
	req := new(CalculateRequest)

	req.Pickup = from
	req.Delivery = to
	req.Weight = &weight
	req.SelfPickup = &selfPickup
	req.SelfDelivery = &selfDelivery

	return req
}

//OverrideFrom overrides pickup (city-sender) field in CalculateRequest struct
func (r *CalculateRequest) OverrideFrom(city *CityRequest) *CalculateRequest {
	r.Pickup = city

	return r
}

//OverrideTo overrides delivery (city-recipient) field in CalculateRequest struct
func (r *CalculateRequest) OverrideTo(city *CityRequest) *CalculateRequest {
	r.Delivery = city

	return r
}

//SetWeight set weight of parcel
func (r *CalculateRequest) SetWeight(weight float64) *CalculateRequest {
	r.Weight = &weight

	return r
}

//SetVolume set volume of parcel
func (r *CalculateRequest) SetVolume(volume float64) *CalculateRequest {
	r.Volume = &volume

	return r
}

//SetServiceCode  set service codes - list of comma-separated values
//If set, DPD service will return cost only for given service codes
func (r *CalculateRequest) SetServiceCode(code string) *CalculateRequest {
	r.ServiceCode = &code

	return r
}

//SetPickupDate set date of parcel pickup& If not set DPD use current date by default
func (r *CalculateRequest) SetPickupDate(time time.Time) *CalculateRequest {
	d := time.Format("2006-01-02")
	r.PickupDate = &d

	return r
}

//SetMaxDays If specific service is set up for request, call of this function with any parameter
//has not affect result of request
func (r *CalculateRequest) SetMaxDays(days int) *CalculateRequest {
	r.MaxDays = &days

	return r
}

//SetMaxCost If specific service is set up for request, call of this function with any parameter
//has not affect result of request
func (r *CalculateRequest) SetMaxCost(cost float64) *CalculateRequest {
	r.MaxCost = &cost

	return r
}

//SetDeclaredValue set cargo declared value
func (r *CalculateRequest) SetDeclaredValue(declaredValue float64) *CalculateRequest {
	r.DeclaredValue = &declaredValue

	return r
}

//SetSelfPickup set client-side delivery to DPD terminal
func (r *CalculateRequest) SetSelfPickup(flag bool) *CalculateRequest {
	r.SelfPickup = &flag

	return r
}

//SetSelfDelivery set DPD-side delivery to their terminal and customer-side pickup on DPD terminal
func (r *CalculateRequest) SetSelfDelivery(flag bool) *CalculateRequest {
	r.SelfDelivery = &flag

	return r
}
