package dpd_sdk

import (
	dpdSoap "git.vseinstrumenti.net/golang-sandbox/dpd-soap"
	"time"
)

type createOrderRequest dpdSoap.DpdOrdersData

type extraParameter dpdSoap.Parameter
type extraService dpdSoap.OrderExtraService
type unitLoad dpdSoap.UnitLoad

type CreateOrderRequest interface {
	SetDatePickup(time time.Time) *createOrderRequest
	SetPayer(payer int64) *createOrderRequest
	SetSender(address address) *createOrderRequest
	SetPickupTimePeriod(period string) *createOrderRequest
	SetRegularNum(num string) *createOrderRequest
	AddOrder(order *order) *createOrderRequest

	toDPDRequest() *dpdSoap.DpdOrdersData
}

func NewCreateOrderRequest() *createOrderRequest {
	return new(createOrderRequest)
}

func (r *createOrderRequest) SetDatePickup(time time.Time) *createOrderRequest {
	d := dpdSoap.Date(time.Format("2006-01-02"))
	r.Header.DatePickup = &d

	return r
}

func (r *createOrderRequest) SetPayer(payer int64) *createOrderRequest {
	r.Header.Payer = &payer

	return r
}

func (r *createOrderRequest) SetSender(address address) *createOrderRequest {
	a := dpdSoap.Address(address)
	r.Header.SenderAddress = &a

	return r
}

func (r *createOrderRequest) SetPickupTimePeriod(period string) *createOrderRequest {
	r.Header.PickupTimePeriod = &period

	return r
}

func (r *createOrderRequest) SetRegularNum(num string) *createOrderRequest {
	r.Header.RegularNum = &num

	return r
}

func (r *createOrderRequest) AddOrder(order *order) *createOrderRequest {
	o := dpdSoap.Order(*order)
	r.Order = append(r.Order, &o)

	return r
}

func (r *createOrderRequest) toDPDRequest() *dpdSoap.DpdOrdersData {
	dpdRequest := dpdSoap.DpdOrdersData(*r)

	return &dpdRequest
}

type order dpdSoap.Order

type Order interface {
	SetInternalOrderNumber(number string) *order
	SetServiceCode(code string) *order
	SetServiceVariant(variant string) *order
	SetCargoNumPack(num int) *order
	SetCargoWeight(weight float64) *order
	SetCargoVolume(volume float64) *order
	SetCargoRegistered(flag bool) *order
	SetCargoValue(value float64) *order
	SerCargoCategory(category string) *order
	SetDeliveryTimePeriod(period string) *order
	SetPaymentType(pType string) *order
	SetReceiverAddress(address *address) *order
	SetReturnAddress(address *address) *order

	AddExtraParameter(parameter *extraParameter) *order
	AddExtraService(service *extraService) *order
	AddParcel(parcel *parcel) *order
	AddUnitLoad(load *unitLoad) *order
}

func NewOrder() *order {
	return new(order)
}

func (o *order) SetInternalOrderNumber(number string) *order {
	o.OrderNumberInternal = &number

	return o
}

func (o *order) SetServiceCode(code string) *order {
	o.ServiceCode = &code

	return o
}

func (o *order) SetServiceVariant(variant string) *order {
	o.ServiceVariant = &variant

	return o
}

func (o *order) SetCargoNumPack(num int) *order {
	o.CargoNumPack = &num

	return o
}

func (o *order) SetCargoWeight(weight float64) *order {
	o.CargoWeight = &weight

	return o
}

func (o *order) SetCargoVolume(volume float64) *order {
	o.CargoVolume = &volume

	return o
}

func (o *order) SetCargoRegistered(flag bool) *order {
	o.CargoRegistered = &flag

	return o
}

func (o *order) SetCargoValue(value float64) *order {
	o.CargoValue = &value

	return o
}

func (o *order) SerCargoCategory(category string) *order {
	o.CargoCategory = &category

	return o
}

func (o *order) SetDeliveryTimePeriod(period string) *order {
	o.DeliveryTimePeriod = &period

	return o
}

func (o *order) SetPaymentType(pType string) *order {
	o.PaymentType = &pType

	return o
}

func (o *order) SetReceiverAddress(address *address) *order {
	a := dpdSoap.Address(*address)
	o.ReceiverAddress = &a

	return o
}

func (o *order) SetReturnAddress(address *address) *order {
	a := dpdSoap.Address(*address)
	o.ReturnAddress = &a

	return o
}

func (o *order) AddExtraParameter(parameter *extraParameter) *order {
	p := dpdSoap.Parameter(*parameter)
	o.ExtraParam = append(o.ExtraParam, &p)

	return o
}

func (o *order) AddExtraService(service *extraService) *order {
	s := dpdSoap.OrderExtraService(*service)
	o.ExtraService = append(o.ExtraService, &s)

	return o
}

func (o *order) AddParcel(parcel *parcel) *order {
	p := dpdSoap.Parcel(*parcel)
	o.Parcel = append(o.Parcel, &p)

	return o
}

func (o *order) AddUnitLoad(load *unitLoad) *order {
	u := dpdSoap.UnitLoad(*load)
	o.UnitLoad = append(o.UnitLoad, &u)

	return o
}

type address dpdSoap.Address

type Address interface {
	SetCode(code string) *address
	SetName(name string) *address
	SetTerminalCode(code string) *address
	SetAddressStr(address string) *address
	SetCountryName(name string) *address
	SetZip(zip string) *address
	SetRegion(region string) *address
	SetCity(city string) *address
	SetStreet(street string) *address
	SetStreetAbbr(abbr string) *address
	SetHouse(house string) *address
	SetHousing(housing string) *address
	SetBuilding(building string) *address
	SetPossession(possession string) *address
	SetExtraInfo(info string) *address
	SetOffice(office string) *address
	SetFlat(flat string) *address
	SetWorkTimeFrom(start string) *address
	SetWorkTimeTo(end string) *address
	SetDinnerTimeFrom(start string) *address
	SetDinnerTimeTo(end string) *address
	SetContactFullName(fullName string) *address
	SetContactPhone(phone string) *address
	SetContactEmail(email string) *address
	SetInstructions(instructions string) *address
	SetNeedPass(flag bool) *address
}

func NewAddress() *address {
	return new(address)
}

func (a *address) SetCode(code string) *address {
	a.Code = &code

	return a
}

func (a *address) SetName(name string) *address {
	a.Name = &name

	return a
}

func (a *address) SetTerminalCode(code string) *address {
	a.TerminalCode = &code

	return a
}

func (a *address) SetAddressStr(address string) *address {
	a.AddressString = &address

	return a
}

func (a *address) SetCountryName(name string) *address {
	a.CountryName = &name

	return a
}

func (a *address) SetZip(zip string) *address {
	a.Index = &zip

	return a
}

func (a *address) SetRegion(region string) *address {
	a.Region = &region

	return a
}

func (a *address) SetCity(city string) *address {
	a.City = &city

	return a
}

func (a *address) SetStreet(street string) *address {
	a.Street = &street

	return a
}

func (a *address) SetStreetAbbr(abbr string) *address {
	a.StreetAbbr = &abbr

	return a
}

func (a *address) SetHouse(house string) *address {
	a.House = &house

	return a
}

func (a *address) SetHousing(abbr string) *address {
	a.StreetAbbr = &abbr

	return a
}

func (a *address) SetBuilding(building string) *address {
	a.Str = &building

	return a
}

func (a *address) SetPossession(possession string) *address {
	a.Vlad = &possession

	return a
}

func (a *address) SetExtraInfo(info string) *address {
	a.ExtraInfo = &info

	return a
}

func (a *address) SetOffice(office string) *address {
	a.Office = &office

	return a
}

func (a *address) SetFlat(flat string) *address {
	a.Flat = &flat

	return a
}

func (a *address) SetWorkTimeFrom(start string) *address {
	a.WorkTimeFrom = &start

	return a
}

func (a *address) SetWorkTimeTo(end string) *address {
	a.WorkTimeTo = &end

	return a
}

func (a *address) SetDinnerTimeFrom(start string) *address {
	a.DinnerTimeFrom = &start

	return a
}

func (a *address) SetDinnerTimeTo(end string) *address {
	a.DinnerTimeTo = &end

	return a
}

func (a *address) SetContactFullName(fullName string) *address {
	a.ContactFio = &fullName

	return a
}

func (a *address) SetContactPhone(phone string) *address {
	a.ContactPhone = &phone

	return a
}

func (a *address) SetContactEmail(email string) *address {
	a.ContactEmail = &email

	return a
}

func (a *address) SetInstructions(instructions string) *address {
	a.Instructions = &instructions

	return a
}

func (a *address) SetNeedPass(flag bool) *address {
	a.NeedPass = &flag

	return a
}

type parcel dpdSoap.Parcel

type Parcel interface {
	SetNumber(number string) *parcel
	SetDPDParcelNumber(number int64) *parcel
	SetNumberForPrint(number string) *parcel
	SetBoxNeeded(needed int) *parcel
	SetWeight(weight float64) *parcel
	SetLength(length float64) *parcel
	SetWidth(width float64) *parcel
	SetHeight(height float64) *parcel
	SetInsuranceCost(cost float64) *parcel
	SetCodAmount(amount float64) *parcel
}

func NewParcel() *parcel {
	return new(parcel)
}

func (p *parcel) SetNumber(number string) *parcel {
	p.Number = &number

	return p
}

func (p *parcel) SetDPDParcelNumber(number int64) *parcel {
	p.DpdParcelNumber = &number

	return p
}

func (p *parcel) SetNumberForPrint(number string) *parcel {
	p.NumberForPrint = &number

	return p
}

func (p *parcel) SetBoxNeeded(needed int) *parcel {
	p.BoxNeeded = &needed

	return p
}

func (p *parcel) SetWeight(weight float64) *parcel {
	p.Weight = &weight

	return p
}

func (p *parcel) SetLength(length float64) *parcel {
	p.Length = &length

	return p
}

func (p *parcel) SetWidth(width float64) *parcel {
	p.Width = &width

	return p
}

func (p *parcel) SetHeight(height float64) *parcel {
	p.Height = &height

	return p
}

func (p *parcel) SetInsuranceCost(cost float64) *parcel {
	p.InsuranceCost = &cost

	return p
}

func (p *parcel) SetCodAmount(amount float64) *parcel {
	p.CodAmount = &amount

	return p
}

type updateOrderRequest dpdSoap.DpdOrderCorrection

type UpdateOrderRequest interface {
	SetDPDOrderNumber(number string) *updateOrderRequest
	SetInternalOrderNumber(number string) *updateOrderRequest
	SetCargoNumPack(num int) *updateOrderRequest
	SetCargoWeight(weight float64) *updateOrderRequest
	SetCargoVolume(volume float64) *updateOrderRequest
	SetCargoValue(value float64) *updateOrderRequest
	SetCargoCategory(category string) *updateOrderRequest

	AddParcel(parcel *parcel) *updateOrderRequest

	toDPDRequest() *dpdSoap.DpdOrderCorrection
}

func NewUpdateOrderRequest() *updateOrderRequest {
	return new(updateOrderRequest)
}

func (r *updateOrderRequest) SetDPDOrderNumber(number string) *updateOrderRequest {
	r.OrderNum = &number

	return r
}

func (r *updateOrderRequest) SetInternalOrderNumber(number string) *updateOrderRequest {
	r.OrderNumberInternal = &number

	return r
}

func (r *updateOrderRequest) SetCargoNumPack(num int) *updateOrderRequest {
	r.CargoNumPack = &num

	return r
}

func (r *updateOrderRequest) SetCargoWeight(weight float64) *updateOrderRequest {
	r.CargoWeight = &weight

	return r
}

func (r *updateOrderRequest) SetCargoVolume(volume float64) *updateOrderRequest {
	r.CargoVolume = &volume

	return r
}

func (r *updateOrderRequest) SetCargoValue(value float64) *updateOrderRequest {
	r.CargoValue = &value

	return r
}

func (r *updateOrderRequest) SetCargoCategory(category string) *updateOrderRequest {
	r.CargoCategory = &category

	return r
}

func (r *updateOrderRequest) AddParcel(parcel *parcel) *updateOrderRequest {
	p := dpdSoap.Parcel(*parcel)
	r.Parcel = append(r.Parcel, &p)

	return r
}

func (r *updateOrderRequest) toDPDRequest() *dpdSoap.DpdOrderCorrection {
	dpdReq := dpdSoap.DpdOrderCorrection(*r)

	return &dpdReq
}

type order2Cancel dpdSoap.OrderCancel

type Order2Cancel interface {
	SetInternalOrderNumber(number string) *order2Cancel
	SetDPDOrderNumber(number string) *order2Cancel
	SetPickupDate(time time.Time) *order2Cancel
}

func NewOrder2Cancel() *order2Cancel {
	return new(order2Cancel)
}

func (o *order2Cancel) SetInternalOrderNumber(number string) *order2Cancel {
	o.OrderNumberInternal = &number

	return o
}

func (o *order2Cancel) SetDPDOrderNumber(number string) *order2Cancel {
	o.OrderNum = &number

	return o
}

func (o *order2Cancel) SetPickupDate(time time.Time) *order2Cancel {
	d := dpdSoap.Date(time.Format("2016-01-02"))
	o.Pickupdate = &d

	return o
}

type cancelOrderRequest dpdSoap.DpdOrderCancellation

type CancelOrderRequest interface {
	AddOrder(order order2Cancel) *cancelOrderRequest

	toDPDRequest() *dpdSoap.DpdOrderCancellation
}

func NewCancelOrderRequest() *cancelOrderRequest {
	return new(cancelOrderRequest)
}

func (r *cancelOrderRequest) AddOrder(order *order2Cancel) *cancelOrderRequest {
	o := dpdSoap.OrderCancel(*order)
	r.Cancel = append(r.Cancel, &o)

	return r
}

func (r *cancelOrderRequest) toDPDRequest() *dpdSoap.DpdOrderCancellation {
	dpdReq := dpdSoap.DpdOrderCancellation(*r)

	return &dpdReq
}
