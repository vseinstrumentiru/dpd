package dpdlib

import "github.com/fiorix/wsdl2go/soap"

var CalculatorNamespace = "http://dpd.ru/ws/calculator/2012-03-20"

type dPDCalculator struct {
	cli *soap.Client
}

func NewDPDCalculator(cli *soap.Client) DPDCalculator {
	return &dPDCalculator{cli}
}

type DPDCalculator interface {
	GetServiceCost(GetServiceCost *GetServiceCost) (*GetServiceCostResponse, error)
	GetServiceCost2(GetServiceCost2 *GetServiceCost2) (*GetServiceCost2Response, error)
	GetServiceCostByParcels(GetServiceCostByParcels *GetServiceCostByParcels) (*GetServiceCostByParcelsResponse, error)
	GetServiceCostByParcels2(GetServiceCostByParcels2 *GetServiceCostByParcels2) (*GetServiceCostByParcels2Response, error)
	GetServiceCostInternational(GetServiceCostInternational *GetServiceCostInternational) (*GetServiceCostInternationalResponse, error)
}

func (p *dPDCalculator) GetServiceCost(GetServiceCost *GetServiceCost) (*GetServiceCostResponse, error) {
	α := struct {
		OperationGetServiceCost `xml:"tns:getServiceCost"`
	}{
		OperationGetServiceCost{
			GetServiceCost,
		},
	}

	γ := struct {
		OperationGetServiceCostResponse `xml:"getServiceCostResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetServiceCost", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetServiceCostResponse, nil
}

func (p *dPDCalculator) GetServiceCost2(GetServiceCost2 *GetServiceCost2) (*GetServiceCost2Response, error) {
	α := struct {
		OperationGetServiceCost2 `xml:"tns:getServiceCost2"`
	}{
		OperationGetServiceCost2{
			GetServiceCost2,
		},
	}

	γ := struct {
		OperationGetServiceCost2Response `xml:"getServiceCost2Response"`
	}{}
	if err := p.cli.RoundTripWithAction("GetServiceCost2", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetServiceCost2Response, nil
}

func (p *dPDCalculator) GetServiceCostByParcels(GetServiceCostByParcels *GetServiceCostByParcels) (*GetServiceCostByParcelsResponse, error) {
	α := struct {
		OperationGetServiceCostByParcels `xml:"tns:getServiceCostByParcels"`
	}{
		OperationGetServiceCostByParcels{
			GetServiceCostByParcels,
		},
	}

	γ := struct {
		OperationGetServiceCostByParcelsResponse `xml:"getServiceCostByParcelsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetServiceCostByParcels", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetServiceCostByParcelsResponse, nil
}

func (p *dPDCalculator) GetServiceCostByParcels2(GetServiceCostByParcels2 *GetServiceCostByParcels2) (*GetServiceCostByParcels2Response, error) {
	α := struct {
		OperationGetServiceCostByParcels2 `xml:"tns:getServiceCostByParcels2"`
	}{
		OperationGetServiceCostByParcels2{
			GetServiceCostByParcels2,
		},
	}

	γ := struct {
		OperationGetServiceCostByParcels2Response `xml:"getServiceCostByParcels2Response"`
	}{}
	if err := p.cli.RoundTripWithAction("GetServiceCostByParcels2", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetServiceCostByParcels2Response, nil
}

func (p *dPDCalculator) GetServiceCostInternational(GetServiceCostInternational *GetServiceCostInternational) (*GetServiceCostInternationalResponse, error) {
	α := struct {
		OperationGetServiceCostInternational `xml:"tns:getServiceCostInternational"`
	}{
		OperationGetServiceCostInternational{
			GetServiceCostInternational,
		},
	}

	γ := struct {
		OperationGetServiceCostInternationalResponse `xml:"getServiceCostInternationalResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetServiceCostInternational", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetServiceCostInternationalResponse, nil
}
