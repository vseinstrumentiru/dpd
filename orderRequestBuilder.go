package dpd

import (
	"time"
)

//NewCreateOrderRequest ...
func NewCreateOrderRequest() *CreateOrderRequest {
	var orders []*Order

	req := &CreateOrderRequest{
		Header: &Header{},
		Order:  orders,
	}

	return req
}

//SetDatePickup ...
func (r *CreateOrderRequest) SetDatePickup(time time.Time) *CreateOrderRequest {
	d := time.Format("2006-01-02")
	r.Header.DatePickup = &d

	return r
}

//SetPayer ...
func (r *CreateOrderRequest) SetPayer(payer int64) *CreateOrderRequest {
	r.Header.Payer = &payer

	return r
}

//SetSender ...
func (r *CreateOrderRequest) SetSender(address Address) *CreateOrderRequest {
	r.Header.SenderAddress = &address

	return r
}

//SetPickupTimePeriod ...
func (r *CreateOrderRequest) SetPickupTimePeriod(period string) *CreateOrderRequest {
	r.Header.PickupTimePeriod = &period

	return r
}

//SetRegularNum regular DPD order number. Ask your manager about correct number if you use it
func (r *CreateOrderRequest) SetRegularNum(num string) *CreateOrderRequest {
	r.Header.RegularNum = &num

	return r
}

//AddOrder to CreateOrderRequest
func (r *CreateOrderRequest) AddOrder(order *Order) *CreateOrderRequest {
	r.Order = append(r.Order, order)

	return r
}

//NewOrder ...
func NewOrder() *Order {
	var parcel []*Parcel
	var extraService []*ExtraService
	var extraParameter []*ExtraParameter
	var unitLoad []*UnitLoadOrder

	return &Order{
		ExtraParam:   extraParameter,
		ExtraService: extraService,
		Parcel:       parcel,
		UnitLoad:     unitLoad,
	}
}

//SetInternalOrderNumber imply client side order number
func (o *Order) SetInternalOrderNumber(number string) *Order {
	o.OrderNumberInternal = &number

	return o
}

//SetServiceCode code of DPD service
func (o *Order) SetServiceCode(code string) *Order {
	o.ServiceCode = &code

	return o
}

//SetServiceVariant delivery variant& Four variants are available(ДД, ДТ, ТД, ТТ)
func (o *Order) SetServiceVariant(variant string) *Order {
	o.ServiceVariant = &variant

	return o
}

//SetCargoCount ...
func (o *Order) SetCargoCount(num int) *Order {
	o.CargoNumPack = &num

	return o
}

//SetCargoWeight ...
func (o *Order) SetCargoWeight(weight float64) *Order {
	o.CargoWeight = &weight

	return o
}

//SetCargoVolume ...
func (o *Order) SetCargoVolume(volume float64) *Order {
	o.CargoVolume = &volume

	return o
}

//SetValuableCargo ...
func (o *Order) SetValuableCargo(flag bool) *Order {
	o.CargoRegistered = &flag

	return o
}

//SetCargoValue ...
func (o *Order) SetCargoValue(value float64) *Order {
	o.CargoValue = &value

	return o
}

//SerCargoCategory ...
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
	o.ReceiverAddress = address

	return o
}

func (o *Order) SetReturnAddress(address *Address) *Order {
	o.ReturnAddress = address

	return o
}

func (o *Order) AddExtraParameter(parameter *ExtraParameter) *Order {
	o.ExtraParam = append(o.ExtraParam, parameter)

	return o
}

func (o *Order) AddExtraService(service *ExtraService) *Order {
	o.ExtraService = append(o.ExtraService, service)

	return o
}

func (o *Order) AddParcel(parcel *Parcel) *Order {
	o.Parcel = append(o.Parcel, parcel)

	return o
}

func (o *Order) AddUnitLoad(load *UnitLoadOrder) *Order {
	o.UnitLoad = append(o.UnitLoad, load)

	return o
}

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

func NewUpdateOrderRequest() *UpdateOrderRequest {
	var parcels []*Parcel

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
	r.Parcel = append(r.Parcel, parcel)

	return r
}

func NewOrderToCancel() *OrderToCancel {
	return new(OrderToCancel)
}

func (o *OrderToCancel) SetInternalOrderNumber(number string) *OrderToCancel {
	o.OrderNumberInternal = &number

	return o
}

func (o *OrderToCancel) SetDPDOrderNumber(number string) *OrderToCancel {
	o.OrderNum = &number

	return o
}

func (o *OrderToCancel) SetPickupDate(time time.Time) *OrderToCancel {
	d := time.Format("2006-01-02")
	o.PickupDate = &d

	return o
}

func NewCancelOrderRequest() *CancelOrderRequest {
	var orders []*OrderToCancel

	req := new(CancelOrderRequest)
	req.Cancel = orders

	return req
}

func (r *CancelOrderRequest) AddOrder(order *OrderToCancel) *CancelOrderRequest {
	r.Cancel = append(r.Cancel, order)

	return r
}
