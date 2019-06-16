package dpd

import (
	dpdSoap "git.vseinstrumenti.net/golang/dpd/soap"
	"github.com/fiorix/wsdl2go/soap"
)

//Client to call DPD SOAP api methods
//
//Names of functions equals original DPD SOAP methods names
type DPDClient struct {
	auth     DPDAuth
	urls     DPDUrls
	services *services
}

//DPD authorization
type DPDAuth struct {
	ClientNumber int64
	ClientKey    string
}

//URLs of DPD services
type DPDUrls struct {
	CalculatorUrl string
	GeographyUrl  string
	OrderUrl      string
	TrackingUrl   string
}

type services struct {
	geography  dpdSoap.DPDGeography2
	order      dpdSoap.DPDOrder
	calculator dpdSoap.DPDCalculator
	tracking   dpdSoap.ParcelTracing
}

//Creates new client
func NewDPDClient(auth DPDAuth, urls DPDUrls) *DPDClient {
	return &DPDClient{
		auth:     auth,
		urls:     urls,
		services: &services{},
	}
}

func (cl *DPDClient) getGeographyService() dpdSoap.DPDGeography2 {
	if cl.services.geography == nil {
		client := soap.Client{
			Namespace: dpdSoap.GeographyNamespace,
			URL:       cl.urls.GeographyUrl,
		}

		cl.services.geography = dpdSoap.NewDPDGeography2(&client)
	}

	return cl.services.geography
}

func (cl *DPDClient) getOrderService() dpdSoap.DPDOrder {
	if cl.services.order == nil {
		client := soap.Client{
			Namespace: dpdSoap.OrderNamespace,
			URL:       cl.urls.OrderUrl,
		}

		cl.services.order = dpdSoap.NewDPDOrder(&client)
	}

	return cl.services.order
}

func (cl *DPDClient) getCalculatorService() dpdSoap.DPDCalculator {
	if cl.services.calculator == nil {
		client := soap.Client{
			Namespace: dpdSoap.CalculatorNamespace,
			URL:       cl.urls.CalculatorUrl,
		}

		cl.services.calculator = dpdSoap.NewDPDCalculator(&client)
	}

	return cl.services.calculator
}

func (cl *DPDClient) getTrackingService() dpdSoap.ParcelTracing {
	if cl.services.tracking == nil {
		client := soap.Client{
			Namespace: dpdSoap.TrackingNamespace,
			URL:       cl.urls.TrackingUrl,
		}

		cl.services.tracking = dpdSoap.NewParcelTracing(&client)
	}

	return cl.services.tracking
}

//Get list of pickup/delivery points with restrictions about weight and dimensions
func (cl *DPDClient) GetParcelShops(r *ParcelShopRequest) ([]*dpdSoap.ParcelShop, error) {
	req := r.toDPDRequest()
	req.Auth = cl.getAuth()
	req.Ns = dpdSoap.GeographyNamespace

	result, err := cl.getGeographyService().GetParcelShops(&dpdSoap.GetParcelShops{
		Request: req,
		Ns:      "",
	})

	if err != nil {
		return nil, err
	}

	return result.Return.ParcelShop, nil
}

//Get terminals list. They don't have any restrictions
func (cl *DPDClient) GetTerminalsSelfDelivery2() ([]*dpdSoap.TerminalSelf, error) {
	result, err := cl.getGeographyService().GetTerminalsSelfDelivery2(&dpdSoap.GetTerminalsSelfDelivery2{
		Ns:   dpdSoap.GeographyNamespace,
		Auth: cl.getAuth(),
	})

	if err != nil {
		return nil, err
	}

	return result.Return.Terminal, nil
}

//Set country code according ISO 3166-1 alpha-2 standard
//
//https://www.iso.org/obp/ui/#search
func (cl *DPDClient) GetCitiesCashPay(countryCode string) ([]*dpdSoap.City, error) {
	result, err := cl.getGeographyService().GetCitiesCashPay(&dpdSoap.GetCitiesCashPay{
		NS: dpdSoap.GeographyNamespace,
		Request: &dpdSoap.DpdCitiesCashPayRequest{
			Auth:        cl.getAuth(),
			CountryCode: &countryCode,
		},
	})

	if err != nil {
		return nil, err
	}

	return result.Return, nil
}

//Calculate cost of delivery for Russia and Customs Union
func (cl *DPDClient) GetServiceCost2(r *CalculateRequest) ([]*dpdSoap.ServiceCost, error) {
	req := r.toDpdRequest()
	req.Auth = cl.getAuth()

	result, err := cl.getCalculatorService().GetServiceCost2(&dpdSoap.GetServiceCost2{
		Ns:      dpdSoap.CalculatorNamespace,
		Request: req,
	})

	if err != nil {
		return nil, err
	}

	return result.Return, nil
}

func (cl *DPDClient) CreateOrder(r *CreateOrderRequest) ([]*dpdSoap.DpdOrderStatus, error) {
	req := r.toDPDRequest()
	req.Auth = cl.getAuth()

	result, err := cl.getOrderService().CreateOrder(&dpdSoap.CreateOrder{
		Orders: req,
	})

	if err != nil {
		return nil, err
	}

	return result.Return, nil
}

//Change order. Add parcels to order
func (cl *DPDClient) AddParcels(r *UpdateOrderRequest) (*dpdSoap.DpdOrderCorrectionStatus, error) {
	req := r.toDPDRequest()
	r.Auth = cl.getAuth()

	result, err := cl.getOrderService().AddParcels(&dpdSoap.AddParcels{
		Parcels: req,
	})

	if err != nil {
		return nil, err
	}

	return result.Return, nil
}

//Change order. Remove parcels from order
func (cl *DPDClient) RemoveParcels(r *UpdateOrderRequest) (*dpdSoap.DpdOrderCorrectionStatus, error) {
	req := r.toDPDRequest()
	r.Auth = cl.getAuth()

	result, err := cl.getOrderService().RemoveParcels(&dpdSoap.RemoveParcels{
		Parcels: req,
	})

	if err != nil {
		return nil, err
	}

	return result.Return, nil
}

func (cl *DPDClient) CancelOrder(r *CancelOrderRequest) ([]*dpdSoap.DpdOrderStatus, error) {
	req := r.toDPDRequest()
	r.Auth = cl.getAuth()

	result, err := cl.getOrderService().CancelOrder(&dpdSoap.CancelOrder{
		Orders: req,
	})

	if err != nil {
		return nil, err
	}

	return result.Return, nil
}

//Return all parcels statuses that changed after last request
func (cl *DPDClient) GetStatesByClient() (*dpdSoap.StateParcels, error) {
	result, err := cl.getTrackingService().GetStatesByClient(&dpdSoap.GetStatesByClient{
		Request: &dpdSoap.RequestClient{
			Auth: cl.getAuth(),
		},
	})

	if err != nil {
		return nil, err
	}

	return result.Return, err
}

//Get list of states for all parcels for specified order
//Order identify by client side order number
func (cl *DPDClient) GetStatesByClientOrder(r *ClientOrderRequest) (*dpdSoap.StateParcels, error) {
	req := r.toDPDRequest()
	req.Auth = cl.getAuth()

	result, err := cl.getTrackingService().GetStatesByClientOrder(&dpdSoap.GetStatesByClientOrder{
		Request: req,
	})

	if err != nil {
		return nil, err
	}

	return result.Return, nil
}

//Get states history for specified parcel
//Parcel identify by client side parcel number
func (cl *DPDClient) GetStatesByDPDOrder(r *DpdOrderRequest) (*dpdSoap.StateParcels, error) {
	req := r.toDPDRequest()
	req.Auth = cl.getAuth()

	result, err := cl.getTrackingService().GetStatesByDPDOrder(&dpdSoap.GetStatesByDPDOrder{
		Request: req,
	})

	if err != nil {
		return nil, err
	}

	return result.Return, nil
}

func (cl *DPDClient) getAuth() *dpdSoap.Auth {
	return &dpdSoap.Auth{
		ClientNumber: &cl.auth.ClientNumber,
		ClientKey:    &cl.auth.ClientKey,
	}
}
