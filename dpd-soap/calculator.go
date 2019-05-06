package dpd_soap

import "github.com/fiorix/wsdl2go/soap"

const CalculatorNamespace = "http://dpd.ru/ws/calculator/2012-03-20"

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

type ServiceCostFault struct {
	Code         *string           `xml:"code,omitempty"`
	DeliveryDups []*CalculatorCity `xml:"deliveryDups,omitempty"`
	Message      *string           `xml:"message,omitempty"`
	PickupDups   []*CalculatorCity `xml:"pickupDups,omitempty"`
}

type ServiceCostFault2 struct {
	Code         *string      `xml:"code,omitempty"`
	DeliveryDups []*CityIndex `xml:"deliveryDups,omitempty"`
	Message      *string      `xml:"message,omitempty"`
	PickupDups   []*CityIndex `xml:"pickupDups,omitempty"`
}

type CalculatorCity struct {
	CityId      *int64  `xml:"cityId,omitempty"`
	CountryCode *string `xml:"countryCode,omitempty"`
	CountryName *string `xml:"countryName,omitempty"`
	RegionCode  *int    `xml:"regionCode,omitempty"`
	RegionName  *string `xml:"regionName,omitempty"`
	CityName    *string `xml:"cityName,omitempty"`
}

type CityIndex struct {
	CityId       *int64  `xml:"cityId,omitempty"`
	CountryCode  *string `xml:"countryCode,omitempty"`
	CountryName  *string `xml:"countryName,omitempty"`
	RegionCode   *int    `xml:"regionCode,omitempty"`
	RegionName   *string `xml:"regionName,omitempty"`
	CityCode     *string `xml:"cityCode,omitempty"`
	CityName     *string `xml:"cityName,omitempty"`
	Abbreviation *string `xml:"abbreviation,omitempty"`
	IndexMin     *string `xml:"indexMin,omitempty"`
	IndexMax     *string `xml:"indexMax,omitempty"`
}

type CityInternationalRequest struct {
	CountryName *string `xml:"countryName,omitempty"`
	CityName    *string `xml:"cityName,omitempty"`
	CityId      *int64  `xml:"cityId,omitempty"`
}

type CityRequest struct {
	CityId      *int64  `xml:"cityId,omitempty"`
	Index       *string `xml:"index,omitempty"`
	CityName    *string `xml:"cityName,omitempty"`
	RegionCode  *int    `xml:"regionCode,omitempty"`
	CountryCode *string `xml:"countryCode,omitempty"`
}

type GetServiceCost struct {
	Request *ServiceCostRequest `xml:"request,omitempty"`
}

type GetServiceCost2 struct {
	Ns      string              `xml:"xmlns,attr"`
	Request *ServiceCostRequest `xml:"request,omitempty"`
}

type GetServiceCost2Response struct {
	Return []*ServiceCost `xml:"return,omitempty"`
}

type GetServiceCostByParcels struct {
	Request *ServiceCostParcelsRequest `xml:"request,omitempty"`
}

type GetServiceCostByParcels2 struct {
	Request *ServiceCostParcelsRequest `xml:"request,omitempty"`
}

type GetServiceCostByParcels2Response struct {
	Return []*ServiceCost `xml:"return,omitempty"`
}

type GetServiceCostByParcelsResponse struct {
	Return []*ServiceCost `xml:"return,omitempty"`
}

type GetServiceCostInternational struct {
	Request *ServiceCostInternationalRequest `xml:"request,omitempty"`
}

type GetServiceCostInternationalResponse struct {
	Return []*ServiceCostInternational `xml:"return,omitempty"`
}

type GetServiceCostResponse struct {
	Return []*ServiceCost `xml:"return,omitempty"`
}

type ParcelRequest struct {
	Weight   *float64 `xml:"weight,omitempty"`
	Length   *float64 `xml:"length,omitempty"`
	Width    *float64 `xml:"width,omitempty"`
	Height   *float64 `xml:"height,omitempty"`
	Quantity *int     `xml:"quantity,omitempty"`
}

type ServiceCost struct {
	ServiceCode *string  `xml:"serviceCode,omitempty"`
	ServiceName *string  `xml:"serviceName,omitempty"`
	Cost        *float64 `xml:"cost,omitempty"`
	Days        *int     `xml:"days,omitempty"`
}

type ServiceCostInternational struct {
	ServiceCode *string  `xml:"serviceCode,omitempty"`
	ServiceName *string  `xml:"serviceName,omitempty"`
	Days        *string  `xml:"days,omitempty"`
	Cost        *float64 `xml:"cost,omitempty"`
	CostPin     *float64 `xml:"costPin,omitempty"`
	Weight      *float64 `xml:"weight,omitempty"`
	Volume      *float64 `xml:"Volume,omitempty"`
}

type ServiceCostInternationalRequest struct {
	Auth          *Auth                     `xml:"auth,omitempty"`
	Pickup        *CityInternationalRequest `xml:"pickup,omitempty"`
	Delivery      *CityInternationalRequest `xml:"delivery,omitempty"`
	SelfPickup    *bool                     `xml:"SelfPickup,omitempty"`
	SelfDelivery  *bool                     `xml:"SelfDelivery,omitempty"`
	Weight        *float64                  `xml:"weight,omitempty"`
	Length        *int64                    `xml:"length,omitempty"`
	Width         *int64                    `xml:"width,omitempty"`
	Height        *int64                    `xml:"height,omitempty"`
	DeclaredValue *float64                  `xml:"declaredValue,omitempty"`
	Insurance     *bool                     `xml:"insurance,omitempty"`
}

type ServiceCostParcelsRequest struct {
	Auth          *Auth            `xml:"auth,omitempty"`
	Pickup        *CityRequest     `xml:"pickup,omitempty"`
	Delivery      *CityRequest     `xml:"delivery,omitempty"`
	SelfPickup    *bool            `xml:"SelfPickup,omitempty"`
	SelfDelivery  *bool            `xml:"SelfDelivery,omitempty"`
	ServiceCode   *string          `xml:"serviceCode,omitempty"`
	PickupDate    *Date            `xml:"pickupDate,omitempty"`
	MaxDays       *int             `xml:"MaxDays,omitempty"`
	MaxCost       *float64         `xml:"MaxCost,omitempty"`
	DeclaredValue *float64         `xml:"declaredValue,omitempty"`
	Parcel        []*ParcelRequest `xml:"parcel,omitempty"`
}

type ServiceCostRequest struct {
	Ns            string       `xml:"xmlns,attr"`
	Auth          *Auth        `xml:"auth,omitempty"`
	Pickup        *CityRequest `xml:"pickup,omitempty"`
	Delivery      *CityRequest `xml:"delivery,omitempty"`
	SelfPickup    *bool        `xml:"selfPickup,omitempty"`
	SelfDelivery  *bool        `xml:"selfDelivery,omitempty"`
	Weight        *float64     `xml:"weight,omitempty"`
	Volume        *float64     `xml:"volume,omitempty"`
	ServiceCode   *string      `xml:"serviceCode,omitempty"`
	PickupDate    *Date        `xml:"pickupDate,omitempty"`
	MaxDays       *int         `xml:"maxDays,omitempty"`
	MaxCost       *float64     `xml:"maxCost,omitempty"`
	DeclaredValue *float64     `xml:"declaredValue,omitempty"`
}

type OperationGetServiceCost struct {
	GetServiceCost *GetServiceCost `xml:"getServiceCost,omitempty"`
}

type OperationGetServiceCostResponse struct {
	GetServiceCostResponse *GetServiceCostResponse `xml:"getServiceCostResponse,omitempty"`
}

type OperationGetServiceCost2 struct {
	GetServiceCost2 *GetServiceCost2 `xml:"getServiceCost2,omitempty"`
}

type OperationGetServiceCost2Response struct {
	GetServiceCost2Response *GetServiceCost2Response `xml:"getServiceCost2Response,omitempty"`
}

type OperationGetServiceCostByParcels struct {
	GetServiceCostByParcels *GetServiceCostByParcels `xml:"getServiceCostByParcels,omitempty"`
}

type OperationGetServiceCostByParcelsResponse struct {
	GetServiceCostByParcelsResponse *GetServiceCostByParcelsResponse `xml:"getServiceCostByParcelsResponse,omitempty"`
}

type OperationGetServiceCostByParcels2 struct {
	GetServiceCostByParcels2 *GetServiceCostByParcels2 `xml:"getServiceCostByParcels2,omitempty"`
}

type OperationGetServiceCostByParcels2Response struct {
	GetServiceCostByParcels2Response *GetServiceCostByParcels2Response `xml:"getServiceCostByParcels2Response,omitempty"`
}

type OperationGetServiceCostInternational struct {
	GetServiceCostInternational *GetServiceCostInternational `xml:"getServiceCostInternational,omitempty"`
}

type OperationGetServiceCostInternationalResponse struct {
	GetServiceCostInternationalResponse *GetServiceCostInternationalResponse `xml:"getServiceCostInternationalResponse,omitempty"`
}
