package dpd_soap

import "github.com/fiorix/wsdl2go/soap"

const TrackingNamespace = "http://dpd.ru/ws/tracing/2011-11-18"

type parcelTracing struct {
	cli *soap.Client
}

func NewParcelTracing(cli *soap.Client) ParcelTracing {
	return &parcelTracing{cli}
}

type ParcelTracing interface {
	GetStatesByClient(GetStatesByClient *GetStatesByClient) (*GetStatesByClientResponse, error)
	GetStatesByClientOrder(GetStatesByClientOrder *GetStatesByClientOrder) (*GetStatesByClientOrderResponse, error)
	GetStatesByClientParcel(GetStatesByClientParcel *GetStatesByClientParcel) (*GetStatesByClientParcelResponse, error)
	GetStatesByDPDOrder(GetStatesByDPDOrder *GetStatesByDPDOrder) (*GetStatesByDPDOrderResponse, error)
}

func (p *parcelTracing) Confirm(Confirm *Confirm) (*ConfirmResponse, error) {
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

func (p *parcelTracing) GetStatesByClient(GetStatesByClient *GetStatesByClient) (*GetStatesByClientResponse, error) {
	α := struct {
		OperationGetStatesByClient `xml:"tns:getStatesByClient"`
	}{
		OperationGetStatesByClient{
			GetStatesByClient,
		},
	}

	γ := struct {
		OperationGetStatesByClientResponse `xml:"getStatesByClientResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetStatesByClient", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetStatesByClientResponse, nil
}

func (p *parcelTracing) GetStatesByClientOrder(GetStatesByClientOrder *GetStatesByClientOrder) (*GetStatesByClientOrderResponse, error) {
	α := struct {
		OperationGetStatesByClientOrder `xml:"tns:getStatesByClientOrder"`
	}{
		OperationGetStatesByClientOrder{
			GetStatesByClientOrder,
		},
	}

	γ := struct {
		OperationGetStatesByClientOrderResponse `xml:"getStatesByClientOrderResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetStatesByClientOrder", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetStatesByClientOrderResponse, nil
}

func (p *parcelTracing) GetStatesByClientParcel(GetStatesByClientParcel *GetStatesByClientParcel) (*GetStatesByClientParcelResponse, error) {
	α := struct {
		OperationGetStatesByClientParcel `xml:"tns:getStatesByClientParcel"`
	}{
		OperationGetStatesByClientParcel{
			GetStatesByClientParcel,
		},
	}

	γ := struct {
		OperationGetStatesByClientParcelResponse `xml:"getStatesByClientParcelResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetStatesByClientParcel", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetStatesByClientParcelResponse, nil
}

func (p *parcelTracing) GetStatesByDPDOrder(GetStatesByDPDOrder *GetStatesByDPDOrder) (*GetStatesByDPDOrderResponse, error) {
	α := struct {
		OperationGetStatesByDPDOrder `xml:"tns:getStatesByDPDOrder"`
	}{
		OperationGetStatesByDPDOrder{
			GetStatesByDPDOrder,
		},
	}

	γ := struct {
		OperationGetStatesByDPDOrderResponse `xml:"getStatesByDPDOrderResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetStatesByDPDOrder", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetStatesByDPDOrderResponse, nil
}

type GetStatesByClient struct {
	Request *RequestClient `xml:"request,omitempty"`
}

type GetStatesByClientOrder struct {
	Request *RequestClientOrder `xml:"request,omitempty"`
}

type GetStatesByClientOrderResponse struct {
	Return *StateParcels `xml:"return,omitempty"`
}

type GetStatesByClientParcel struct {
	Request *RequestClientParcel `xml:"request,omitempty"`
}

type GetStatesByClientParcelResponse struct {
	Return *StateParcels `xml:"return,omitempty"`
}

type GetStatesByClientResponse struct {
	Return *StateParcels `xml:"return,omitempty"`
}

type GetStatesByDPDOrder struct {
	Request *RequestDpdOrder `xml:"request,omitempty"`
}

type GetStatesByDPDOrderResponse struct {
	Return *StateParcels `xml:"return,omitempty"`
}

type RequestClient struct {
	Auth *Auth `xml:"auth,omitempty"`
}

type RequestClientOrder struct {
	Auth          *Auth   `xml:"auth,omitempty"`
	ClientOrderNr *string `xml:"clientOrderNr,omitempty"`
	PickupDate    *Date   `xml:"pickupDate,omitempty"`
}

type RequestClientParcel struct {
	Auth           *Auth   `xml:"auth,omitempty"`
	ClientParcelNr *string `xml:"clientParcelNr,omitempty"`
	PickupDate     *Date   `xml:"pickupDate,omitempty"`
}

type RequestDpdOrder struct {
	Auth       *Auth   `xml:"auth,omitempty"`
	DpdOrderNr *string `xml:"dpdOrderNr,omitempty"`
	PickupYear *int    `xml:"pickupYear,omitempty"`
}

type StateParcel struct {
	ClientOrderNr    *string   `xml:"clientOrderNr,omitempty"`
	ClientParcelNr   *string   `xml:"clientParcelNr,omitempty"`
	DpdOrderNr       *string   `xml:"dpdOrderNr,omitempty"`
	DpdParcelNr      *string   `xml:"dpdParcelNr,omitempty"`
	PickupDate       *Date     `xml:"pickupDate,omitempty"`
	DpdOrderReNr     *string   `xml:"dpdOrderReNr,omitempty"`
	DpdParcelReNr    *string   `xml:"dpdParcelReNr,omitempty"`
	IsReturn         *bool     `xml:"isReturn,omitempty"`
	PlanDeliveryDate *Date     `xml:"planDeliveryDate,omitempty"`
	NewState         *string   `xml:"newState,omitempty"`
	TransitionTime   *DateTime `xml:"transitionTime,omitempty"`
	TerminalCode     *string   `xml:"terminalCode,omitempty"`
	TerminalCity     *string   `xml:"terminalCity,omitempty"`
	IncidentCode     *string   `xml:"incidentCode,omitempty"`
	IncidentName     *string   `xml:"incidentName,omitempty"`
	Consignee        *string   `xml:"consignee,omitempty"`
}

type StateParcels struct {
	DocId          *int64         `xml:"docId,omitempty"`
	DocDate        *DateTime      `xml:"docDate,omitempty"`
	ClientNumber   *int64         `xml:"clientNumber,omitempty"`
	ResultComplete *bool          `xml:"resultComplete,omitempty"`
	States         []*StateParcel `xml:"states,omitempty"`
}

type OperationGetStatesByClient struct {
	GetStatesByClient *GetStatesByClient `xml:"getStatesByClient,omitempty"`
}

type OperationGetStatesByClientResponse struct {
	GetStatesByClientResponse *GetStatesByClientResponse `xml:"getStatesByClientResponse,omitempty"`
}

type OperationGetStatesByClientOrder struct {
	GetStatesByClientOrder *GetStatesByClientOrder `xml:"getStatesByClientOrder,omitempty"`
}

type OperationGetStatesByClientOrderResponse struct {
	GetStatesByClientOrderResponse *GetStatesByClientOrderResponse `xml:"getStatesByClientOrderResponse,omitempty"`
}

type OperationGetStatesByClientParcel struct {
	GetStatesByClientParcel *GetStatesByClientParcel `xml:"getStatesByClientParcel,omitempty"`
}

type OperationGetStatesByClientParcelResponse struct {
	GetStatesByClientParcelResponse *GetStatesByClientParcelResponse `xml:"getStatesByClientParcelResponse,omitempty"`
}

type OperationGetStatesByDPDOrder struct {
	GetStatesByDPDOrder *GetStatesByDPDOrder `xml:"getStatesByDPDOrder,omitempty"`
}

type OperationGetStatesByDPDOrderResponse struct {
	GetStatesByDPDOrderResponse *GetStatesByDPDOrderResponse `xml:"getStatesByDPDOrderResponse,omitempty"`
}
