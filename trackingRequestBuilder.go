package dpd

import (
	"time"
)

func NewClientOrderRequest() *TrackByClientOrderRequest {
	return new(TrackByClientOrderRequest)
}

func (r *TrackByClientOrderRequest) SetClientOrderNumber(number string) *TrackByClientOrderRequest {
	r.ClientOrderNr = &number

	return r
}

func (r *TrackByClientOrderRequest) SetPickupDate(time time.Time) *TrackByClientOrderRequest {
	d := time.Format("2006-01-02")
	r.PickupDate = &d

	return r
}

func NewDpdOrderRequest() *TrackByDPDOrderRequest {
	return new(TrackByDPDOrderRequest)
}

func (r *TrackByDPDOrderRequest) SetDPDOrderNumber(number string) *TrackByDPDOrderRequest {
	r.DpdOrderNr = &number

	return r
}

func (r *TrackByDPDOrderRequest) SetPickupYear(year int) *TrackByDPDOrderRequest {
	r.PickupYear = &year

	return r
}
