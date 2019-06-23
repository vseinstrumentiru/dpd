package dpd

import (
	"time"

	dpdSoap "github.com/vseinstrumentiru/dpd/soap"
)

//Запрос на трекинг по номеру заказа клиента
type ClientOrderRequest dpdSoap.RequestClientOrder

func NewClientOrderRequest() *ClientOrderRequest {
	return new(ClientOrderRequest)
}

func (r *ClientOrderRequest) SetClientOrderNumber(number string) *ClientOrderRequest {
	r.ClientOrderNr = &number

	return r
}

func (r *ClientOrderRequest) SetPickupDate(time time.Time) *ClientOrderRequest {
	d := dpdSoap.Date(time.Format("2006-01-02"))
	r.PickupDate = &d

	return r
}

func (r *ClientOrderRequest) toDPDRequest() *dpdSoap.RequestClientOrder {
	dpdReq := dpdSoap.RequestClientOrder(*r)

	return &dpdReq
}

//Запрос на трекинг по посылке. Посылка идентифицируется по номеру посылки в информационной системе клиента.
type ClientParcelRequest dpdSoap.RequestClientParcel

func NewClientParcelRequest() *ClientParcelRequest {
	return new(ClientParcelRequest)
}

func (r *ClientParcelRequest) SetClientParcelNumber(number string) *ClientParcelRequest {
	r.ClientParcelNr = &number

	return r
}

func (r *ClientParcelRequest) SetPickupDate(time time.Time) *ClientParcelRequest {
	d := dpdSoap.Date(time.Format("2006-01-02"))
	r.PickupDate = &d

	return r
}

//Запрос на трекинг по номеру заказа DPD
type DpdOrderRequest dpdSoap.RequestDpdOrder

func NewDpdOrderRequest() *DpdOrderRequest {
	return new(DpdOrderRequest)
}

func (r *DpdOrderRequest) SetDPDOrderNumber(number string) *DpdOrderRequest {
	r.DpdOrderNr = &number

	return r
}

func (r *DpdOrderRequest) SetPickupYear(year int) *DpdOrderRequest {
	r.PickupYear = &year

	return r
}

func (r *DpdOrderRequest) toDPDRequest() *dpdSoap.RequestDpdOrder {
	dpdReq := dpdSoap.RequestDpdOrder(*r)

	return &dpdReq
}

// Запрос на подтверждение эвентов получених от GetEvents
type ConfirmRequest dpdSoap.RequestConfirm

func NewConfirmRequest() *ConfirmRequest {
	return new(ConfirmRequest)
}

func (r *ConfirmRequest) SetDocId(docId int64) *ConfirmRequest {
	r.DocId = &docId

	return r
}

type DpdTrackingEventRequest dpdSoap.EventTrackingRequest

func (r *DpdTrackingEventRequest) SetDateFrom(from time.Time) *DpdTrackingEventRequest {
	d := dpdSoap.DateTime(from.Format("2006-01-02 15:04:05"))
	r.DateFrom = &d

	return r
}

func (r *DpdTrackingEventRequest) SetDateTo(to time.Time) *DpdTrackingEventRequest {
	d := dpdSoap.DateTime(to.Format("2006-01-02 15:04:05"))
	r.DateTo = &d

	return r
}

func (r *DpdTrackingEventRequest) SetMaxRowCount(count int) *DpdTrackingEventRequest {
	r.MaxRowCount = &count

	return r
}

func (r *DpdTrackingEventRequest) toDPDRequest() *dpdSoap.EventTrackingRequest {
	dpdReq := dpdSoap.EventTrackingRequest(*r)

	return &dpdReq
}
