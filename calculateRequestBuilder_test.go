package dpd

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewCity(t *testing.T) {
	type args struct {
		cityID int64
	}

	var cityID int64 = 234

	tests := []struct {
		name string
		args args
		want *CityRequest
	}{
		{
			"City " + strconv.Itoa(int(cityID)),
			args{
				cityID: cityID,
			},
			&CityRequest{
				CityID: &cityID,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCity(tt.args.cityID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCityRequest_SetIndex(t *testing.T) {
	type args struct {
		index string
	}

	index := "603000"

	tests := []struct {
		name string
		c    *CityRequest
		args args
		want *CityRequest
	}{
		{
			"Set index",
			&CityRequest{},
			args{
				index: index,
			},
			&CityRequest{
				Index: &index,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.SetIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CityRequest.SetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCityRequest_SetCityName(t *testing.T) {
	type args struct {
		name string
	}

	cityName := "Город которого нет"

	tests := []struct {
		name string
		c    *CityRequest
		args args
		want *CityRequest
	}{
		{
			"Set city name",
			&CityRequest{},
			args{
				cityName,
			},
			&CityRequest{
				CityName: &cityName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.SetCityName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CityRequest.SetCityName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCityRequest_SetRegionCode(t *testing.T) {
	type args struct {
		code int
	}

	regionCode := 123

	tests := []struct {
		name string
		c    *CityRequest
		args args
		want *CityRequest
	}{
		{
			"Set region code",
			&CityRequest{},
			args{
				regionCode,
			},
			&CityRequest{
				RegionCode: &regionCode,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.SetRegionCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CityRequest.SetRegionCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCityRequest_SetCountryCode(t *testing.T) {
	type args struct {
		code string
	}

	countyCode := "RU"

	tests := []struct {
		name string
		c    *CityRequest
		args args
		want *CityRequest
	}{
		{
			"Set country code",
			&CityRequest{},
			args{
				countyCode,
			},
			&CityRequest{
				CountryCode: &countyCode,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.SetCountryCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CityRequest.SetCountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCalculateRequest(t *testing.T) {
	type args struct {
		from         *CityRequest
		to           *CityRequest
		weight       float64
		selfPickup   bool
		selfDelivery bool
	}

	from := NewCity(123)
	to := NewCity(321)
	weight := 12.5
	selfPickup := false
	selfDelivery := false

	tests := []struct {
		name string
		args args
		want *CalculateRequest
	}{
		{
			"Constructor",
			args{
				from,
				to,
				weight,
				selfPickup,
				selfDelivery,
			},
			&CalculateRequest{
				Pickup:       from,
				Delivery:     to,
				Weight:       &weight,
				SelfPickup:   &selfPickup,
				SelfDelivery: &selfDelivery,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCalculateRequest(
				tt.args.from,
				tt.args.to,
				tt.args.weight,
				tt.args.selfPickup,
				tt.args.selfDelivery,
			); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCalculateRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetPickup(t *testing.T) {
	type args struct {
		city *CityRequest
	}

	cityRequest := NewCity(123)

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set pickup",
			&CalculateRequest{},
			args{
				cityRequest,
			},
			&CalculateRequest{
				Pickup: cityRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.OverrideFrom(tt.args.city); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.OverrideFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetDelivery(t *testing.T) {
	type args struct {
		city *CityRequest
	}

	cityRequest := NewCity(123)

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set delivery",
			&CalculateRequest{},
			args{
				cityRequest,
			},
			&CalculateRequest{
				Delivery: cityRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.OverrideTo(tt.args.city); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.OverrideTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetWeight(t *testing.T) {
	type args struct {
		weight float64
	}

	weight := 1.23

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set weight",
			&CalculateRequest{},
			args{
				weight,
			},
			&CalculateRequest{
				Weight: &weight,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetWeight(tt.args.weight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.SetWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetVolume(t *testing.T) {
	type args struct {
		volume float64
	}

	volume := 12.34

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set volume",
			&CalculateRequest{},
			args{
				volume,
			},
			&CalculateRequest{
				Volume: &volume,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetVolume(tt.args.volume); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.SetVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetServiceCode(t *testing.T) {
	type args struct {
		code string
	}

	serviceCode := "BZN,GBP,SCN"

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set service cost",
			&CalculateRequest{},
			args{
				serviceCode,
			},
			&CalculateRequest{
				ServiceCode: &serviceCode,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetServiceCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.SetServiceCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetPickupDate(t *testing.T) {
	type args struct {
		time time.Time
	}

	currentTime := time.Now()
	dpdDate := currentTime.Format("2006-01-02")

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set pickup date",
			&CalculateRequest{},
			args{
				currentTime,
			},
			&CalculateRequest{
				PickupDate: &dpdDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetPickupDate(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.SetPickupDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetMaxDays(t *testing.T) {
	type args struct {
		days int
	}

	days := 1

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set max days",
			&CalculateRequest{},
			args{
				days,
			},
			&CalculateRequest{
				MaxDays: &days,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetMaxDays(tt.args.days); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.SetMaxDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetMaxCost(t *testing.T) {
	type args struct {
		cost float64
	}

	cost := 123.23

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set max cost",
			&CalculateRequest{},
			args{
				cost,
			},
			&CalculateRequest{
				MaxCost: &cost,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetMaxCost(tt.args.cost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.SetMaxCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetDeclaredValue(t *testing.T) {
	type args struct {
		declaredValue float64
	}

	declaredValue := 2123.43

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set declared value",
			&CalculateRequest{},
			args{
				declaredValue,
			},
			&CalculateRequest{
				DeclaredValue: &declaredValue,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetDeclaredValue(tt.args.declaredValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.SetDeclaredValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetSelfPickup(t *testing.T) {
	type args struct {
		flag bool
	}

	pickup := true

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set self pickup",
			&CalculateRequest{},
			args{
				pickup,
			},
			&CalculateRequest{
				SelfPickup: &pickup,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetSelfPickup(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.SetSelfPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateRequest_SetSelfDelivery(t *testing.T) {
	type args struct {
		flag bool
	}

	delivery := false

	tests := []struct {
		name string
		r    *CalculateRequest
		args args
		want *CalculateRequest
	}{
		{
			"Set self delivery",
			&CalculateRequest{},
			args{
				delivery,
			},
			&CalculateRequest{
				SelfDelivery: &delivery,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetSelfDelivery(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateRequest.SetSelfDelivery() = %v, want %v", got, tt.want)
			}
		})
	}
}
