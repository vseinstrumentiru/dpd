package dpd

import (
	"reflect"
	"testing"
	"time"

	dpdSoap "github.com/seacomandor/dpd/soap"
)

func createEmptyOrderRequest() *CreateOrderRequest {
	var orders []*dpdSoap.Order

	return &CreateOrderRequest{
		Header: &dpdSoap.Header{},
		Order:  orders,
	}
}

func TestNewCreateOrderRequest(t *testing.T) {
	tests := []struct {
		name string
		want *CreateOrderRequest
	}{
		{
			"Constructor",
			createEmptyOrderRequest(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCreateOrderRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateOrderRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateOrderRequest_SetDatePickup(t *testing.T) {
	type args struct {
		time time.Time
	}

	date := time.Now()
	dpdDate := dpdSoap.Date(date.Format("2006-01-02"))

	want := createEmptyOrderRequest()
	want.Header.DatePickup = &dpdDate

	tests := []struct {
		name string
		r    *CreateOrderRequest
		args args
		want *CreateOrderRequest
	}{
		{
			"Set date pickup",
			NewCreateOrderRequest(),
			args{
				date,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetDatePickup(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrderRequest.SetDatePickup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateOrderRequest_SetPayer(t *testing.T) {
	type args struct {
		payer int64
	}

	var payer int64 = 2321123

	want := createEmptyOrderRequest()
	want.Header.Payer = &payer

	tests := []struct {
		name string
		r    *CreateOrderRequest
		args args
		want *CreateOrderRequest
	}{
		{
			"Set payer",
			NewCreateOrderRequest(),
			args{
				payer,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetPayer(tt.args.payer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrderRequest.SetPayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateOrderRequest_SetSender(t *testing.T) {
	type args struct {
		address Address
	}

	address := NewAddress()
	dpdAddress := dpdSoap.Address(*address)

	want := createEmptyOrderRequest()
	want.Header.SenderAddress = &dpdAddress

	tests := []struct {
		name string
		r    *CreateOrderRequest
		args args
		want *CreateOrderRequest
	}{
		{
			"Set sender",
			NewCreateOrderRequest(),
			args{
				*address,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetSender(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrderRequest.SetSender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateOrderRequest_SetPickupTimePeriod(t *testing.T) {
	type args struct {
		period string
	}

	period := "11-18"

	want := createEmptyOrderRequest()
	want.Header.PickupTimePeriod = &period

	tests := []struct {
		name string
		r    *CreateOrderRequest
		args args
		want *CreateOrderRequest
	}{
		{
			"Set pikcup time period",
			NewCreateOrderRequest(),
			args{
				period,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetPickupTimePeriod(tt.args.period); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrderRequest.SetPickupTimePeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateOrderRequest_SetRegularNum(t *testing.T) {
	type args struct {
		num string
	}

	num := "any string"

	want := createEmptyOrderRequest()
	want.Header.RegularNum = &num

	tests := []struct {
		name string
		r    *CreateOrderRequest
		args args
		want *CreateOrderRequest
	}{
		{
			"Set regular num",
			NewCreateOrderRequest(),
			args{
				num,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetRegularNum(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrderRequest.SetRegularNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateOrderRequest_AddOrder(t *testing.T) {
	type args struct {
		order *Order
	}

	order := new(Order)

	want := createEmptyOrderRequest()
	want.Order = append(want.Order, &dpdSoap.Order{})

	tests := []struct {
		name string
		r    *CreateOrderRequest
		args args
		want *CreateOrderRequest
	}{
		{
			"Add order",
			NewCreateOrderRequest(),
			args{
				order,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.AddOrder(tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrderRequest.AddOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateOrderRequest_toDPDRequest(t *testing.T) {
	var orders []*dpdSoap.Order

	tests := []struct {
		name string
		r    *CreateOrderRequest
		want *dpdSoap.DpdOrdersData
	}{
		{
			"converter",
			NewCreateOrderRequest(),
			&dpdSoap.DpdOrdersData{
				Header: &dpdSoap.Header{},
				Order:  orders,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.toDPDRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateOrderRequest.toDPDRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createEmptyOrder() *Order {
	var parcel []*dpdSoap.Parcel
	var extraService []*dpdSoap.OrderExtraService
	var extraParameter []*dpdSoap.Parameter
	var unitLoad []*dpdSoap.UnitLoad

	return &Order{
		ExtraParam:   extraParameter,
		ExtraService: extraService,
		Parcel:       parcel,
		UnitLoad:     unitLoad,
	}
}

func TestNewOrder(t *testing.T) {
	tests := []struct {
		name string
		want *Order
	}{
		{
			"Constructor",
			createEmptyOrder(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetInternalOrderNumber(t *testing.T) {
	type args struct {
		number string
	}

	number := "any string"

	want := createEmptyOrder()
	want.OrderNumberInternal = &number

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set internal order number",
			NewOrder(),
			args{
				number,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetInternalOrderNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetInternalOrderNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetServiceCode(t *testing.T) {
	type args struct {
		code string
	}

	code := "any string"

	want := createEmptyOrder()
	want.ServiceCode = &code

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set service code",
			NewOrder(),
			args{
				code,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetServiceCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetServiceCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetServiceVariant(t *testing.T) {
	type args struct {
		variant string
	}

	var variant = "any string"

	want := createEmptyOrder()
	want.ServiceVariant = &variant

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set service variant",
			NewOrder(),
			args{
				variant,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetServiceVariant(tt.args.variant); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetServiceVariant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetCargoNumPack(t *testing.T) {
	type args struct {
		num int
	}

	num := 123

	want := createEmptyOrder()
	want.CargoNumPack = &num

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set cargo num pack",
			NewOrder(),
			args{
				num,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetCargoNumPack(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetCargoNumPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetCargoWeight(t *testing.T) {
	type args struct {
		weight float64
	}

	weight := 124.12

	want := createEmptyOrder()
	want.CargoWeight = &weight

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set cargo weight",
			NewOrder(),
			args{
				weight,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetCargoWeight(tt.args.weight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetCargoWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetCargoVolume(t *testing.T) {
	type args struct {
		volume float64
	}

	volume := 12.32

	want := createEmptyOrder()
	want.CargoVolume = &volume

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set cargo volume",
			NewOrder(),
			args{
				volume,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetCargoVolume(tt.args.volume); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetCargoVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetCargoRegistered(t *testing.T) {
	type args struct {
		flag bool
	}

	flag := false

	want := createEmptyOrder()
	want.CargoRegistered = &flag

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set cargo registered",
			NewOrder(),
			args{
				flag,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetCargoRegistered(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetCargoRegistered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetCargoValue(t *testing.T) {
	type args struct {
		value float64
	}

	value := 123.12

	want := createEmptyOrder()
	want.CargoValue = &value

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set cargo value",
			NewOrder(),
			args{
				value,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetCargoValue(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetCargoValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SerCargoCategory(t *testing.T) {
	type args struct {
		category string
	}

	category := "any string"

	want := createEmptyOrder()
	want.CargoCategory = &category

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set cargo category",
			NewOrder(),
			args{
				category,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SerCargoCategory(tt.args.category); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SerCargoCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetDeliveryTimePeriod(t *testing.T) {
	type args struct {
		period string
	}

	period := "any string"

	want := createEmptyOrder()
	want.DeliveryTimePeriod = &period

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set delivery time period",
			NewOrder(),
			args{
				period,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetDeliveryTimePeriod(tt.args.period); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetDeliveryTimePeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetPaymentType(t *testing.T) {
	type args struct {
		pType string
	}

	var pType = "any string"

	want := createEmptyOrder()
	want.PaymentType = &pType

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set payment type",
			NewOrder(),
			args{
				pType,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetPaymentType(tt.args.pType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetPaymentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetReceiverAddress(t *testing.T) {
	type args struct {
		address *Address
	}

	address := new(Address)

	want := createEmptyOrder()
	want.ReceiverAddress = &dpdSoap.Address{}

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set receiver address",
			NewOrder(),
			args{
				address,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetReceiverAddress(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetReceiverAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_SetReturnAddress(t *testing.T) {
	type args struct {
		address *Address
	}

	address := new(Address)

	want := createEmptyOrder()
	want.ReturnAddress = &dpdSoap.Address{}

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Set return address",
			NewOrder(),
			args{
				address,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetReturnAddress(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.SetReturnAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_AddExtraParameter(t *testing.T) {
	type args struct {
		parameter *ExtraParameter
	}

	parameter := new(ExtraParameter)

	want := createEmptyOrder()
	want.ExtraParam = append(want.ExtraParam, &dpdSoap.Parameter{})

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Add extra parameter",
			NewOrder(),
			args{
				parameter,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.AddExtraParameter(tt.args.parameter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.AddExtraParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_AddExtraService(t *testing.T) {
	type args struct {
		service *ExtraService
	}

	service := new(ExtraService)

	want := createEmptyOrder()
	want.ExtraService = append(want.ExtraService, &dpdSoap.OrderExtraService{})

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Add extra service",
			NewOrder(),
			args{
				service,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.AddExtraService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.AddExtraService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_AddParcel(t *testing.T) {
	type args struct {
		parcel *Parcel
	}

	parcel := new(Parcel)

	want := createEmptyOrder()
	want.Parcel = append(want.Parcel, &dpdSoap.Parcel{})

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Add parcel",
			NewOrder(),
			args{
				parcel,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.AddParcel(tt.args.parcel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.AddParcel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_AddUnitLoad(t *testing.T) {
	type args struct {
		load *UnitLoad
	}

	unit := new(UnitLoad)

	want := createEmptyOrder()
	want.UnitLoad = append(want.UnitLoad, &dpdSoap.UnitLoad{})

	tests := []struct {
		name string
		o    *Order
		args args
		want *Order
	}{
		{
			"Add unit load",
			NewOrder(),
			args{
				unit,
			},
			want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.AddUnitLoad(tt.args.load); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.AddUnitLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAddress(t *testing.T) {
	tests := []struct {
		name string
		want *Address
	}{
		{
			"Constructor",
			&Address{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetCode(t *testing.T) {
	type args struct {
		code string
	}

	code := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set code",
			NewAddress(),
			args{
				code,
			},
			&Address{
				Code: &code,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetName(t *testing.T) {
	type args struct {
		name string
	}

	name := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set name",
			NewAddress(),
			args{
				name,
			},
			&Address{
				Name: &name,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetTerminalCode(t *testing.T) {
	type args struct {
		code string
	}

	terminalCode := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set terminal code",
			NewAddress(),
			args{
				terminalCode,
			},
			&Address{
				TerminalCode: &terminalCode,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetTerminalCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetTerminalCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetAddressString(t *testing.T) {
	type args struct {
		address string
	}

	address := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set address string",
			NewAddress(),
			args{
				address,
			},
			&Address{
				AddressString: &address,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetAddressString(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetAddressString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetCountryName(t *testing.T) {
	type args struct {
		name string
	}

	name := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set country name",
			NewAddress(),
			args{
				name,
			},
			&Address{
				CountryName: &name,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetCountryName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetCountryName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetZip(t *testing.T) {
	type args struct {
		zip string
	}

	zip := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set zip",
			NewAddress(),
			args{
				zip,
			},
			&Address{
				Index: &zip,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetZip(tt.args.zip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetZip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetRegion(t *testing.T) {
	type args struct {
		region string
	}

	region := "Any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set region",
			NewAddress(),
			args{
				region,
			},
			&Address{
				Region: &region,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetRegion(tt.args.region); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetCity(t *testing.T) {
	type args struct {
		city string
	}

	city := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set city",
			NewAddress(),
			args{
				city,
			},
			&Address{
				City: &city,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetCity(tt.args.city); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetStreet(t *testing.T) {
	type args struct {
		street string
	}

	street := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set street",
			NewAddress(),
			args{
				street,
			},
			&Address{
				Street: &street,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetStreet(tt.args.street); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetStreet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetStreetAbbr(t *testing.T) {
	type args struct {
		abbr string
	}

	abbr := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set street abbreviation",
			NewAddress(),
			args{
				abbr,
			},
			&Address{
				StreetAbbr: &abbr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetStreetAbbr(tt.args.abbr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetStreetAbbr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetHouse(t *testing.T) {
	type args struct {
		house string
	}

	house := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set house",
			NewAddress(),
			args{
				house,
			},
			&Address{
				House: &house,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetHouse(tt.args.house); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetHousing(t *testing.T) {
	type args struct {
		housing string
	}

	housing := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set housing",
			NewAddress(),
			args{
				housing,
			},
			&Address{
				HouseKorpus: &housing,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetHousing(tt.args.housing); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetHousing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetBuilding(t *testing.T) {
	type args struct {
		building string
	}

	building := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set building",
			NewAddress(),
			args{
				building,
			},
			&Address{
				Str: &building,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetBuilding(tt.args.building); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetPossession(t *testing.T) {
	type args struct {
		possession string
	}

	possession := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set possession",
			NewAddress(),
			args{
				possession,
			},
			&Address{
				Vlad: &possession,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetPossession(tt.args.possession); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetPossession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetExtraInfo(t *testing.T) {
	type args struct {
		info string
	}

	info := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set extra info",
			NewAddress(),
			args{
				info,
			},
			&Address{
				ExtraInfo: &info,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetExtraInfo(tt.args.info); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetExtraInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetOffice(t *testing.T) {
	type args struct {
		office string
	}

	office := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set office",
			NewAddress(),
			args{
				office,
			},
			&Address{
				Office: &office,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetOffice(tt.args.office); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetOffice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetFlat(t *testing.T) {
	type args struct {
		flat string
	}

	flat := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set flat",
			NewAddress(),
			args{
				flat,
			},
			&Address{
				Flat: &flat,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetFlat(tt.args.flat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetFlat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetWorkTimeFrom(t *testing.T) {
	type args struct {
		start string
	}

	start := "10"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set work time from",
			NewAddress(),
			args{
				start,
			},
			&Address{
				WorkTimeFrom: &start,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetWorkTimeFrom(tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetWorkTimeFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetWorkTimeTo(t *testing.T) {
	type args struct {
		end string
	}

	end := "19"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set work time to",
			NewAddress(),
			args{
				end,
			},
			&Address{
				WorkTimeTo: &end,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetWorkTimeTo(tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetWorkTimeTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetDinnerTimeFrom(t *testing.T) {
	type args struct {
		start string
	}

	start := ""

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set dinner time from",
			NewAddress(),
			args{
				start,
			},
			&Address{
				DinnerTimeFrom: &start,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetDinnerTimeFrom(tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetDinnerTimeFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetDinnerTimeTo(t *testing.T) {
	type args struct {
		end string
	}

	end := "string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set dinner time to",
			NewAddress(),
			args{
				end,
			},
			&Address{
				DinnerTimeTo: &end,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetDinnerTimeTo(tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetDinnerTimeTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetContactFullName(t *testing.T) {
	type args struct {
		fullName string
	}

	fullName := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set contact full name",
			NewAddress(),
			args{
				fullName,
			},
			&Address{
				ContactFio: &fullName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetContactFullName(tt.args.fullName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetContactFullName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetContactPhone(t *testing.T) {
	type args struct {
		phone string
	}

	phone := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set contact phone",
			NewAddress(),
			args{
				phone,
			},
			&Address{
				ContactPhone: &phone,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetContactPhone(tt.args.phone); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetContactPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetContactEmail(t *testing.T) {
	type args struct {
		email string
	}

	email := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set contact email",
			NewAddress(),
			args{
				email,
			},
			&Address{
				ContactEmail: &email,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetContactEmail(tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetContactEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetInstructions(t *testing.T) {
	type args struct {
		instructions string
	}

	instructions := "any string"

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set instructions",
			NewAddress(),
			args{
				instructions,
			},
			&Address{
				Instructions: &instructions,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetInstructions(tt.args.instructions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_SetNeedPass(t *testing.T) {
	type args struct {
		flag bool
	}

	flag := false

	tests := []struct {
		name string
		a    *Address
		args args
		want *Address
	}{
		{
			"Set need pass",
			NewAddress(),
			args{
				flag,
			},
			&Address{
				NeedPass: &flag,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SetNeedPass(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Address.SetNeedPass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewParcel(t *testing.T) {
	tests := []struct {
		name string
		want *Parcel
	}{
		{
			"Constructor",
			&Parcel{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParcel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParcel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcel_SetNumber(t *testing.T) {
	type args struct {
		number string
	}

	number := "any string"

	tests := []struct {
		name string
		p    *Parcel
		args args
		want *Parcel
	}{
		{
			"Set number",
			NewParcel(),
			args{
				number,
			},
			&Parcel{
				Number: &number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SetNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parcel.SetNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcel_SetDPDParcelNumber(t *testing.T) {
	type args struct {
		number int64
	}

	var number int64 = 324987987

	tests := []struct {
		name string
		p    *Parcel
		args args
		want *Parcel
	}{
		{
			"Set DPD parcel number",
			NewParcel(),
			args{
				number,
			},
			&Parcel{
				DpdParcelNumber: &number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SetDPDParcelNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parcel.SetDPDParcelNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcel_SetNumberForPrint(t *testing.T) {
	type args struct {
		number string
	}

	number := "any string"

	tests := []struct {
		name string
		p    *Parcel
		args args
		want *Parcel
	}{
		{
			"Set number for print",
			NewParcel(),
			args{
				number,
			},
			&Parcel{
				NumberForPrint: &number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SetNumberForPrint(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parcel.SetNumberForPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcel_SetBoxNeeded(t *testing.T) {
	type args struct {
		needed int
	}

	needed := 2

	tests := []struct {
		name string
		p    *Parcel
		args args
		want *Parcel
	}{
		{
			"Set box needed",
			NewParcel(),
			args{
				needed,
			},
			&Parcel{
				BoxNeeded: &needed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SetBoxNeeded(tt.args.needed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parcel.SetBoxNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcel_SetWeight(t *testing.T) {
	type args struct {
		weight float64
	}

	weight := 1.23

	tests := []struct {
		name string
		p    *Parcel
		args args
		want *Parcel
	}{
		{
			"Set weight",
			NewParcel(),
			args{
				weight,
			},
			&Parcel{
				Weight: &weight,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SetWeight(tt.args.weight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parcel.SetWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcel_SetLength(t *testing.T) {
	type args struct {
		length float64
	}

	length := 1.34

	tests := []struct {
		name string
		p    *Parcel
		args args
		want *Parcel
	}{
		{
			"Set length",
			NewParcel(),
			args{
				length,
			},
			&Parcel{
				Length: &length,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SetLength(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parcel.SetLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcel_SetWidth(t *testing.T) {
	type args struct {
		width float64
	}

	width := 1.32

	tests := []struct {
		name string
		p    *Parcel
		args args
		want *Parcel
	}{
		{
			"Set width",
			NewParcel(),
			args{
				width,
			},
			&Parcel{
				Width: &width,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SetWidth(tt.args.width); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parcel.SetWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcel_SetHeight(t *testing.T) {
	type args struct {
		height float64
	}

	height := 1.23

	tests := []struct {
		name string
		p    *Parcel
		args args
		want *Parcel
	}{
		{
			"Set height",
			NewParcel(),
			args{
				height,
			},
			&Parcel{
				Height: &height,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SetHeight(tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parcel.SetHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcel_SetInsuranceCost(t *testing.T) {
	type args struct {
		cost float64
	}

	cost := 1000.00

	tests := []struct {
		name string
		p    *Parcel
		args args
		want *Parcel
	}{
		{
			"Set insurance cost",
			NewParcel(),
			args{
				cost,
			},
			&Parcel{
				InsuranceCost: &cost,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SetInsuranceCost(tt.args.cost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parcel.SetInsuranceCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcel_SetCodAmount(t *testing.T) {
	type args struct {
		amount float64
	}

	amount := 123.12

	tests := []struct {
		name string
		p    *Parcel
		args args
		want *Parcel
	}{
		{
			"Set cod amount",
			NewParcel(),
			args{
				amount,
			},
			&Parcel{
				CodAmount: &amount,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.SetCodAmount(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parcel.SetCodAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUpdateOrderRequest(t *testing.T) {
	tests := []struct {
		name string
		want *UpdateOrderRequest
	}{
		{
			"Constructor",
			&UpdateOrderRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUpdateOrderRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdateOrderRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrderRequest_SetDPDOrderNumber(t *testing.T) {
	type args struct {
		number string
	}

	number := "any string"

	tests := []struct {
		name string
		r    *UpdateOrderRequest
		args args
		want *UpdateOrderRequest
	}{
		{
			"Set DPD order number",
			NewUpdateOrderRequest(),
			args{
				number,
			},
			&UpdateOrderRequest{
				OrderNum: &number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetDPDOrderNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrderRequest.SetDPDOrderNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrderRequest_SetInternalOrderNumber(t *testing.T) {
	type args struct {
		number string
	}

	number := "any string"

	tests := []struct {
		name string
		r    *UpdateOrderRequest
		args args
		want *UpdateOrderRequest
	}{
		{
			"Set internal order number",
			NewUpdateOrderRequest(),
			args{
				number,
			},
			&UpdateOrderRequest{
				OrderNumberInternal: &number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetInternalOrderNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrderRequest.SetInternalOrderNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrderRequest_SetCargoNumPack(t *testing.T) {
	type args struct {
		num int
	}

	num := 123

	tests := []struct {
		name string
		r    *UpdateOrderRequest
		args args
		want *UpdateOrderRequest
	}{
		{
			"Set cargo num pack",
			NewUpdateOrderRequest(),
			args{
				num,
			},
			&UpdateOrderRequest{
				CargoNumPack: &num,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetCargoNumPack(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrderRequest.SetCargoNumPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrderRequest_SetCargoWeight(t *testing.T) {
	type args struct {
		weight float64
	}

	weight := 12.23

	tests := []struct {
		name string
		r    *UpdateOrderRequest
		args args
		want *UpdateOrderRequest
	}{
		{
			"Set cargo weight",
			NewUpdateOrderRequest(),
			args{
				weight,
			},
			&UpdateOrderRequest{
				CargoWeight: &weight,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetCargoWeight(tt.args.weight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrderRequest.SetCargoWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrderRequest_SetCargoVolume(t *testing.T) {
	type args struct {
		volume float64
	}

	volume := 12.2

	tests := []struct {
		name string
		r    *UpdateOrderRequest
		args args
		want *UpdateOrderRequest
	}{
		{
			"Set cargo volume",
			NewUpdateOrderRequest(),
			args{
				volume,
			},
			&UpdateOrderRequest{
				CargoVolume: &volume,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetCargoVolume(tt.args.volume); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrderRequest.SetCargoVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrderRequest_SetCargoValue(t *testing.T) {
	type args struct {
		value float64
	}

	value := 213.21

	tests := []struct {
		name string
		r    *UpdateOrderRequest
		args args
		want *UpdateOrderRequest
	}{
		{
			"Set cargo value",
			NewUpdateOrderRequest(),
			args{
				value,
			},
			&UpdateOrderRequest{
				CargoValue: &value,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetCargoValue(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrderRequest.SetCargoValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrderRequest_SetCargoCategory(t *testing.T) {
	type args struct {
		category string
	}

	category := "any string"

	tests := []struct {
		name string
		r    *UpdateOrderRequest
		args args
		want *UpdateOrderRequest
	}{
		{
			"Set cargo category",
			NewUpdateOrderRequest(),
			args{
				category,
			},
			&UpdateOrderRequest{
				CargoCategory: &category,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetCargoCategory(tt.args.category); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrderRequest.SetCargoCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrderRequest_AddParcel(t *testing.T) {
	type args struct {
		parcel *Parcel
	}

	parcel := NewParcel()
	parcels := make([]*dpdSoap.Parcel, 0)

	parcels = append(parcels, &dpdSoap.Parcel{})

	tests := []struct {
		name string
		r    *UpdateOrderRequest
		args args
		want *UpdateOrderRequest
	}{
		{
			"Add parcel",
			NewUpdateOrderRequest(),
			args{
				parcel,
			},
			&UpdateOrderRequest{
				Parcel: parcels,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.AddParcel(tt.args.parcel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrderRequest.AddParcel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateOrderRequest_toDPDRequest(t *testing.T) {
	tests := []struct {
		name string
		r    *UpdateOrderRequest
		want *dpdSoap.DpdOrderCorrection
	}{
		{
			"Convert to DPD request",
			NewUpdateOrderRequest(),
			&dpdSoap.DpdOrderCorrection{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.toDPDRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateOrderRequest.toDPDRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOrder2Cancel(t *testing.T) {
	tests := []struct {
		name string
		want *Order2Cancel
	}{
		{
			"Constructor",
			&Order2Cancel{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrder2Cancel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder2Cancel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder2Cancel_SetInternalOrderNumber(t *testing.T) {
	type args struct {
		number string
	}

	number := "any string"

	tests := []struct {
		name string
		o    *Order2Cancel
		args args
		want *Order2Cancel
	}{
		{
			"Set internal order number",
			NewOrder2Cancel(),
			args{
				number,
			},
			&Order2Cancel{
				OrderNumberInternal: &number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetInternalOrderNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order2Cancel.SetInternalOrderNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder2Cancel_SetDPDOrderNumber(t *testing.T) {
	type args struct {
		number string
	}

	number := "any string"

	tests := []struct {
		name string
		o    *Order2Cancel
		args args
		want *Order2Cancel
	}{
		{
			"Set DPD order number",
			NewOrder2Cancel(),
			args{
				number,
			},
			&Order2Cancel{
				OrderNum: &number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetDPDOrderNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order2Cancel.SetDPDOrderNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder2Cancel_SetPickupDate(t *testing.T) {
	type args struct {
		time time.Time
	}

	now := time.Now()
	dpdDate := dpdSoap.Date(now.Format("2006-01-02"))

	tests := []struct {
		name string
		o    *Order2Cancel
		args args
		want *Order2Cancel
	}{
		{
			"Set pikcup date",
			NewOrder2Cancel(),
			args{
				now,
			},
			&Order2Cancel{
				Pickupdate: &dpdDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.SetPickupDate(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order2Cancel.SetPickupDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCancelOrderRequest(t *testing.T) {
	var orders []*dpdSoap.OrderCancel

	tests := []struct {
		name string
		want *CancelOrderRequest
	}{
		{
			"Constructor",
			&CancelOrderRequest{
				Cancel: orders,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCancelOrderRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCancelOrderRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCancelOrderRequest_AddOrder(t *testing.T) {
	type args struct {
		order *Order2Cancel
	}

	var orders []*dpdSoap.OrderCancel

	orders = append(orders, &dpdSoap.OrderCancel{})

	tests := []struct {
		name string
		r    *CancelOrderRequest
		args args
		want *CancelOrderRequest
	}{
		{
			"Add order",
			NewCancelOrderRequest(),
			args{
				&Order2Cancel{},
			},
			&CancelOrderRequest{
				Cancel: orders,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.AddOrder(tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CancelOrderRequest.AddOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCancelOrderRequest_toDPDRequest(t *testing.T) {
	tests := []struct {
		name string
		r    *CancelOrderRequest
		want *dpdSoap.DpdOrderCancellation
	}{
		{
			"Converter",
			NewCancelOrderRequest(),
			&dpdSoap.DpdOrderCancellation{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.toDPDRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CancelOrderRequest.toDPDRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
