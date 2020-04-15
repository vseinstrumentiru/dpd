package dpd

import (
	"github.com/fiorix/wsdl2go/soap"
)

const (
	trackingByDPDOrderNamespace = "http://dpd.ru/ws/tracing/2012-06-25"

	//CityCODPaymentLimit represents default C.O.D. payment limit for all dpd cities in which C.O.D. is available
	CityCODPaymentLimit = 250000
)

//Client for DPD soap api
type Client struct {
	auth     *Auth
	services *soapClients
}

//ServiceUrls DPD services urls
type ServiceUrls struct {
	CalculatorURL string
	GeographyURL  string
	OrderURL      string
	TrackingURL   string
}

type soapClients struct {
	geography  *soap.Client
	order      *soap.Client
	calculator *soap.Client
	tracking   *soap.Client
}

//NewClient creates new client
func NewClient(clientNumber int64, clientKey string, urls ServiceUrls) *Client {
	return &Client{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			geography: &soap.Client{
				Namespace: geographyNamespace,
				URL:       urls.GeographyURL,
			},
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       urls.OrderURL,
			},
			calculator: &soap.Client{
				Namespace: calculatorNamespace,
				URL:       urls.CalculatorURL,
			},
			tracking: &soap.Client{
				Namespace: trackingNamespace,
				URL:       urls.TrackingURL,
			},
		},
	}
}

//GetParcelShops return list of pickup/delivery points with restrictions about weight and dimensions
func (cl *Client) GetParcelShops(req *ParcelShopRequest) ([]*ParcelShop, error) {
	req.Auth = cl.auth
	req.Namespace = ""

	a := struct {
		operationGetParcelShops `xml:"tns:getParcelShops"`
	}{
		operationGetParcelShops{
			&getParcelShops{
				req,
				geographyNamespace,
			},
		},
	}

	y := struct {
		operationGetParcelShopsResponse `xml:"getParcelShopsResponse"`
	}{}

	if err := cl.services.geography.RoundTripWithAction("GetParcelShops", a, &y); err != nil {
		return nil, err
	}

	return y.operationGetParcelShopsResponse.GetParcelShopsResponse.Return.ParcelShop, nil
}

//GetTerminalsSelfDelivery2 return terminals list. They don't have any restrictions
func (cl *Client) GetTerminalsSelfDelivery2() ([]*Terminal, error) {
	auth := cl.auth
	auth.Namespace = new(string)

	a := struct {
		operationGetTerminalsSelfDelivery2 `xml:"tns:getTerminalsSelfDelivery2"`
	}{
		operationGetTerminalsSelfDelivery2{
			&getTerminalsSelfDelivery2Request{
				Auth:      auth,
				Namespace: geographyNamespace,
			},
		},
	}

	y := struct {
		operationGetTerminalsSelfDelivery2Response `xml:"getTerminalsSelfDelivery2Response"`
	}{}

	if err := cl.services.geography.RoundTripWithAction("getTerminalsSelfDelivery2Request", a, &y); err != nil {
		return nil, err
	}

	return y.Response.Return.Terminal, nil
}

//GetCitiesCashPay countryCode must comply with ISO 3166-1 alpha-2
func (cl *Client) GetCitiesCashPay(countryCode string) ([]*City, error) {
	a := struct {
		operationGetCitiesCashPay `xml:"tns:getCitiesCashPay"`
	}{
		operationGetCitiesCashPay{
			&getCitiesCashPay{
				Namespace: geographyNamespace,
				Request: &citiesCashPayRequest{
					Auth:        cl.auth,
					CountryCode: &countryCode,
				},
			},
		},
	}

	y := struct {
		operationGetCitiesCashPayResponse `xml:"getCitiesCashPayResponse"`
	}{}

	if err := cl.services.geography.RoundTripWithAction("GetCitiesCashPay", a, &y); err != nil {
		return nil, err
	}

	return y.GetCitiesCashPayResponse.Return, nil
}

//GetServiceCost2 return cost of delivery for Russia and Customs Union
func (cl *Client) GetServiceCost2(req *CalculateRequest) ([]*ServiceCostResponse, error) {
	req.Auth = cl.auth
	a := struct {
		operationGetServiceCost2 `xml:"tns:getServiceCost2"`
	}{
		operationGetServiceCost2{
			&getServiceCostRequest{
				Namespace: calculatorNamespace,
				Request:   req,
			},
		},
	}

	y := struct {
		operationGetServiceCost2Response `xml:"getServiceCost2Response"`
	}{}

	if err := cl.services.calculator.RoundTripWithAction("GetServiceCost2", a, &y); err != nil {
		return nil, err
	}

	return y.GetServiceCostResponse.Return, nil
}

//CreateOrder create a delivery order
func (cl *Client) CreateOrder(req *CreateOrderRequest) ([]*OrderStatus, error) {
	req.Auth = cl.auth
	a := struct {
		operationCreateOrder `xml:"tns:createOrder"`
	}{
		operationCreateOrder{
			&createOrder{
				Namespace: orderNamespace,
				Orders:    req,
			},
		},
	}

	y := struct {
		operationCreateOrderResponse `xml:"createOrderResponse"`
	}{}

	if err := cl.services.order.RoundTripWithAction("CreateOrder", a, &y); err != nil {
		return nil, err
	}

	return y.CreateOrderResponse.Return, nil
}

//AddParcels change order with parcel addition
func (cl *Client) AddParcels(req *UpdateOrderRequest) (*OrderCorrectionStatus, error) {
	req.Auth = cl.auth

	a := struct {
		operationAddParcels `xml:"tns:addParcels"`
	}{
		operationAddParcels{
			&addParcels{
				req,
			},
		},
	}

	y := struct {
		operationAddParcelsResponse `xml:"addParcelsResponse"`
	}{}

	if err := cl.services.order.RoundTripWithAction("AddParcels", a, &y); err != nil {
		return nil, err
	}

	return y.AddParcelsResponse.Return, nil
}

//RemoveParcels change order with parcel deletion
func (cl *Client) RemoveParcels(req *UpdateOrderRequest) (*OrderCorrectionStatus, error) {
	req.Auth = cl.auth
	a := struct {
		operationRemoveParcels `xml:"tns:removeParcels"`
	}{
		operationRemoveParcels{
			&removeParcels{
				req,
			},
		},
	}

	y := struct {
		operationRemoveParcelsResponse `xml:"removeParcelsResponse"`
	}{}

	if err := cl.services.order.RoundTripWithAction("RemoveParcels", a, &y); err != nil {
		return nil, err
	}

	return y.RemoveParcelsResponse.Return, nil
}

//CancelOrder cancel a delivery order
func (cl *Client) CancelOrder(req *CancelOrderRequest) ([]*OrderStatus, error) {
	req.Auth = cl.auth
	a := struct {
		operationCancelOrder `xml:"tns:cancelOrder"`
	}{
		operationCancelOrder{
			&cancelOrder{
				req,
			},
		},
	}

	y := struct {
		operationCancelOrderResponse `xml:"cancelOrderResponse"`
	}{}
	if err := cl.services.order.RoundTripWithAction("CancelOrder", a, &y); err != nil {
		return nil, err
	}

	return y.CancelOrderResponse.Return, nil
}

//GetOrderStatus returns order creation status
func (cl *Client) GetOrderStatus(req []*InternalOrderNumber) ([]*OrderStatus, error) {
	a := struct {
		operationGetOrderStatus `xml:"tns:getOrderStatus"`
	}{
		operationGetOrderStatus{
			&getOrderStatus{
				Namespace: orderNamespace,
				OrderStatus: &getOrderStatusRequest{
					Auth:  cl.auth,
					Order: req,
				},
			},
		},
	}

	y := struct {
		operationGetOrderStatusResponse `xml:"getOrderStatusResponse"`
	}{}
	if err := cl.services.order.RoundTripWithAction("GetOrderStatus", a, &y); err != nil {
		return nil, err
	}
	return y.GetOrderStatusResponse.Return, nil
}

//GetStatesByClient returns all statuses of client parcels since previous method call
func (cl *Client) GetStatesByClient() (*ParcelsStates, error) {
	a := struct {
		operationGetStatesByClient `xml:"tns:getStatesByClient"`
	}{
		operationGetStatesByClient{
			&getStatesByClient{
				&getStatesByClientRequest{
					Auth: cl.auth,
				},
			},
		},
	}

	y := struct {
		operationGetStatesByClientResponse `xml:"getStatesByClientResponse"`
	}{}

	if err := cl.services.tracking.RoundTripWithAction("GetStatesByClient", a, &y); err != nil {
		return nil, err
	}

	return y.GetStatesByClientResponse.Return, nil
}

//GetStatesByClientOrder returns list of states for all parcels for specified order
//Order identify by client side order number
func (cl *Client) GetStatesByClientOrder(req *TrackByClientOrderRequest) (*ParcelsStates, error) {
	req.Auth = cl.auth
	a := struct {
		operationGetStatesByClientOrder `xml:"tns:getStatesByClientOrder"`
	}{
		operationGetStatesByClientOrder{
			&getStatesByClientOrder{
				req,
			},
		},
	}

	y := struct {
		operationGetStatesByClientOrderResponse `xml:"getStatesByClientOrderResponse"`
	}{}

	if err := cl.services.tracking.RoundTripWithAction("GetStatesByClientOrder", a, &y); err != nil {
		return nil, err
	}

	return y.GetStatesByClientOrderResponse.Return, nil
}

//GetStatesByDPDOrder returns states history for specified parcel
//Parcel identify by client side parcel number
func (cl *Client) GetStatesByDPDOrder(req *TrackByDPDOrderRequest) (*ParcelsStates, error) {
	req.Auth = cl.auth
	a := struct {
		operationGetStatesByDPDOrder `xml:"tns:getStatesByDPDOrder"`
	}{
		operationGetStatesByDPDOrder{
			&getStatesByDPDOrder{
				Namespace: trackingByDPDOrderNamespace,
				Request:   req,
			},
		},
	}

	y := struct {
		operationGetStatesByDPDOrderResponse `xml:"getStatesByDPDOrderResponse"`
	}{}
	if err := cl.services.tracking.RoundTripWithAction("GetStatesByDPDOrder", a, &y); err != nil {
		return nil, err
	}

	return y.GetStatesByDPDOrderResponse.Return, nil
}
