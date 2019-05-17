package dpd

import (
	dpdSoap "git.vseinstrumenti.net/golang-sandbox/dpd/soap"
	"time"
)

// Запрос на создание заказа
type CreateOrderRequest dpdSoap.DpdOrdersData

type ExtraParameter dpdSoap.Parameter
type ExtraService dpdSoap.OrderExtraService
type UnitLoad dpdSoap.UnitLoad

func NewCreateOrderRequest() *CreateOrderRequest {
	var orders []*dpdSoap.Order

	req := &CreateOrderRequest{
		Header: &dpdSoap.Header{},
		Order:  orders,
	}

	return req
}

func (r *CreateOrderRequest) SetDatePickup(time time.Time) *CreateOrderRequest {
	d := dpdSoap.Date(time.Format("2006-01-02"))
	r.Header.DatePickup = &d

	return r
}

func (r *CreateOrderRequest) SetPayer(payer int64) *CreateOrderRequest {
	r.Header.Payer = &payer

	return r
}

func (r *CreateOrderRequest) SetSender(address Address) *CreateOrderRequest {
	a := dpdSoap.Address(address)
	r.Header.SenderAddress = &a

	return r
}

func (r *CreateOrderRequest) SetPickupTimePeriod(period string) *CreateOrderRequest {
	r.Header.PickupTimePeriod = &period

	return r
}

func (r *CreateOrderRequest) SetRegularNum(num string) *CreateOrderRequest {
	r.Header.RegularNum = &num

	return r
}

func (r *CreateOrderRequest) AddOrder(order *Order) *CreateOrderRequest {
	o := dpdSoap.Order(*order)
	r.Order = append(r.Order, &o)

	return r
}

func (r *CreateOrderRequest) toDPDRequest() *dpdSoap.DpdOrdersData {
	dpdRequest := dpdSoap.DpdOrdersData(*r)

	return &dpdRequest
}

type Order dpdSoap.Order

func NewOrder() *Order {
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

func (o *Order) SetInternalOrderNumber(number string) *Order {
	o.OrderNumberInternal = &number

	return o
}

func (o *Order) SetServiceCode(code string) *Order {
	o.ServiceCode = &code

	return o
}

func (o *Order) SetServiceVariant(variant string) *Order {
	o.ServiceVariant = &variant

	return o
}

func (o *Order) SetCargoNumPack(num int) *Order {
	o.CargoNumPack = &num

	return o
}

func (o *Order) SetCargoWeight(weight float64) *Order {
	o.CargoWeight = &weight

	return o
}

func (o *Order) SetCargoVolume(volume float64) *Order {
	o.CargoVolume = &volume

	return o
}

func (o *Order) SetCargoRegistered(flag bool) *Order {
	o.CargoRegistered = &flag

	return o
}

func (o *Order) SetCargoValue(value float64) *Order {
	o.CargoValue = &value

	return o
}

func (o *Order) SerCargoCategory(category string) *Order {
	o.CargoCategory = &category

	return o
}

func (o *Order) SetDeliveryTimePeriod(period string) *Order {
	o.DeliveryTimePeriod = &period

	return o
}

func (o *Order) SetPaymentType(pType string) *Order {
	o.PaymentType = &pType

	return o
}

func (o *Order) SetReceiverAddress(address *Address) *Order {
	a := dpdSoap.Address(*address)
	o.ReceiverAddress = &a

	return o
}

func (o *Order) SetReturnAddress(address *Address) *Order {
	a := dpdSoap.Address(*address)
	o.ReturnAddress = &a

	return o
}

func (o *Order) AddExtraParameter(parameter *ExtraParameter) *Order {
	p := dpdSoap.Parameter(*parameter)
	o.ExtraParam = append(o.ExtraParam, &p)

	return o
}

func (o *Order) AddExtraService(service *ExtraService) *Order {
	s := dpdSoap.OrderExtraService(*service)
	o.ExtraService = append(o.ExtraService, &s)

	return o
}

func (o *Order) AddParcel(parcel *Parcel) *Order {
	p := dpdSoap.Parcel(*parcel)
	o.Parcel = append(o.Parcel, &p)

	return o
}

func (o *Order) AddUnitLoad(load *UnitLoad) *Order {
	u := dpdSoap.UnitLoad(*load)
	o.UnitLoad = append(o.UnitLoad, &u)

	return o
}

type Address dpdSoap.Address

func NewAddress() *Address {
	return new(Address)
}

func (a *Address) SetCode(code string) *Address {
	a.Code = &code

	return a
}

func (a *Address) SetName(name string) *Address {
	a.Name = &name

	return a
}

func (a *Address) SetTerminalCode(code string) *Address {
	a.TerminalCode = &code

	return a
}

func (a *Address) SetAddressString(address string) *Address {
	a.AddressString = &address

	return a
}

func (a *Address) SetCountryName(name string) *Address {
	a.CountryName = &name

	return a
}

func (a *Address) SetZip(zip string) *Address {
	a.Index = &zip

	return a
}

func (a *Address) SetRegion(region string) *Address {
	a.Region = &region

	return a
}

func (a *Address) SetCity(city string) *Address {
	a.City = &city

	return a
}

func (a *Address) SetStreet(street string) *Address {
	a.Street = &street

	return a
}

func (a *Address) SetStreetAbbr(abbr string) *Address {
	a.StreetAbbr = &abbr

	return a
}

func (a *Address) SetHouse(house string) *Address {
	a.House = &house

	return a
}

func (a *Address) SetHousing(housing string) *Address {
	a.HouseKorpus = &housing

	return a
}

func (a *Address) SetBuilding(building string) *Address {
	a.Str = &building

	return a
}

func (a *Address) SetPossession(possession string) *Address {
	a.Vlad = &possession

	return a
}

func (a *Address) SetExtraInfo(info string) *Address {
	a.ExtraInfo = &info

	return a
}

func (a *Address) SetOffice(office string) *Address {
	a.Office = &office

	return a
}

func (a *Address) SetFlat(flat string) *Address {
	a.Flat = &flat

	return a
}

func (a *Address) SetWorkTimeFrom(start string) *Address {
	a.WorkTimeFrom = &start

	return a
}

func (a *Address) SetWorkTimeTo(end string) *Address {
	a.WorkTimeTo = &end

	return a
}

func (a *Address) SetDinnerTimeFrom(start string) *Address {
	a.DinnerTimeFrom = &start

	return a
}

func (a *Address) SetDinnerTimeTo(end string) *Address {
	a.DinnerTimeTo = &end

	return a
}

func (a *Address) SetContactFullName(fullName string) *Address {
	a.ContactFio = &fullName

	return a
}

func (a *Address) SetContactPhone(phone string) *Address {
	a.ContactPhone = &phone

	return a
}

func (a *Address) SetContactEmail(email string) *Address {
	a.ContactEmail = &email

	return a
}

func (a *Address) SetInstructions(instructions string) *Address {
	a.Instructions = &instructions

	return a
}

func (a *Address) SetNeedPass(flag bool) *Address {
	a.NeedPass = &flag

	return a
}

type Parcel dpdSoap.Parcel

func NewParcel() *Parcel {
	return new(Parcel)
}

func (p *Parcel) SetNumber(number string) *Parcel {
	p.Number = &number

	return p
}

func (p *Parcel) SetDPDParcelNumber(number int64) *Parcel {
	p.DpdParcelNumber = &number

	return p
}

func (p *Parcel) SetNumberForPrint(number string) *Parcel {
	p.NumberForPrint = &number

	return p
}

func (p *Parcel) SetBoxNeeded(needed int) *Parcel {
	p.BoxNeeded = &needed

	return p
}

func (p *Parcel) SetWeight(weight float64) *Parcel {
	p.Weight = &weight

	return p
}

func (p *Parcel) SetLength(length float64) *Parcel {
	p.Length = &length

	return p
}

func (p *Parcel) SetWidth(width float64) *Parcel {
	p.Width = &width

	return p
}

func (p *Parcel) SetHeight(height float64) *Parcel {
	p.Height = &height

	return p
}

func (p *Parcel) SetInsuranceCost(cost float64) *Parcel {
	p.InsuranceCost = &cost

	return p
}

func (p *Parcel) SetCodAmount(amount float64) *Parcel {
	p.CodAmount = &amount

	return p
}

type UpdateOrderRequest dpdSoap.DpdOrderCorrection

func NewUpdateOrderRequest() *UpdateOrderRequest {
	var parcels []*dpdSoap.Parcel

	req := new(UpdateOrderRequest)
	req.Parcel = parcels

	return new(UpdateOrderRequest)
}

func (r *UpdateOrderRequest) SetDPDOrderNumber(number string) *UpdateOrderRequest {
	r.OrderNum = &number

	return r
}

func (r *UpdateOrderRequest) SetInternalOrderNumber(number string) *UpdateOrderRequest {
	r.OrderNumberInternal = &number

	return r
}

func (r *UpdateOrderRequest) SetCargoNumPack(num int) *UpdateOrderRequest {
	r.CargoNumPack = &num

	return r
}

func (r *UpdateOrderRequest) SetCargoWeight(weight float64) *UpdateOrderRequest {
	r.CargoWeight = &weight

	return r
}

func (r *UpdateOrderRequest) SetCargoVolume(volume float64) *UpdateOrderRequest {
	r.CargoVolume = &volume

	return r
}

func (r *UpdateOrderRequest) SetCargoValue(value float64) *UpdateOrderRequest {
	r.CargoValue = &value

	return r
}

func (r *UpdateOrderRequest) SetCargoCategory(category string) *UpdateOrderRequest {
	r.CargoCategory = &category

	return r
}

func (r *UpdateOrderRequest) AddParcel(parcel *Parcel) *UpdateOrderRequest {
	p := dpdSoap.Parcel(*parcel)
	r.Parcel = append(r.Parcel, &p)

	return r
}

func (r *UpdateOrderRequest) toDPDRequest() *dpdSoap.DpdOrderCorrection {
	dpdReq := dpdSoap.DpdOrderCorrection(*r)

	return &dpdReq
}

type Order2Cancel dpdSoap.OrderCancel

func NewOrder2Cancel() *Order2Cancel {
	return new(Order2Cancel)
}

func (o *Order2Cancel) SetInternalOrderNumber(number string) *Order2Cancel {
	o.OrderNumberInternal = &number

	return o
}

func (o *Order2Cancel) SetDPDOrderNumber(number string) *Order2Cancel {
	o.OrderNum = &number

	return o
}

func (o *Order2Cancel) SetPickupDate(time time.Time) *Order2Cancel {
	d := dpdSoap.Date(time.Format("2006-01-02"))
	o.Pickupdate = &d

	return o
}

type CancelOrderRequest dpdSoap.DpdOrderCancellation

func NewCancelOrderRequest() *CancelOrderRequest {
	var orders []*dpdSoap.OrderCancel

	req := new(CancelOrderRequest)
	req.Cancel = orders

	return req
}

func (r *CancelOrderRequest) AddOrder(order *Order2Cancel) *CancelOrderRequest {
	o := dpdSoap.OrderCancel(*order)
	r.Cancel = append(r.Cancel, &o)

	return r
}

func (r *CancelOrderRequest) toDPDRequest() *dpdSoap.DpdOrderCancellation {
	dpdReq := dpdSoap.DpdOrderCancellation(*r)

	return &dpdReq
}
