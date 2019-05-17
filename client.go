package dpd_sdk

import (
	dpdSoap "git.vseinstrumenti.net/golang-sandbox/dpd-sdk/dpd-soap"
	"github.com/fiorix/wsdl2go/soap"
)

type DpdClient struct {
	auth        DPDAuth
	countryCode string
	urls        DPDUrls
	services    *services
}

type DPDAuth struct {
	ClientNumber int64
	ClientKey    string
}

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

func NewDPDClient(auth DPDAuth, urls DPDUrls, countryCode string) *DpdClient {
	return &DpdClient{
		auth:        auth,
		urls:        urls,
		countryCode: countryCode,
		services:    &services{},
	}
}

func (cl *DpdClient) getGeographyService() dpdSoap.DPDGeography2 {
	if cl.services.geography == nil {
		client := soap.Client{
			Namespace: dpdSoap.GeographyNamespace,
			URL:       cl.urls.GeographyUrl,
		}

		cl.services.geography = dpdSoap.NewDPDGeography2(&client)
	}

	return cl.services.geography
}

func (cl *DpdClient) getOrderService() dpdSoap.DPDOrder {
	if cl.services.order == nil {
		client := soap.Client{
			Namespace: dpdSoap.OrderNamespace,
			URL:       cl.urls.OrderUrl,
		}

		cl.services.order = dpdSoap.NewDPDOrder(&client)
	}

	return cl.services.order
}

func (cl *DpdClient) getCalculatorService() dpdSoap.DPDCalculator {
	if cl.services.calculator == nil {
		client := soap.Client{
			Namespace: dpdSoap.CalculatorNamespace,
			URL:       cl.urls.CalculatorUrl,
		}

		cl.services.calculator = dpdSoap.NewDPDCalculator(&client)
	}

	return cl.services.calculator
}

func (cl *DpdClient) getTrackingService() dpdSoap.ParcelTracing {
	if cl.services.tracking == nil {
		client := soap.Client{
			Namespace: dpdSoap.TrackingNamespace,
			URL:       cl.urls.TrackingUrl,
		}

		cl.services.tracking = dpdSoap.NewParcelTracing(&client)
	}

	return cl.services.tracking
}

func (cl *DpdClient) GetParcelShops(r *ParcelShopRequest) ([]*dpdSoap.ParcelShop, error) {
	req := r.toDPDRequest()
	req.Auth = cl.getAuth()
	req.Ns = dpdSoap.GeographyNamespace

	result, err := cl.getGeographyService().GetParcelShops(&dpdSoap.GetParcelShops{
		Request: req,
		Ns:      dpdSoap.GeographyNamespace,
	})

	if err != nil {
		return nil, err
	}

	return result.Return.ParcelShop, nil
}

func (cl *DpdClient) GetTerminalsSelfDelivery2() ([]*dpdSoap.TerminalSelf, error) {
	result, err := cl.getGeographyService().GetTerminalsSelfDelivery2(&dpdSoap.GetTerminalsSelfDelivery2{
		Ns:   dpdSoap.GeographyNamespace,
		Auth: cl.getAuth(),
	})

	if err != nil {
		return nil, err
	}

	return result.Return.Terminal, nil
}

func (cl *DpdClient) GetCitiesCashPay(countryCode string) ([]*dpdSoap.City, error) {
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

func (cl *DpdClient) GetServiceCost2(r *CalculateRequest) ([]*dpdSoap.ServiceCost, error) {
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

func (cl *DpdClient) CreateOrder(r *CreateOrderRequest) ([]*dpdSoap.DpdOrderStatus, error) {
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

func (cl *DpdClient) AddParcels(r *UpdateOrderRequest) (*dpdSoap.DpdOrderCorrectionStatus, error) {
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

func (cl *DpdClient) RemoveParcels(r *UpdateOrderRequest) (*dpdSoap.DpdOrderCorrectionStatus, error) {
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

func (cl *DpdClient) CancelOrder(r *CancelOrderRequest) ([]*dpdSoap.DpdOrderStatus, error) {
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

func (cl *DpdClient) GetStatesByClient() (*dpdSoap.StateParcels, error) {
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

func (cl *DpdClient) GetStatesByClientOrder(r *ClientOrderRequest) (*dpdSoap.StateParcels, error) {
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

func (cl *DpdClient) GetStatesByDPDOrder(r *DpdOrderRequest) (*dpdSoap.StateParcels, error) {
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

func (cl *DpdClient) getAuth() *dpdSoap.Auth {
	return &dpdSoap.Auth{
		ClientNumber: &cl.auth.ClientNumber,
		ClientKey:    &cl.auth.ClientKey,
	}
}
