package dpd

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/fiorix/wsdl2go/soap"
)

func ExampleNewClient() {
	NewClient(anyNumber, anyKey, ServiceUrls{
		CalculatorURL: "",
		OrderURL:      "",
		GeographyURL:  "",
		TrackingURL:   "",
	})
}

func ExampleClient_GetCitiesCashPay() {
	client := NewClient(0, "", ServiceUrls{})

	cities, err := client.GetCitiesCashPay("RU")
	if err != nil {
		log.Fatalf("Error occurred %s", err.Error())
	}

	for _, value := range cities {
		fmt.Println(*value.CityName)
	}
}

func ExampleClient_GetServiceCost2() {
	client := NewClient(0, "", ServiceUrls{})

	req := NewCalculateRequest(
		NewCity(123).SetCityName("Анапа"),
		NewCity(321).SetCityName("Апана"),
		1.89,
		false,
		false)

	offers, err := client.GetServiceCost2(req)
	if err != nil {
		log.Fatalf("Error occurred %s", err.Error())
	}

	for _, value := range offers {
		fmt.Printf("Delivery cost: %b", *value.Cost)
	}
}

func ExampleClient_GetParcelShops() {
	client := NewClient(0, "", ServiceUrls{})

	req := NewParcelShopRequest()

	//For retrieving all points
	req.SetCountryCode("RU")

	//For retrieving points in specific city
	req.SetCountryCode("RU").
		SetCityCode("1243124") //DPD native, city identifier

	shops, err := client.GetParcelShops(req)
	if err != nil {
		log.Fatalf("Error occurred %s", err.Error())
	}

	for _, value := range shops {
		fmt.Printf("Shop: %s", *value.Brand)
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		clientNumber int64
		clientKey    string
		urls         ServiceUrls
	}

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			"Client constructor",
			args{
				clientNumber,
				clientKey,
				ServiceUrls{
					"", "", "", "",
				},
			},
			&Client{
				&Auth{
					ClientNumber: &clientNumber,
					ClientKey:    &clientKey,
				},
				&soapClients{
					geography: &soap.Client{
						Namespace: geographyNamespace,
						URL:       "",
					},
					calculator: &soap.Client{
						Namespace: calculatorNamespace,
						URL:       "",
					},
					order: &soap.Client{
						Namespace: orderNamespace,
						URL:       "",
					},
					tracking: &soap.Client{
						Namespace: trackingNamespace,
						URL:       "",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.clientNumber, tt.args.clientKey, tt.args.urls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CreateOrder(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}
	type args struct {
		req *CreateOrderRequest
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getCreateOrderSuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       errServer.URL,
			},
		},
	}

	datePickup := time.Now().Format("2006-01-02")

	senderName := "Туда-сюда контора"
	senderCountry := thisCountry
	senderRegion := "Московская обл."
	senderCity := "Котельники"
	senderStreet := "Ленина"
	senderStreetAbbr := "ул."
	senderHouse := "3"
	senderContactFio := "Туда-Сюда Человек"
	senderContactPhone := "+790000000000"
	pickupTimePeriod := "9-18"

	orderNumberInternal := internalOrderNumber
	serviceCode := "PCL"
	serviceVariant := "ДД"
	cargoNumPack := 1
	cargoWeight := 5.3
	cargoRegistered := false
	cargoCategory := "Категория"

	receiverName := "Инокентий"
	receiverCountryName := thisCountry
	receiverRegion := "Нижугородская обл."
	receiverCity := "Арзамас"
	receiverStreet := "Ленина"
	receiverStreetAbbr := "ул"
	receiverHouse := "8"
	receiverFlat := "5"
	receiverContactFio := "Инокентий Мутный"
	receiverContactPhone := "+79888888888"

	orderNum := dpdOrderNumber
	status := "OK"

	arg := args{&CreateOrderRequest{
		Header: &Header{
			DatePickup: &datePickup,
			SenderAddress: &Address{
				Name:         &senderName,
				CountryName:  &senderCountry,
				Region:       &senderRegion,
				City:         &senderCity,
				Street:       &senderStreet,
				StreetAbbr:   &senderStreetAbbr,
				House:        &senderHouse,
				ContactFio:   &senderContactFio,
				ContactPhone: &senderContactPhone,
			},
			PickupTimePeriod: &pickupTimePeriod,
		},
		Order: []*Order{
			{
				OrderNumberInternal: &orderNumberInternal,
				ServiceCode:         &serviceCode,
				ServiceVariant:      &serviceVariant,
				CargoNumPack:        &cargoNumPack,
				CargoWeight:         &cargoWeight,
				CargoRegistered:     &cargoRegistered,
				CargoCategory:       &cargoCategory,
				ReturnAddress: &Address{
					Name:         &receiverName,
					CountryName:  &receiverCountryName,
					Region:       &receiverRegion,
					City:         &receiverCity,
					Street:       &receiverStreet,
					StreetAbbr:   &receiverStreetAbbr,
					House:        &receiverHouse,
					Flat:         &receiverFlat,
					ContactFio:   &receiverContactFio,
					ContactPhone: &receiverContactPhone,
				},
			},
		},
	},
	}

	want := []*OrderStatus{
		{
			OrderNumberInternal: &orderNumberInternal,
			OrderNum:            &orderNum,
			Status:              &status,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*OrderStatus
		wantErr bool
	}{
		{
			"Success",
			okFields,
			arg,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			arg,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.CreateOrder(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_AddParcels(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}
	type args struct {
		req *UpdateOrderRequest
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getAddParcelsSuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       errServer.URL,
			},
		},
	}

	orderNum := dpdOrderNumber
	orderNumberInternal := internalOrderNumber
	cargoNumPack := 2
	cargoWeight := 1.1
	cargoVolume := 0.05
	cargoCategory := "category"
	statusOK := "OK"

	arg := args{
		&UpdateOrderRequest{
			Auth: &Auth{
				ClientNumber: &clientNumber,
				ClientKey:    &clientKey,
			},
			OrderNum:            &orderNum,
			OrderNumberInternal: &orderNumberInternal,
			CargoNumPack:        &cargoNumPack,
			CargoWeight:         &cargoWeight,
			CargoVolume:         &cargoVolume,
			CargoCategory:       &cargoCategory,
		},
	}

	want := &OrderCorrectionStatus{
		OrderNum: &orderNum,
		Status:   &statusOK,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrderCorrectionStatus
		wantErr bool
	}{
		{
			"Success",
			okFields,
			arg,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			arg,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.AddParcels(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.AddParcels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.AddParcels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RemoveParcels(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}
	type args struct {
		req *UpdateOrderRequest
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getRemoveParcelsSuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       errServer.URL,
			},
		},
	}

	dpdOrderNum := dpdOrderNumber
	orderNumberinternal := internalOrderNumber
	cargoNumPack := 1
	cargoWeight := 1.1
	cargoVolume := 0.05
	cargoCategory := "category"
	parcelNum := "1"
	statusOK := "OK"

	arg := args{
		&UpdateOrderRequest{
			Auth: &Auth{
				ClientNumber: &clientNumber,
				ClientKey:    &clientKey,
			},
			OrderNum:            &dpdOrderNum,
			OrderNumberInternal: &orderNumberinternal,
			CargoNumPack:        &cargoNumPack,
			CargoWeight:         &cargoWeight,
			CargoVolume:         &cargoVolume,
			CargoCategory:       &cargoCategory,
			Parcel: []*Parcel{
				{
					Number: &parcelNum,
				},
			},
		},
	}

	want := &OrderCorrectionStatus{
		OrderNum: &dpdOrderNum,
		Status:   &statusOK,
		ParcelStatus: []*ParcelStatus{
			{
				Number: &parcelNum,
				Status: &statusOK,
			},
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OrderCorrectionStatus
		wantErr bool
	}{
		{
			"Success",
			okFields,
			arg,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			arg,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.RemoveParcels(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.RemoveParcels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.RemoveParcels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_CancelOrder(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}
	type args struct {
		req *CancelOrderRequest
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getCancelOrderSuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       errServer.URL,
			},
		},
	}

	orderNum := dpdOrderNumber
	pickupDate := string("2019-07-15")

	orderNumberInternal := internalOrderNumber
	status := "Canceled"

	arg := args{
		&CancelOrderRequest{
			Cancel: []*OrderToCancel{
				{
					OrderNum:   &orderNum,
					PickupDate: &pickupDate,
				},
			},
		},
	}

	want :=
		[]*OrderStatus{
			{
				OrderNumberInternal: &orderNumberInternal,
				OrderNum:            &orderNum,
				Status:              &status,
			},
		}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*OrderStatus
		wantErr bool
	}{
		{
			"Success",
			okFields,
			arg,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			arg,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.CancelOrder(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CancelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetOrderStatus(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}
	type args struct {
		req []*InternalOrderNumber
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getOrderStatusSuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			order: &soap.Client{
				Namespace: orderNamespace,
				URL:       errServer.URL,
			},
		},
	}

	orderNumberInternal := internalOrderNumber

	arg := args{
		[]*InternalOrderNumber{
			{
				OrderNumberInternal: &orderNumberInternal,
			},
		},
	}

	orderNum := dpdOrderNumber
	status := "OrderCancelled"

	want := []*OrderStatus{
		{
			OrderNumberInternal: &orderNumberInternal,
			OrderNum:            &orderNum,
			Status:              &status,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*OrderStatus
		wantErr bool
	}{
		{
			"Success",
			okFields,
			arg,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			arg,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.GetOrderStatus(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetOrderStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetOrderStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetServiceCost2(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}
	type args struct {
		req *CalculateRequest
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getServiceCostSuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			calculator: &soap.Client{
				Namespace: calculatorNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			calculator: &soap.Client{
				Namespace: calculatorNamespace,
				URL:       errServer.URL,
			},
		},
	}

	a, b := true, false
	weight, volume := 1.1, 0.05
	cityID := int64(21321421)

	arg := args{
		&CalculateRequest{
			Namespace: calculatorNamespace,
			Pickup: &CityRequest{
				CityID: &cityID,
			},
			Delivery: &CityRequest{
				CityID: &cityID,
			},
			SelfPickup:   &b,
			SelfDelivery: &a,
			Weight:       &weight,
			Volume:       &volume,
		},
	}

	serviceCode1, serviceCode2 := "BZP", "CUR"
	serviceName1, serviceName2 := "DPD 18:00", "DPD CLASSIC domestic"
	cost1, cost2 := 7310.71, 6946.2
	days1, days2 := 13, 9

	want := []*ServiceCostResponse{
		{
			ServiceCode: &serviceCode1,
			ServiceName: &serviceName1,
			Cost:        &cost1,
			Days:        &days1,
		},
		{
			ServiceCode: &serviceCode2,
			ServiceName: &serviceName2,
			Cost:        &cost2,
			Days:        &days2,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ServiceCostResponse
		wantErr bool
	}{
		{
			"Success",
			okFields,
			arg,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			arg,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.GetServiceCost2(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetServiceCost2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetServiceCost2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTerminalsSelfDelivery2(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getTerminalsSelfDelivery2SuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			geography: &soap.Client{
				Namespace: geographyNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			geography: &soap.Client{
				Namespace: geographyNamespace,
				URL:       errServer.URL,
			},
		},
	}

	terminalCode := "M11"
	terminalName := "Москва -  M11 Илимская"

	cityID := int64(49694102)
	countryCode := "RU"
	regionCode := "77"
	regionName := "Москва"
	cityCode := "77000000000"
	cityName := regionName
	index := "129301"
	street := "Касаткина"
	streetAbbr := "улица"
	houseNo := "11"
	structure := "3"

	address := GeographyAddress{
		CityID:      &cityID,
		CountryCode: &countryCode,
		RegionCode:  &regionCode,
		RegionName:  &regionName,
		CityCode:    &cityCode,
		CityName:    &cityName,
		Index:       &index,
		Street:      &street,
		StreetAbbr:  &streetAbbr,
		HouseNo:     &houseNo,
		Structure:   &structure,
	}

	latitude := 55.826581
	longitude := 37.660667

	geo := GeoCoordinates{
		Latitude:  &latitude,
		Longitude: &longitude,
	}

	schedulesOperations := [3]string{
		"Payment",
		"PaymentByBankCard",
		"SelfPickup",
	}

	weekDays1 := "Пн,Вт,Ср,Чт,Пт"
	weekDays2 := "Сб,Вс"

	workTime := "10:00 - 21:00"

	schedule := make([]*Schedule, 0)

	for i := 0; i < 3; i++ {
		operation := schedulesOperations[i]

		schedule = append(schedule, &Schedule{
			Operation: &operation,
			Timetable: []*Timetable{
				{
					WeekDays: &weekDays1,
					WorkTime: &workTime,
				},
				{
					WeekDays: &weekDays2,
					WorkTime: &workTime,
				},
			},
		})
	}

	extraService := "ТРМ"
	serviceCode := "NDY"

	want := []*Terminal{{
		TerminalCode:   &terminalCode,
		TerminalName:   &terminalName,
		Address:        &address,
		GeoCoordinates: &geo,
		Schedule:       schedule,
		ExtraService: []*ExtraServiceParameters{
			{
				EsCode: &extraService,
			},
		},
		Services: []*string{&serviceCode},
	}}

	tests := []struct {
		name    string
		fields  fields
		want    []*Terminal
		wantErr bool
	}{
		{
			"Success",
			okFields,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.GetTerminalsSelfDelivery2()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.getTerminalsSelfDelivery2Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.getTerminalsSelfDelivery2Request() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetParcelShops(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}
	type args struct {
		req *ParcelShopRequest
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getParcelShopsSuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			geography: &soap.Client{
				Namespace: geographyNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			geography: &soap.Client{
				Namespace: geographyNamespace,
				URL:       errServer.URL,
			},
		},
	}

	countryCode := "RU"

	arg := args{
		&ParcelShopRequest{
			CountryCode: &countryCode,
			Namespace:   geographyNamespace,
		},
	}

	want := []*ParcelShop{
		getParcelShopResponse(),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ParcelShop
		wantErr bool
	}{
		{
			"Success",
			okFields,
			arg,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			arg,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.GetParcelShops(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetParcelShops() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetParcelShops() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetCitiesCashPay(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}
	type args struct {
		countryCode string
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getCitiesCashPaySuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			geography: &soap.Client{
				Namespace: geographyNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			geography: &soap.Client{
				Namespace: geographyNamespace,
				URL:       errServer.URL,
			},
		},
	}

	countryCode := "RU"

	arg := args{
		countryCode,
	}

	cityID1, cityID2 := int64(48951627), int64(48994107)
	countryName := thisCountry
	regionCode1, regionCode2 := 42, 66
	regionName1, regionName2 := "Кемеровская", "Свердловская"
	cityCode1, cityCode2 := "42000009000", "66000001000"
	cityName1, cityName2 := "Кемерово", "Екатеринбург"
	abbrevation := "г"
	indexMin1, indexMax1, indexMin2, indexMax2 := "650000", "650993", "620000", "624130"

	want := []*City{
		{
			CityID:       &cityID1,
			CountryCode:  &countryCode,
			CountryName:  &countryName,
			RegionCode:   &regionCode1,
			RegionName:   &regionName1,
			CityCode:     &cityCode1,
			CityName:     &cityName1,
			Abbreviation: &abbrevation,
			IndexMin:     &indexMin1,
			IndexMax:     &indexMax1,
		},
		{
			CityID:       &cityID2,
			CountryCode:  &countryCode,
			CountryName:  &countryName,
			RegionCode:   &regionCode2,
			RegionName:   &regionName2,
			CityCode:     &cityCode2,
			CityName:     &cityName2,
			Abbreviation: &abbrevation,
			IndexMin:     &indexMin2,
			IndexMax:     &indexMax2,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*City
		wantErr bool
	}{
		{
			"Success",
			okFields,
			arg,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			arg,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.GetCitiesCashPay(tt.args.countryCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetCitiesCashPay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetCitiesCashPay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetStatesByClient(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getStatesByClientSuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			tracking: &soap.Client{
				Namespace: trackingNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			tracking: &soap.Client{
				Namespace: trackingNamespace,
				URL:       errServer.URL,
			},
		},
	}

	docId := int64(17543659022)
	docDate := "2019-07-18T15:59:46"
	clientNumberResp := int64(1001027554)
	resultComplete := true

	orderNumberDPD := dpdOrderNumber
	orderNumberInternal := internalOrderNumber
	parcelNumberDPD := "271227267"
	pickupDate := "2019-07-18"
	planDeliveryDate := "2019-07-19"
	newState := "OnRoad"
	transitionTime := "2019-07-18T08:52:45"
	terminalCode := "M16"
	terminalCity := "Москва"

	want := &ParcelsStates{
		DocID:          &docId,
		DocDate:        &docDate,
		ClientNumber:   &clientNumberResp,
		ResultComplete: &resultComplete,
		States: []*ParcelState{
			{
				ClientOrderNr:    &orderNumberInternal,
				DpdOrderNr:       &orderNumberDPD,
				DpdParcelNr:      &parcelNumberDPD,
				PickupDate:       &pickupDate,
				PlanDeliveryDate: &planDeliveryDate,
				NewState:         &newState,
				TransitionTime:   &transitionTime,
				TerminalCode:     &terminalCode,
				TerminalCity:     &terminalCity,
			},
		},
	}

	tests := []struct {
		name    string
		fields  fields
		want    *ParcelsStates
		wantErr bool
	}{
		{
			"Success",
			okFields,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.GetStatesByClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetStatesByClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetStatesByClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetStatesByClientOrder(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}
	type args struct {
		req *TrackByClientOrderRequest
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getStatesByClientOrderSuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey
	orderNumberInternal := internalOrderNumber

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			tracking: &soap.Client{
				Namespace: trackingNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			tracking: &soap.Client{
				Namespace: trackingNamespace,
				URL:       errServer.URL,
			},
		},
	}

	arg := args{
		&TrackByClientOrderRequest{
			ClientOrderNr: &orderNumberInternal,
		},
	}

	docId := int64(17543035562)
	docDate := "2019-07-18T11:16:46"
	clientNumberResp := int64(1001027554)
	resultComplete := true

	orderNumberDPD := dpdOrderNumber
	parcelNumberDPD := "271227267"
	pickupDate := "2019-07-18"
	planDeliveryDate := "2019-07-19"
	newState := "OnTerminalPickup"
	transitionTime := "2019-07-18T08:51:21"
	terminalCode := "M16"
	terminalCity := "Москва"

	want := &ParcelsStates{
		DocID:          &docId,
		DocDate:        &docDate,
		ClientNumber:   &clientNumberResp,
		ResultComplete: &resultComplete,
		States: []*ParcelState{
			{
				ClientOrderNr:    &orderNumberInternal,
				DpdOrderNr:       &orderNumberDPD,
				DpdParcelNr:      &parcelNumberDPD,
				PickupDate:       &pickupDate,
				PlanDeliveryDate: &planDeliveryDate,
				NewState:         &newState,
				TransitionTime:   &transitionTime,
				TerminalCode:     &terminalCode,
				TerminalCity:     &terminalCity,
			},
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ParcelsStates
		wantErr bool
	}{
		{
			"Sucess",
			okFields,
			arg,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			args{&TrackByClientOrderRequest{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.GetStatesByClientOrder(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetStatesByClientOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetStatesByClientOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetStatesByDPDOrder(t *testing.T) {
	type fields struct {
		auth     *Auth
		services *soapClients
	}
	type args struct {
		req *TrackByDPDOrderRequest
	}

	okServer := makeTestHTTPServer(t, http.StatusOK, getStatesByDPDOrderSuccessResponse())
	defer okServer.Close()

	errServer := makeTestHTTPServer(t, http.StatusInternalServerError, []byte(""))
	defer errServer.Close()

	clientNumber := int64(anyNumber)
	clientKey := anyKey
	orderNumberDPD := dpdOrderNumber

	okFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			tracking: &soap.Client{
				Namespace: trackingNamespace,
				URL:       okServer.URL,
			},
		},
	}

	errFields := fields{
		auth: &Auth{
			ClientNumber: &clientNumber,
			ClientKey:    &clientKey,
		},
		services: &soapClients{
			tracking: &soap.Client{
				Namespace: trackingNamespace,
				URL:       errServer.URL,
			},
		},
	}

	arg := args{
		&TrackByDPDOrderRequest{
			DpdOrderNr: &orderNumberDPD,
		},
	}

	docId := int64(17543580110)
	docDate := "2019-07-18T15:24:32"
	clientNumberResp := int64(1001027554)
	resultComplete := true

	orderNumberInternal := internalOrderNumber
	parcelNumberDPD := "271227267"
	pickupDate := "2019-07-18"
	planDeliveryDate := "2019-07-19"
	newState := "Delivered"
	transitionTime := "2019-07-18T08:58:00"
	terminalCode := "GOJ"
	terminalCity := "Нижний Новгород"
	consignee := "Инокентий Мутный"

	want := &ParcelsStates{
		DocID:          &docId,
		DocDate:        &docDate,
		ClientNumber:   &clientNumberResp,
		ResultComplete: &resultComplete,
		States: []*ParcelState{
			{
				ClientOrderNr:    &orderNumberInternal,
				DpdOrderNr:       &orderNumberDPD,
				DpdParcelNr:      &parcelNumberDPD,
				PickupDate:       &pickupDate,
				PlanDeliveryDate: &planDeliveryDate,
				NewState:         &newState,
				TransitionTime:   &transitionTime,
				TerminalCode:     &terminalCode,
				TerminalCity:     &terminalCity,
				Consignee:        &consignee,
			},
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ParcelsStates
		wantErr bool
	}{
		{
			"Success",
			okFields,
			arg,
			want,
			false,
		},
		{
			"Fail",
			errFields,
			arg,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &Client{
				auth:     tt.fields.auth,
				services: tt.fields.services,
			}
			got, err := cl.GetStatesByDPDOrder(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetStatesByDPDOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetStatesByDPDOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

/***************************************************************
**********************        Staff        *********************
****************************************************************/
func getServiceCostSuccessResponse() []byte {
	return []byte(`
			<?xml version='1.0' encoding='UTF-8'?>
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
				<S:Body>
					<ns2:getServiceCost2Response xmlns:ns2="http://dpd.ru/ws/calculator/2012-03-20">
						<return>
							<serviceCode>BZP</serviceCode>
							<serviceName>DPD 18:00</serviceName>
							<cost>7310.71</cost>
							<days>13</days>
						</return>
						<return>
							<serviceCode>CUR</serviceCode>
							<serviceName>DPD CLASSIC domestic</serviceName>
							<cost>6946.2</cost>
							<days>9</days>
						</return>
					</ns2:getServiceCost2Response>
				</S:Body>
			</S:Envelope>
`)
}

func getTerminalsSelfDelivery2SuccessResponse() []byte {
	return []byte(`
			<?xml version='1.0' encoding='UTF-8'?>
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
				<S:Body>
					<ns2:getTerminalsSelfDelivery2Response xmlns:ns2="http://dpd.ru/ws/geography/2015-05-20">
						<return>
							<terminal>
								<terminalCode>M11</terminalCode>
								<terminalName>Москва -  M11 Илимская</terminalName>
								<address>
									<cityId>49694102</cityId>
									<countryCode>RU</countryCode>
									<regionCode>77</regionCode>
									<regionName>Москва</regionName>
									<cityCode>77000000000</cityCode>
									<cityName>Москва</cityName>
									<index>129301</index>
									<street>Касаткина</street>
									<streetAbbr>улица</streetAbbr>
									<houseNo>11</houseNo>
									<structure>3</structure>
								</address>
								<geoCoordinates>
									<latitude>55.826581</latitude>
									<longitude>37.660667</longitude>
								</geoCoordinates>
								<schedule>
									<operation>Payment</operation>
									<timetable>
										<weekDays>Пн,Вт,Ср,Чт,Пт</weekDays>
										<workTime>10:00 - 21:00</workTime>
									</timetable>
									<timetable>
										<weekDays>Сб,Вс</weekDays>
										<workTime>10:00 - 21:00</workTime>
									</timetable>
								</schedule>
								<schedule>
									<operation>PaymentByBankCard</operation>
									<timetable>
										<weekDays>Пн,Вт,Ср,Чт,Пт</weekDays>
										<workTime>10:00 - 21:00</workTime>
									</timetable>
									<timetable>
										<weekDays>Сб,Вс</weekDays>
										<workTime>10:00 - 21:00</workTime>
									</timetable>
								</schedule>
								<schedule>
									<operation>SelfPickup</operation>
									<timetable>
										<weekDays>Пн,Вт,Ср,Чт,Пт</weekDays>
										<workTime>10:00 - 21:00</workTime>
									</timetable>
									<timetable>
										<weekDays>Сб,Вс</weekDays>
										<workTime>10:00 - 21:00</workTime>
									</timetable>
								</schedule>
								<extraService>
									<esCode>ТРМ</esCode>
								</extraService>
								<services>
									<serviceCode>NDY</serviceCode>
								</services>
							</terminal>
						</return>
					</ns2:getTerminalsSelfDelivery2Response>
				</S:Body>
			</S:Envelope>
		`)
}

func getCitiesCashPaySuccessResponse() []byte {
	return []byte(`
			<?xml version='1.0' encoding='UTF-8'?>
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
				<S:Body>
					<ns2:getCitiesCashPayResponse xmlns:ns2="http://dpd.ru/ws/geography/2015-05-20">
						<return>
							<cityId>48951627</cityId>
							<countryCode>RU</countryCode>
							<countryName>Россия</countryName>
							<regionCode>42</regionCode>
							<regionName>Кемеровская</regionName>
							<cityCode>42000009000</cityCode>
							<cityName>Кемерово</cityName>
							<abbreviation>г</abbreviation>
							<indexMin>650000</indexMin>
							<indexMax>650993</indexMax>
						</return>
						<return>
							<cityId>48994107</cityId>
							<countryCode>RU</countryCode>
							<countryName>Россия</countryName>
							<regionCode>66</regionCode>
							<regionName>Свердловская</regionName>
							<cityCode>66000001000</cityCode>
							<cityName>Екатеринбург</cityName>
							<abbreviation>г</abbreviation>
							<indexMin>620000</indexMin>
							<indexMax>624130</indexMax>
						</return>
					</ns2:getCitiesCashPayResponse>
				</S:Body>
			</S:Envelope>
		`)
}

func getParcelShopsSuccessResponse() []byte {
	return []byte(`
			<?xml version='1.0' encoding='UTF-8'?>
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
				<S:Body>
					<ns2:getParcelShopsResponse xmlns:ns2="http://dpd.ru/ws/geography/2015-05-20">
						<return>
							<parcelShop>
								<code>001G</code>
								<parcelShopType>ПВП</parcelShopType>
								<state>Open</state>
								<address>
									<cityId>49694167</cityId>
									<countryCode>RU</countryCode>
									<regionCode>78</regionCode>
									<regionName>Санкт-Петербург</regionName>
									<cityCode>78000000000</cityCode>
									<cityName>Санкт-Петербург</cityName>
									<index>196135</index>
									<street>Алтайская</street>
									<streetAbbr>ул</streetAbbr>
									<houseNo>16</houseNo>
									<descript>Выход из метро в сторону аэропорта, от Московской пл</descript>
								</address>
								<brand>Главпункт</brand>
								<clientDepartmentNum>Moskovskaya-A16</clientDepartmentNum>
								<geoCoordinates>
									<latitude>59.849681</latitude>
									<longitude>30.327345</longitude>
								</geoCoordinates>
								<limits>
									<maxShipmentWeight>70</maxShipmentWeight>
									<maxWeight>20</maxWeight>
									<maxLength>200</maxLength>
									<maxWidth>200</maxWidth>
									<maxHeight>200</maxHeight>
									<dimensionSum>200</dimensionSum>
								</limits>
								<schedule>
									<operation>Payment</operation>
									<timetable>
										<weekDays>Пн,Вт,Ср,Чт,Пт,Сб,Вс</weekDays>
										<workTime>10:00 - 21:00</workTime>
									</timetable>
								</schedule>
								<schedule>
									<operation>PaymentByBankCard</operation>
									<timetable>
										<weekDays>Пн,Вт,Ср,Чт,Пт,Сб,Вс</weekDays>
										<workTime>10:00 - 21:00</workTime>
									</timetable>
								</schedule>
								<schedule>
									<operation>SelfDelivery</operation>
									<timetable>
										<weekDays>Пн,Вт,Ср,Чт,Пт,Сб,Вс</weekDays>
										<workTime>10:00 - 21:00</workTime>
									</timetable>
								</schedule>
								<extraService>
									<esCode>НПП</esCode>
									<params>
										<name>sum_npp</name>
										<value>100000</value>
									</params>
								</extraService>
								<extraService>
									<esCode>ОЖД</esCode>
									<params>
										<name>reason_delay</name>
										<value>ПРОС</value>
									</params>
								</extraService>
								<extraService>
									<esCode>ТРМ</esCode>
								</extraService>
								<services>
									<serviceCode>BZP</serviceCode>
								</services>
							</parcelShop>
						</return>
					</ns2:getParcelShopsResponse>
				</S:Body>
			</S:Envelope>
		`)
}

func getParcelShopResponse() *ParcelShop {
	code := "001G"
	parcelShopType := "ПВП"
	state := "Open"

	cityID := int64(49694167)
	countryCode := "RU"
	regionCode := "78"
	regionName := "Санкт-Петербург"
	cityCode := "78000000000"
	cityName := regionName
	index := "196135"
	street := "Алтайская"
	streetAbbr := "ул"
	houseNo := "16"
	description := `Выход из метро в сторону аэропорта, от Московской пл`

	address := GeographyAddress{
		CityID:      &cityID,
		CountryCode: &countryCode,
		RegionCode:  &regionCode,
		RegionName:  &regionName,
		CityCode:    &cityCode,
		CityName:    &cityName,
		Index:       &index,
		Street:      &street,
		StreetAbbr:  &streetAbbr,
		HouseNo:     &houseNo,
		Description: &description,
	}

	brand := "Главпункт"
	clientDepartmentNum := "Moskovskaya-A16"

	latitude := 59.849681
	longitude := 30.327345

	geoCoordinates := GeoCoordinates{
		Latitude:  &latitude,
		Longitude: &longitude,
	}

	maxShipmentWeight := 70.0
	maxWeight := 20.0
	maxLength := 200.0
	maxWidth := 200.0
	maxHeight := 200.0
	dimensionsSum := 200.0

	limits := Limits{
		MaxShipmentWeight: &maxShipmentWeight,
		MaxWeight:         &maxWeight,
		MaxLength:         &maxLength,
		MaxWidth:          &maxWidth,
		MaxHeight:         &maxHeight,
		DimensionSum:      &dimensionsSum,
	}

	schedules := make([]*Schedule, 0)

	operations := [3]string{
		"Payment",
		"PaymentByBankCard",
		"SelfDelivery",
	}

	weekDays := "Пн,Вт,Ср,Чт,Пт,Сб,Вс"
	workTime := "10:00 - 21:00"

	for i := 0; i < 3; i++ {
		operation := operations[i]
		schedules = append(schedules, &Schedule{
			Operation: &operation,
			Timetable: []*Timetable{
				{
					WeekDays: &weekDays,
					WorkTime: &workTime,
				},
			},
		})
	}

	extraServices := make([]*ExtraServiceParameters, 0)

	services := [3]string{
		"НПП",
		"ОЖД",
		"ТРМ",
	}
	serviceParamNames := [2]string{
		"sum_npp",
		"reason_delay",
	}
	serviceParamValue := [2]string{
		"100000",
		"ПРОС",
	}

	for i := 0; i < 3; i++ {
		name, param := "", ""

		if len(serviceParamNames) > i && len(serviceParamValue) > i {
			name = serviceParamNames[i]
			param = serviceParamValue[i]
		}

		esCode := services[i]

		es := ExtraServiceParameters{
			EsCode: &esCode,
		}

		if name != "" && param != "" {
			es.Params = []*ExtraServiceParameter{
				{
					Name:  &name,
					Value: &param,
				},
			}
		}

		extraServices = append(extraServices, &es)
	}

	serviceCode := "BZP"

	return &ParcelShop{
		Code:                &code,
		ParcelShopType:      &parcelShopType,
		State:               &state,
		Address:             &address,
		Brand:               &brand,
		ClientDepartmentNum: &clientDepartmentNum,
		GeoCoordinates:      &geoCoordinates,
		Limits:              &limits,
		Schedule:            schedules,
		ExtraService:        extraServices,
		Services:            []*string{&serviceCode},
	}
}

func getCreateOrderSuccessResponse() []byte {
	return []byte(`
			<?xml version='1.0' encoding='UTF-8'?>
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
				<S:Body>
					<ns2:createOrderResponse xmlns:ns2="http://dpd.ru/ws/order2/2012-04-04">
						<return>
							<orderNumberInternal>1901-200100-42389</orderNumberInternal>
							<orderNum>RU019202656</orderNum>
							<status>OK</status>
						</return>
					</ns2:createOrderResponse>
				</S:Body>
			</S:Envelope>
	`)
}

func getAddParcelsSuccessResponse() []byte {
	return []byte(`
			<?xml version='1.0' encoding='UTF-8'?>
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
				<S:Body>
					<ns2:addParcelsResponse xmlns:ns2="http://dpd.ru/ws/order2/2012-04-04">
						<return>
							<orderNum>RU019202656</orderNum>
							<status>OK</status>
						</return>
					</ns2:addParcelsResponse>
				</S:Body>
			</S:Envelope>
		`)
}

func getRemoveParcelsSuccessResponse() []byte {
	return []byte(`
		<?xml version='1.0' encoding='UTF-8'?>
		<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
			<S:Body>
				<ns2:removeParcelsResponse xmlns:ns2="http://dpd.ru/ws/order2/2012-04-04">
					<return>
						<orderNum>RU019202656</orderNum>
						<status>OK</status>
						<parcelStatus>
							<number>1</number>
							<status>OK</status>
						</parcelStatus>
					</return>
				</ns2:removeParcelsResponse>
			</S:Body>
		</S:Envelope>
	`)
}

func getCancelOrderSuccessResponse() []byte {
	return []byte(`
			<?xml version='1.0' encoding='UTF-8'?>
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
				<S:Body>
					<ns2:cancelOrderResponse xmlns:ns2="http://dpd.ru/ws/order2/2012-04-04">
						<return>
							<orderNumberInternal>1901-200100-42389</orderNumberInternal>
							<orderNum>RU019202656</orderNum>
							<status>Canceled</status>
						</return>
					</ns2:cancelOrderResponse>
				</S:Body>
			</S:Envelope>
	`)
}

func getOrderStatusSuccessResponse() []byte {
	return []byte(`
			<?xml version='1.0' encoding='UTF-8'?>
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
				<S:Body>
					<ns2:getOrderStatusResponse xmlns:ns2="http://dpd.ru/ws/order2/2012-04-04">
						<return>
							<orderNumberInternal>1901-200100-42389</orderNumberInternal>
							<orderNum>RU019202656</orderNum>
							<status>OrderCancelled</status>
						</return>
					</ns2:getOrderStatusResponse>
				</S:Body>
			</S:Envelope>
	`)
}

func getStatesByClientOrderSuccessResponse() []byte {
	return []byte(`
			<?xml version='1.0' encoding='UTF-8'?>
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
				<S:Body>
					<ns2:getStatesByClientOrderResponse xmlns:ns2="http://dpd.ru/ws/tracing/2011-11-18">
						<return>
							<docId>17543035562</docId>
							<docDate>2019-07-18T11:16:46</docDate>
							<clientNumber>1001027554</clientNumber>
							<resultComplete>true</resultComplete>
							<states>
								<clientOrderNr>1901-200100-42389</clientOrderNr>
								<dpdOrderNr>RU019202656</dpdOrderNr>
								<dpdParcelNr>271227267</dpdParcelNr>
								<pickupDate>2019-07-18</pickupDate>
								<planDeliveryDate>2019-07-19</planDeliveryDate>
								<newState>OnTerminalPickup</newState>
								<transitionTime>2019-07-18T08:51:21</transitionTime>
								<terminalCode>M16</terminalCode>
								<terminalCity>Москва</terminalCity>
							</states>
						</return>
					</ns2:getStatesByClientOrderResponse>
				</S:Body>
			</S:Envelope>
	`)
}

func getStatesByDPDOrderSuccessResponse() []byte {
	return []byte(`
	<?xml version='1.0' encoding='UTF-8'?>
	<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
		<S:Body>
			<ns2:getStatesByDPDOrderResponse xmlns:ns2="http://dpd.ru/ws/tracing/2011-11-18">
				<return>
					<docId>17543580110</docId>
					<docDate>2019-07-18T15:24:32</docDate>
					<clientNumber>1001027554</clientNumber>
					<resultComplete>true</resultComplete>
					<states>
						<clientOrderNr>1901-200100-42389</clientOrderNr>
						<dpdOrderNr>RU019202656</dpdOrderNr>
						<dpdParcelNr>271227267</dpdParcelNr>
						<pickupDate>2019-07-18</pickupDate>
						<planDeliveryDate>2019-07-19</planDeliveryDate>
						<newState>Delivered</newState>
						<transitionTime>2019-07-18T08:58:00</transitionTime>
						<terminalCode>GOJ</terminalCode>
						<terminalCity>Нижний Новгород</terminalCity>
						<consignee>Инокентий Мутный</consignee>
					</states>
				</return>
			</ns2:getStatesByDPDOrderResponse>
		</S:Body>
	</S:Envelope>
		`)
}

func getStatesByClientSuccessResponse() []byte {
	return []byte(`
		<?xml version='1.0' encoding='UTF-8'?>
		<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
			<S:Body>
				<ns2:getStatesByClientResponse xmlns:ns2="http://dpd.ru/ws/tracing/2011-11-18">
					<return>
						<docId>17543659022</docId>
						<docDate>2019-07-18T15:59:46</docDate>
						<clientNumber>1001027554</clientNumber>
						<resultComplete>true</resultComplete>
						<states>
							<clientOrderNr>1901-200100-42389</clientOrderNr>
							<dpdOrderNr>RU019202656</dpdOrderNr>
							<dpdParcelNr>271227267</dpdParcelNr>
							<pickupDate>2019-07-18</pickupDate>
							<planDeliveryDate>2019-07-19</planDeliveryDate>
							<newState>OnRoad</newState>
							<transitionTime>2019-07-18T08:52:45</transitionTime>
							<terminalCode>M16</terminalCode>
							<terminalCity>Москва</terminalCity>
						</states>
					</return>
				</ns2:getStatesByClientResponse>
			</S:Body>
		</S:Envelope>
	`)
}
