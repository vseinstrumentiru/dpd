package dpd

const calculatorNamespace = "http://dpd.ru/ws/Calculator/2012-03-20"

type operationGetServiceCost2 struct {
	GetServiceCost2 *getServiceCostRequest `xml:"getServiceCost2,omitempty"`
}

type getServiceCostRequest struct {
	Namespace string            `xml:"xmlns,attr"`
	Request   *CalculateRequest `xml:"request,omitempty"`
}

//CalculateRequest GetServiceCost2 request body
type CalculateRequest struct {
	Namespace     string       `xml:"xmlns,attr"`
	Auth          *Auth        `xml:"auth,omitempty"`
	Pickup        *CityRequest `xml:"pickup,omitempty"`
	Delivery      *CityRequest `xml:"delivery,omitempty"`
	SelfPickup    *bool        `xml:"selfPickup,omitempty"`
	SelfDelivery  *bool        `xml:"selfDelivery,omitempty"`
	Weight        *float64     `xml:"weight,omitempty"`
	Volume        *float64     `xml:"volume,omitempty"`
	ServiceCode   *string      `xml:"serviceCode,omitempty"`
	PickupDate    *string      `xml:"pickupDate,omitempty"`
	MaxDays       *int         `xml:"maxDays,omitempty"`
	MaxCost       *float64     `xml:"maxCost,omitempty"`
	DeclaredValue *float64     `xml:"declaredValue,omitempty"`
}

//CityRequest Pickup and Delivery part of CalculateRequest
type CityRequest struct {
	CityID      *int64  `xml:"cityId,omitempty"`
	Index       *string `xml:"index,omitempty"`
	CityName    *string `xml:"cityName,omitempty"`
	RegionCode  *int    `xml:"regionCode,omitempty"`
	CountryCode *string `xml:"countryCode,omitempty"`
}

type getServiceCostResponse struct {
	Return []*ServiceCostResponse `xml:"return,omitempty"`
}

//ServiceCostResponse GetServiceCost2 return body
type ServiceCostResponse struct {
	ServiceCode *string  `xml:"serviceCode,omitempty"`
	ServiceName *string  `xml:"serviceName,omitempty"`
	Cost        *float64 `xml:"cost,omitempty"`
	Days        *int     `xml:"days,omitempty"`
}

type operationGetServiceCost2Response struct {
	GetServiceCostResponse *getServiceCostResponse `xml:"getServiceCost2Response,omitempty"`
}
