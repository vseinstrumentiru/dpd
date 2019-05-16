package dpd_sdk

import (
	dpdSoap "git.vseinstrumenti.net/golang-sandbox/dpd-soap"
	"time"
)

type clientOrderRequest dpdSoap.RequestClientOrder

type ClientOrderRequest interface {
	SetClientOrderNumber(number string) *clientOrderRequest
	SetPickupDate(time time.Time) *clientOrderRequest

	toDPDRequest() *dpdSoap.RequestClientOrder
}

func NewClientOrderRequest() *clientOrderRequest {
	return new(clientOrderRequest)
}

func (r *clientOrderRequest) SetClientOrderNumber(number string) *clientOrderRequest {
	r.ClientOrderNr = &number

	return r
}

func (r *clientOrderRequest) SetPickupDate(time time.Time) *clientOrderRequest {
	d := dpdSoap.Date(time.Format("2016-01-02"))
	r.PickupDate = &d

	return r
}

func (r *clientOrderRequest) toDPDRequest() *dpdSoap.RequestClientOrder {
	dpdReq := dpdSoap.RequestClientOrder(*r)

	return &dpdReq
}

type clientParcelRequest dpdSoap.RequestClientParcel

type ClientParcelRequest interface {
	SetClientParcelNumber(number string) *clientParcelRequest
	SetPickupDate(time time.Time) *clientParcelRequest
}

func NewClientParcelRequest() *clientParcelRequest {
	return new(clientParcelRequest)
}

func (r *clientParcelRequest) SetClientParcelNumber(number string) *clientParcelRequest {
	r.ClientParcelNr = &number

	return r
}

func (r *clientParcelRequest) SetPickupDate(time time.Time) *clientParcelRequest {
	d := dpdSoap.Date(time.Format("2006-01-02"))
	r.PickupDate = &d

	return r
}

type dpdOrderRequest dpdSoap.RequestDpdOrder

type DpdOrderRequest interface {
	SetDPDOrderNumber(number string) *dpdOrderRequest
	SetPickupYear(year int) *dpdOrderRequest

	toDPDRequest() *dpdSoap.RequestDpdOrder
}

func NewDpdOrderRequest() *dpdOrderRequest {
	return new(dpdOrderRequest)
}

func (r *dpdOrderRequest) SetDPDOrderNumber(number string) *dpdOrderRequest {
	r.DpdOrderNr = &number

	return r
}

func (r *dpdOrderRequest) SetPickupYear(year int) *dpdOrderRequest {
	r.PickupYear = &year

	return r
}

func (r *dpdOrderRequest) toDPDRequest() *dpdSoap.RequestDpdOrder {
	dpdReq := dpdSoap.RequestDpdOrder(*r)

	return &dpdReq
}

type confirmRequest dpdSoap.RequestConfirm

type ConfirmRequest interface {
	SetDocId(docID int64) *confirmRequest
}

func NewConfirmRequest() *confirmRequest {
	return new(confirmRequest)
}

func (r *confirmRequest) SetDocId(docId int64) *confirmRequest {
	r.DocId = &docId

	return r
}

type dpdTrackingEventRequest dpdSoap.EventTrackingRequest

type DPDTrackingOrderRequest interface {
	SetDateFrom(from time.Time) *dpdTrackingEventRequest
	SetDateTo(to time.Time) *dpdTrackingEventRequest
	SetMaxRowCount(count int) *dpdTrackingEventRequest

	toDPDRequest() *dpdSoap.EventTrackingRequest
}

func (r *dpdTrackingEventRequest) SetDateFrom(from time.Time) *dpdTrackingEventRequest {
	d := dpdSoap.DateTime(from.Format("2006-01-02 15:04:05"))
	r.DateFrom = &d

	return r
}

func (r *dpdTrackingEventRequest) SetDateTo(to time.Time) *dpdTrackingEventRequest {
	d := dpdSoap.DateTime(to.Format("2006-01-02 15:04:05"))
	r.DateTo = &d

	return r
}

func (r *dpdTrackingEventRequest) SetMaxRowCount(count int) *dpdTrackingEventRequest {
	r.MaxRowCount = &count

	return r
}

func (r *dpdTrackingEventRequest) toDPDRequest() *dpdSoap.EventTrackingRequest {
	dpdReq := dpdSoap.EventTrackingRequest(*r)

	return &dpdReq
}
