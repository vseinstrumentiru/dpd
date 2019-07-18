package dpd

import (
	"reflect"
	"testing"
)

func TestNewParcelShopRequest(t *testing.T) {
	tests := []struct {
		name string
		want *ParcelShopRequest
	}{
		{
			"New parcel shop",
			&ParcelShopRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParcelShopRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParcelShopRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcelShopRequest_SetCountryCode(t *testing.T) {
	type args struct {
		code string
	}

	countryCode := "RU"

	tests := []struct {
		name string
		r    *ParcelShopRequest
		args args
		want *ParcelShopRequest
	}{
		{
			"Set country code",
			&ParcelShopRequest{},
			args{
				countryCode,
			},
			&ParcelShopRequest{
				CountryCode: &countryCode,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetCountryCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParcelShopRequest.SetCountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcelShopRequest_SetRegionCode(t *testing.T) {
	type args struct {
		code string
	}

	code := "234324"

	tests := []struct {
		name string
		r    *ParcelShopRequest
		args args
		want *ParcelShopRequest
	}{
		{
			"Set region code",
			&ParcelShopRequest{},
			args{
				code,
			},
			&ParcelShopRequest{
				RegionCode: &code,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetRegionCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParcelShopRequest.SetRegionCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcelShopRequest_SetCityCode(t *testing.T) {
	type args struct {
		code string
	}

	cityCode := "123ffew!"

	tests := []struct {
		name string
		r    *ParcelShopRequest
		args args
		want *ParcelShopRequest
	}{
		{
			"Set city code",
			&ParcelShopRequest{},
			args{
				cityCode,
			},
			&ParcelShopRequest{
				CityCode: &cityCode,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetCityCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParcelShopRequest.SetCityCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParcelShopRequest_SetCityName(t *testing.T) {
	type args struct {
		name string
	}

	cityName := "City name"

	tests := []struct {
		name string
		r    *ParcelShopRequest
		args args
		want *ParcelShopRequest
	}{
		{
			"Set city name",
			&ParcelShopRequest{},
			args{
				cityName,
			},
			&ParcelShopRequest{
				CityName: &cityName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetCityName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParcelShopRequest.SetCityName() = %v, want %v", got, tt.want)
			}
		})
	}
}
