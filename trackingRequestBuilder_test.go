package dpd

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewClientOrderRequest(t *testing.T) {
	tests := []struct {
		name string
		want *TrackByClientOrderRequest
	}{
		{
			"Constructor",
			&TrackByClientOrderRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientOrderRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientOrderRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientOrderRequest_SetClientOrderNumber(t *testing.T) {
	type args struct {
		number string
	}
	orderNumber := anyString

	tests := []struct {
		name string
		r    *TrackByClientOrderRequest
		args args
		want *TrackByClientOrderRequest
	}{
		{
			"Set order number",
			&TrackByClientOrderRequest{},
			args{
				orderNumber,
			},
			&TrackByClientOrderRequest{
				ClientOrderNr: &orderNumber,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetClientOrderNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrackByClientOrderRequest.SetClientOrderNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientOrderRequest_SetPickupDate(t *testing.T) {
	type args struct {
		time time.Time
	}

	now := time.Now()
	dpdDate := now.Format("2006-01-02")

	tests := []struct {
		name string
		r    *TrackByClientOrderRequest
		args args
		want *TrackByClientOrderRequest
	}{
		{
			"Set pickup date",
			&TrackByClientOrderRequest{},
			args{
				now,
			},
			&TrackByClientOrderRequest{
				PickupDate: &dpdDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetPickupDate(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrackByClientOrderRequest.SetPickupDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDpdOrderRequest(t *testing.T) {
	tests := []struct {
		name string
		want *TrackByDPDOrderRequest
	}{
		{
			"Constructor",
			&TrackByDPDOrderRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDpdOrderRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDpdOrderRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDpdOrderRequest_SetDPDOrderNumber(t *testing.T) {
	type args struct {
		number string
	}

	number := anyString

	tests := []struct {
		name string
		r    *TrackByDPDOrderRequest
		args args
		want *TrackByDPDOrderRequest
	}{
		{
			"DPD order number",
			&TrackByDPDOrderRequest{},
			args{
				number,
			},
			&TrackByDPDOrderRequest{
				DpdOrderNr: &number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetDPDOrderNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrackByDPDOrderRequest.SetDPDOrderNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDpdOrderRequest_SetPickupYear(t *testing.T) {
	type args struct {
		year int
	}

	year, err := strconv.Atoi(time.Now().Format("2006"))

	if err != nil {
		t.Errorf("TrackByDPDOrderRequest.SetPickupYear(), err %s", err.Error())
		t.SkipNow()
	}

	tests := []struct {
		name string
		r    *TrackByDPDOrderRequest
		args args
		want *TrackByDPDOrderRequest
	}{
		{
			"Set pickup year",
			&TrackByDPDOrderRequest{},
			args{
				year,
			},
			&TrackByDPDOrderRequest{
				PickupYear: &year,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetPickupYear(tt.args.year); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrackByDPDOrderRequest.SetPickupYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
