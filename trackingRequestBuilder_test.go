package dpd

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	dpdSoap "github.com/seacomandor/dpd/soap"
)

func TestNewClientOrderRequest(t *testing.T) {
	tests := []struct {
		name string
		want *ClientOrderRequest
	}{
		{
			"Constructor",
			&ClientOrderRequest{},
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
	orderNumber := "any string"

	tests := []struct {
		name string
		r    *ClientOrderRequest
		args args
		want *ClientOrderRequest
	}{
		{
			"Set order number",
			&ClientOrderRequest{},
			args{
				orderNumber,
			},
			&ClientOrderRequest{
				ClientOrderNr: &orderNumber,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetClientOrderNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientOrderRequest.SetClientOrderNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientOrderRequest_SetPickupDate(t *testing.T) {
	type args struct {
		time time.Time
	}

	now := time.Now()
	dpdDate := dpdSoap.Date(now.Format("2006-01-02"))

	tests := []struct {
		name string
		r    *ClientOrderRequest
		args args
		want *ClientOrderRequest
	}{
		{
			"Set pickup date",
			&ClientOrderRequest{},
			args{
				now,
			},
			&ClientOrderRequest{
				PickupDate: &dpdDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetPickupDate(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientOrderRequest.SetPickupDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientOrderRequest_toDPDRequest(t *testing.T) {
	tests := []struct {
		name string
		r    *ClientOrderRequest
		want *dpdSoap.RequestClientOrder
	}{
		{
			"Converter",
			&ClientOrderRequest{},
			&dpdSoap.RequestClientOrder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.toDPDRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientOrderRequest.toDPDRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClientParcelRequest(t *testing.T) {
	tests := []struct {
		name string
		want *ClientParcelRequest
	}{
		{
			"Constructor",
			&ClientParcelRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientParcelRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientParcelRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientParcelRequest_SetClientParcelNumber(t *testing.T) {
	type args struct {
		number string
	}

	number := "any string"

	tests := []struct {
		name string
		r    *ClientParcelRequest
		args args
		want *ClientParcelRequest
	}{
		{
			"Set client parcel number",
			&ClientParcelRequest{},
			args{
				number,
			},
			&ClientParcelRequest{
				ClientParcelNr: &number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetClientParcelNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientParcelRequest.SetClientParcelNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientParcelRequest_SetPickupDate(t *testing.T) {
	type args struct {
		time time.Time
	}

	now := time.Now()
	dpdDate := dpdSoap.Date(now.Format("2006-01-02"))

	tests := []struct {
		name string
		r    *ClientParcelRequest
		args args
		want *ClientParcelRequest
	}{
		{
			"Set pickup date",
			&ClientParcelRequest{},
			args{
				now,
			},
			&ClientParcelRequest{
				PickupDate: &dpdDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetPickupDate(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientParcelRequest.SetPickupDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDpdOrderRequest(t *testing.T) {
	tests := []struct {
		name string
		want *DpdOrderRequest
	}{
		{
			"Constructor",
			&DpdOrderRequest{},
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

	number := "any string"

	tests := []struct {
		name string
		r    *DpdOrderRequest
		args args
		want *DpdOrderRequest
	}{
		{
			"DPD order number",
			&DpdOrderRequest{},
			args{
				number,
			},
			&DpdOrderRequest{
				DpdOrderNr: &number,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetDPDOrderNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DpdOrderRequest.SetDPDOrderNumber() = %v, want %v", got, tt.want)
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
		t.Errorf("DpdOrderRequest.SetPickupYear(), err %s", err.Error())
		t.SkipNow()
	}

	tests := []struct {
		name string
		r    *DpdOrderRequest
		args args
		want *DpdOrderRequest
	}{
		{
			"Set pickup year",
			&DpdOrderRequest{},
			args{
				year,
			},
			&DpdOrderRequest{
				PickupYear: &year,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetPickupYear(tt.args.year); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DpdOrderRequest.SetPickupYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDpdOrderRequest_toDPDRequest(t *testing.T) {
	tests := []struct {
		name string
		r    *DpdOrderRequest
		want *dpdSoap.RequestDpdOrder
	}{
		{
			"Converter",
			&DpdOrderRequest{},
			&dpdSoap.RequestDpdOrder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.toDPDRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DpdOrderRequest.toDPDRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfirmRequest(t *testing.T) {
	tests := []struct {
		name string
		want *ConfirmRequest
	}{
		{
			"Constructor",
			&ConfirmRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfirmRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfirmRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfirmRequest_SetDocId(t *testing.T) {
	type args struct {
		docId int64
	}

	var docID int64 = 12341241242

	tests := []struct {
		name string
		r    *ConfirmRequest
		args args
		want *ConfirmRequest
	}{
		{
			"Set doc id",
			&ConfirmRequest{},
			args{
				docID,
			},
			&ConfirmRequest{
				DocId: &docID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetDocId(tt.args.docId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfirmRequest.SetDocId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDpdTrackingEventRequest_SetDateFrom(t *testing.T) {
	type args struct {
		from time.Time
	}

	from := time.Now()
	dpdDateTime := dpdSoap.DateTime(from.Format("2006-01-02 15:04:05"))

	tests := []struct {
		name string
		r    *DpdTrackingEventRequest
		args args
		want *DpdTrackingEventRequest
	}{
		{
			"Set date from",
			&DpdTrackingEventRequest{},
			args{
				from,
			},
			&DpdTrackingEventRequest{
				DateFrom: &dpdDateTime,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetDateFrom(tt.args.from); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DpdTrackingEventRequest.SetDateFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDpdTrackingEventRequest_SetDateTo(t *testing.T) {
	type args struct {
		to time.Time
	}

	to := time.Now()
	dpdDateTime := dpdSoap.DateTime(to.Format("2006-01-02 15:04:05"))

	tests := []struct {
		name string
		r    *DpdTrackingEventRequest
		args args
		want *DpdTrackingEventRequest
	}{
		{
			"Set date to",
			&DpdTrackingEventRequest{},
			args{
				to,
			},
			&DpdTrackingEventRequest{
				DateTo: &dpdDateTime,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetDateTo(tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DpdTrackingEventRequest.SetDateTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDpdTrackingEventRequest_SetMaxRowCount(t *testing.T) {
	type args struct {
		count int
	}

	count := 2

	tests := []struct {
		name string
		r    *DpdTrackingEventRequest
		args args
		want *DpdTrackingEventRequest
	}{
		{
			"Max row count",
			&DpdTrackingEventRequest{},
			args{
				count,
			},
			&DpdTrackingEventRequest{
				MaxRowCount: &count,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetMaxRowCount(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DpdTrackingEventRequest.SetMaxRowCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDpdTrackingEventRequest_toDPDRequest(t *testing.T) {
	tests := []struct {
		name string
		r    *DpdTrackingEventRequest
		want *dpdSoap.EventTrackingRequest
	}{
		{
			"Converter",
			&DpdTrackingEventRequest{},
			&dpdSoap.EventTrackingRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.toDPDRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DpdTrackingEventRequest.toDPDRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
