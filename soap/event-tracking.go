package soap

import (
	"github.com/fiorix/wsdl2go/soap"
)

const EventTrackingNamespace = "http://dpd.ru/ws/event-tracking/2016-02-15"

func NewDPDEventTracking(cli *soap.Client) DPDEventTracking {
	return &dPDEventTracking{cli}
}

type DPDEventTracking interface {
	Confirm(Confirm *Confirm) (*ConfirmResponse, error)
	GetEvents(GetEvents *GetEvents) (*GetEventsResponse, error)
	GetTrakingOrderLink(GetTrakingOrderLink *GetTrakingOrderLink) (*GetTrakingOrderLinkResponse, error)
}

type Confirm struct {
	Request *RequestConfirm `xml:"request,omitempty"`
}

type ConfirmResponse struct {
	Return *string `xml:"return,omitempty"`
}

type EventTackingResponse struct {
	DocId          *int64       `xml:"docId,omitempty"`
	DocDate        *DateTime    `xml:"docDate,omitempty"`
	ClientNumber   *int64       `xml:"clientNumber,omitempty"`
	ResultComplete *bool        `xml:"resultComplete,omitempty"`
	Event          []*EventType `xml:"event,omitempty"`
}

type EventTrackingRequest struct {
	Auth        *Auth     `xml:"auth,omitempty"`
	DateFrom    *DateTime `xml:"dateFrom,omitempty"`
	DateTo      *DateTime `xml:"dateTo,omitempty"`
	MaxRowCount *int      `xml:"maxRowCount,omitempty"`
}

type EventTrackingUnitLoad struct {
	Article        *string `xml:"article,omitempty"`
	Descript       *string `xml:"descript,omitempty"`
	Declared_value *string `xml:"declared_value,omitempty"`
	Parcel_num     *string `xml:"parcel_num,omitempty"`
	Npp_amount     *string `xml:"npp_amount,omitempty"`
	Vat_percent    *int    `xml:"vat_percent,omitempty"`
	Without_vat    *bool   `xml:"without_vat,omitempty"`
	Count          *int    `xml:"count,omitempty"`
	State_name     *string `xml:"state_name,omitempty"`
}

type EventType struct {
	ClientOrderNr      *string          `xml:"clientOrderNr,omitempty"`
	ClientCodeInternal *string          `xml:"clientCodeInternal,omitempty"`
	ClientParcelNr     *string          `xml:"clientParcelNr,omitempty"`
	DpdOrderNr         *string          `xml:"dpdOrderNr,omitempty"`
	DpdParcelNr        *string          `xml:"dpdParcelNr,omitempty"`
	EventNumber        *string          `xml:"eventNumber,omitempty"`
	EventCode          *string          `xml:"eventCode,omitempty"`
	EventName          *string          `xml:"eventName,omitempty"`
	ReasonCode         *string          `xml:"reasonCode,omitempty"`
	ReasonName         *string          `xml:"reasonName,omitempty"`
	EventDate          *string          `xml:"eventDate,omitempty"`
	Parameter          []*ParameterType `xml:"parameter,omitempty"`
}

type ExtTrackingLinkRequest struct {
	Auth  *Auth          `xml:"auth,omitempty"`
	Order *TrackingOrder `xml:"order,omitempty"`
}

type ExtTrackingLinkResponse struct {
	Link   *string `xml:"link,omitempty"`
	Status *string `xml:"status,omitempty"`
}

type GetEvents struct {
	Request *EventTrackingRequest `xml:"request,omitempty"`
}

type GetEventsResponse struct {
	Return *EventTackingResponse `xml:"return,omitempty"`
}

type GetTrakingOrderLink struct {
	TrackingOrder *ExtTrackingLinkRequest `xml:"trackingOrder,omitempty"`
}

type GetTrakingOrderLinkResponse struct {
	Return *ExtTrackingLinkResponse `xml:"return,omitempty"`
}

type TrackingOrder struct {
	OrderNumberDPD                *string `xml:"orderNumberDPD,omitempty"`
	OrderNumberInternal           *string `xml:"orderNumberInternal,omitempty"`
	OrderNumberInternalAdditional *string `xml:"orderNumberInternalAdditional,omitempty"`
}

type ParameterType struct {
	ParamName     *string                  `xml:"paramName,omitempty"`
	ValueString   *string                  `xml:"valueString,omitempty"`
	ValueDecimal  *float64                 `xml:"valueDecimal,omitempty"`
	ValueDateTime *DateTime                `xml:"valueDateTime,omitempty"`
	Value         []*EventTrackingUnitLoad `xml:"value,omitempty"`
	Type          string                   `xml:"type,attr,omitempty"`
}

type RequestConfirm struct {
	Auth  *Auth  `xml:"auth,omitempty"`
	DocId *int64 `xml:"docId,omitempty"`
}

type OperationConfirm struct {
	Confirm *Confirm `xml:"confirm,omitempty"`
}

type OperationConfirmResponse struct {
	ConfirmResponse *ConfirmResponse `xml:"confirmResponse,omitempty"`
}

type OperationGetEvents struct {
	GetEvents *GetEvents `xml:"getEvents,omitempty"`
}

type OperationGetEventsResponse struct {
	GetEventsResponse *GetEventsResponse `xml:"getEventsResponse,omitempty"`
}

type OperationGetTrakingOrderLink struct {
	GetTrakingOrderLink *GetTrakingOrderLink `xml:"getTrakingOrderLink,omitempty"`
}

type OperationGetTrakingOrderLinkResponse struct {
	GetTrakingOrderLinkResponse *GetTrakingOrderLinkResponse `xml:"getTrakingOrderLinkResponse,omitempty"`
}

type dPDEventTracking struct {
	cli *soap.Client
}

func (p *dPDEventTracking) Confirm(Confirm *Confirm) (*ConfirmResponse, error) {
	α := struct {
		OperationConfirm `xml:"tns:confirm"`
	}{
		OperationConfirm{
			Confirm,
		},
	}

	γ := struct {
		OperationConfirmResponse `xml:"confirmResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("Confirm", α, &γ); err != nil {
		return nil, err
	}
	return γ.ConfirmResponse, nil
}

func (p *dPDEventTracking) GetEvents(GetEvents *GetEvents) (*GetEventsResponse, error) {
	α := struct {
		OperationGetEvents `xml:"tns:getEvents"`
	}{
		OperationGetEvents{
			GetEvents,
		},
	}

	γ := struct {
		OperationGetEventsResponse `xml:"getEventsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetEvents", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetEventsResponse, nil
}

func (p *dPDEventTracking) GetTrakingOrderLink(GetTrakingOrderLink *GetTrakingOrderLink) (*GetTrakingOrderLinkResponse, error) {
	α := struct {
		OperationGetTrakingOrderLink `xml:"tns:getTrakingOrderLink"`
	}{
		OperationGetTrakingOrderLink{
			GetTrakingOrderLink,
		},
	}

	γ := struct {
		OperationGetTrakingOrderLinkResponse `xml:"getTrakingOrderLinkResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetTrakingOrderLink", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetTrakingOrderLinkResponse, nil
}
