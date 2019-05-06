package dpdlib

import "github.com/fiorix/wsdl2go/soap"

var TrackingNamespace = "http://dpd.ru/ws/tracing/2011-11-18"

type parcelTracing struct {
	cli *soap.Client
}

func NewParcelTracing(cli *soap.Client) ParcelTracing {
	return &parcelTracing{cli}
}

type ParcelTracing interface {
	Confirm(Confirm *Confirm) (*ConfirmResponse, error)
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
