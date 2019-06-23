package dpd

import (
	"reflect"
	"testing"

	"github.com/fiorix/wsdl2go/soap"
	dpdSoap "github.com/vseinstrumentiru/dpd/soap"
)

func ExampleNewDPDClient() {
	NewDPDClient(
		DPDAuth{
			ClientKey:    "key",
			ClientNumber: 1000000,
		},
		DPDUrls{
			CalculatorUrl: "",
			OrderUrl:      "",
			GeographyUrl:  "",
			TrackingUrl:   "",
		},
	)
}

func ExampleDPDClient_GetCitiesCashPay() {
	client := NewDPDClient(DPDAuth{}, DPDUrls{})

	client.GetCitiesCashPay("RU")
}

func ExampleDPDClient_GetServiceCost2() {
	client := NewDPDClient(DPDAuth{}, DPDUrls{})

	req := NewCalculateRequest(
		NewCity(123).SetCityName("Анапа"),
		NewCity(321).SetCityName("Апана"),
		1.89,
		false,
		false)

	client.GetServiceCost2(req)
}

func ExampleDPDClient_GetParcelShops() {
	client := NewDPDClient(DPDAuth{}, DPDUrls{})

	req := NewParcelShopRequest()

	//For retrieving all points
	req.SetCountryCode("RU")

	//For retrieving points in specific city
	req.SetCountryCode("RU").
		SetCityCode("1243124") //DPD native, city identifier

	client.GetParcelShops(req)
}

var auth = DPDAuth{
	ClientNumber: 123,
	ClientKey:    "client_key",
}

var urls = DPDUrls{
	CalculatorUrl: "url",
	GeographyUrl:  "url",
	OrderUrl:      "url",
	TrackingUrl:   "url",
}

func TestNewDPDClient(t *testing.T) {
	type args struct {
		auth DPDAuth
		urls DPDUrls
	}
	tests := []struct {
		name string
		args args
		want *DPDClient
	}{
		{
			"Client with Auth and Urls",
			args{
				auth,
				urls,
			},

			&DPDClient{
				auth,
				urls,
				&services{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDPDClient(tt.args.auth, tt.args.urls); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("NewDPDClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDPDClient_getGeographyService(t *testing.T) {
	dpdClient := NewDPDClient(auth, urls)
	soapClient := soap.Client{
		Namespace: dpdSoap.GeographyNamespace,
		URL:       urls.GeographyUrl,
	}

	want := dpdSoap.NewDPDGeography2(&soapClient)

	tests := []struct {
		name string
		cl   *DPDClient
		want dpdSoap.DPDGeography2
	}{
		{
			"Get geography service",
			dpdClient,
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cl.getGeographyService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DPDClient.getGeographyService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDPDClient_getOrderService(t *testing.T) {
	dpdClient := NewDPDClient(auth, urls)
	soapClient := soap.Client{
		Namespace: dpdSoap.OrderNamespace,
		URL:       urls.OrderUrl,
	}

	want := dpdSoap.NewDPDOrder(&soapClient)

	tests := []struct {
		name string
		cl   *DPDClient
		want dpdSoap.DPDOrder
	}{
		{
			"Get order service",
			dpdClient,
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cl.getOrderService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DPDClient.getOrderService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDPDClient_getCalculatorService(t *testing.T) {
	dpdClient := NewDPDClient(auth, urls)
	soapClient := soap.Client{
		Namespace: dpdSoap.CalculatorNamespace,
		URL:       urls.CalculatorUrl,
	}

	want := dpdSoap.NewDPDCalculator(&soapClient)

	tests := []struct {
		name string
		cl   *DPDClient
		want dpdSoap.DPDCalculator
	}{
		{
			"Get calculator service",
			dpdClient,
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cl.getCalculatorService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DPDClient.getCalculatorService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDPDClient_getTrackingService(t *testing.T) {
	dpdClient := NewDPDClient(auth, urls)
	soapClient := soap.Client{
		Namespace: dpdSoap.TrackingNamespace,
		URL:       urls.TrackingUrl,
	}

	want := dpdSoap.NewParcelTracing(&soapClient)

	tests := []struct {
		name string
		cl   *DPDClient
		want dpdSoap.ParcelTracing
	}{
		{
			"Get tracking service",
			dpdClient,
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cl.getTrackingService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DPDClient.getTrackingService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDPDClient_getAuth(t *testing.T) {
	dpdClient := NewDPDClient(auth, urls)

	tests := []struct {
		name string
		cl   *DPDClient
		want *dpdSoap.Auth
	}{
		{
			"Get auth",
			dpdClient,
			&dpdSoap.Auth{
				ClientNumber: &auth.ClientNumber,
				ClientKey:    &auth.ClientKey,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cl.getAuth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DPDClient.getAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
