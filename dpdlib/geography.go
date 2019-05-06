package dpdlib

import "github.com/fiorix/wsdl2go/soap"

var GeographyNamespace = "http://dpd.ru/ws/geography/2015-05-20"

func NewDPDGeography2(cli *soap.Client) DPDGeography2 {
	return &dPDGeography2{cli}
}

type DPDGeography2 interface {
	GetCitiesCashPay(GetCitiesCashPay *GetCitiesCashPay) (*GetCitiesCashPayResponse, error)
	GetParcelShops(GetParcelShops *GetParcelShops) (*GetParcelShopsResponse, error)
	GetPossibleExtraService(GetPossibleExtraService *GetPossibleExtraService) (*GetPossibleExtraServiceResponse, error)
	GetStoragePeriod(GetStoragePeriod *GetStoragePeriod) (*GetStoragePeriodResponse, error)
	GetTerminalsSelfDelivery2(GetTerminalsSelfDelivery2 *GetTerminalsSelfDelivery2) (*GetTerminalsSelfDelivery2Response, error)
}

func (p *dPDGeography2) GetCitiesCashPay(GetCitiesCashPay *GetCitiesCashPay) (*GetCitiesCashPayResponse, error) {
	α := struct {
		OperationGetCitiesCashPay `xml:"tns:getCitiesCashPay"`
	}{
		OperationGetCitiesCashPay{
			GetCitiesCashPay,
		},
	}

	γ := struct {
		OperationGetCitiesCashPayResponse `xml:"getCitiesCashPayResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetCitiesCashPay", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetCitiesCashPayResponse, nil
}

func (p *dPDGeography2) GetParcelShops(GetParcelShops *GetParcelShops) (*GetParcelShopsResponse, error) {
	α := struct {
		OperationGetParcelShops `xml:"tns:getParcelShops"`
	}{
		OperationGetParcelShops{
			GetParcelShops,
		},
	}

	γ := struct {
		OperationGetParcelShopsResponse `xml:"getParcelShopsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetParcelShops", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetParcelShopsResponse, nil
}

func (p *dPDGeography2) GetPossibleExtraService(GetPossibleExtraService *GetPossibleExtraService) (*GetPossibleExtraServiceResponse, error) {
	α := struct {
		OperationGetPossibleExtraService `xml:"tns:getPossibleExtraService"`
	}{
		OperationGetPossibleExtraService{
			GetPossibleExtraService,
		},
	}

	γ := struct {
		OperationGetPossibleExtraServiceResponse `xml:"getPossibleExtraServiceResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetPossibleExtraService", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetPossibleExtraServiceResponse, nil
}

func (p *dPDGeography2) GetStoragePeriod(GetStoragePeriod *GetStoragePeriod) (*GetStoragePeriodResponse, error) {
	α := struct {
		OperationGetStoragePeriod `xml:"tns:getStoragePeriod"`
	}{
		OperationGetStoragePeriod{
			GetStoragePeriod,
		},
	}

	γ := struct {
		OperationGetStoragePeriodResponse `xml:"getStoragePeriodResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetStoragePeriod", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetStoragePeriodResponse, nil
}

func (p *dPDGeography2) GetTerminalsSelfDelivery2(GetTerminalsSelfDelivery2 *GetTerminalsSelfDelivery2) (*GetTerminalsSelfDelivery2Response, error) {
	α := struct {
		OperationGetTerminalsSelfDelivery2 `xml:"tns:getTerminalsSelfDelivery2"`
	}{
		OperationGetTerminalsSelfDelivery2{
			GetTerminalsSelfDelivery2,
		},
	}

	γ := struct {
		OperationGetTerminalsSelfDelivery2Response `xml:"getTerminalsSelfDelivery2Response"`
	}{}
	if err := p.cli.RoundTripWithAction("GetTerminalsSelfDelivery2", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetTerminalsSelfDelivery2Response, nil
}
