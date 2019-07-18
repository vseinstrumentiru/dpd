package dpd

const orderNamespace = "http://dpd.ru/ws/order2/2012-04-04"

type operationAddParcels struct {
	AddParcels *addParcels `xml:"addParcels,omitempty"`
}

type addParcels struct {
	Parcels *UpdateOrderRequest `xml:"parcels,omitempty"`
}

//UpdateOrderRequest AddParcels\RemoveParcels request body
type UpdateOrderRequest struct {
	Auth                *Auth     `xml:"auth,omitempty"`
	OrderNum            *string   `xml:"orderNum,omitempty"`
	OrderNumberInternal *string   `xml:"orderNumberInternal,omitempty"`
	CargoNumPack        *int      `xml:"cargoNumPack,omitempty"`
	CargoWeight         *float64  `xml:"cargoWeight,omitempty"`
	CargoVolume         *float64  `xml:"cargoVolume,omitempty"`
	CargoValue          *float64  `xml:"cargoValue,omitempty"`
	CargoCategory       *string   `xml:"cargoCategory,omitempty"`
	Parcel              []*Parcel `xml:"parcel,omitempty"`
}

//Parcel to add or remove
type Parcel struct {
	Number           *string  `xml:"number,omitempty"`
	DpdParcelNumber  *int64   `xml:"dpdParcelNumber,omitempty"`
	NumberForPrint   *string  `xml:"number_for_print,omitempty"`
	BoxNeeded        *int     `xml:"box_needed,omitempty"`
	Weight           *float64 `xml:"weight,omitempty"`
	Length           *float64 `xml:"length,omitempty"`
	Width            *float64 `xml:"width,omitempty"`
	Height           *float64 `xml:"height,omitempty"`
	InsuranceCost    *float64 `xml:"insuranceCost,omitempty"`
	InsuranceCostVat *float64 `xml:"insuranceCostVat,omitempty"`
	CodAmount        *float64 `xml:"codAmount,omitempty"`
}

type operationAddParcelsResponse struct {
	AddParcelsResponse *addParcelsResponse `xml:"addParcelsResponse,omitempty"`
}

type addParcelsResponse struct {
	Return *OrderCorrectionStatus `xml:"return,omitempty"`
}

//OrderCorrectionStatus AddParcels\RemoveParcels response body
type OrderCorrectionStatus struct {
	OrderNum     *string         `xml:"orderNum,omitempty"`
	Status       *string         `xml:"status,omitempty"`
	ErrorMessage *string         `xml:"errorMessage,omitempty"`
	ParcelStatus []*ParcelStatus `xml:"parcelStatus,omitempty"`
}

//ParcelsStatus correction status
type ParcelStatus struct {
	Number       *string `xml:"number,omitempty"`
	Status       *string `xml:"status,omitempty"`
	ErrorMessage *string `xml:"errorMessage,omitempty"`
}

type operationCancelOrder struct {
	CancelOrder *cancelOrder `xml:"cancelOrder,omitempty"`
}

type cancelOrder struct {
	Orders *CancelOrderRequest `xml:"orders,omitempty"`
}

//CancelOrderRequest CancelOrder request body
type CancelOrderRequest struct {
	Auth   *Auth            `xml:"auth,omitempty"`
	Cancel []*OrderToCancel `xml:"cancel,omitempty"`
}

//OrderToCancel ...
type OrderToCancel struct {
	OrderNumberInternal *string `xml:"orderNumberInternal,omitempty"`
	OrderNum            *string `xml:"orderNum,omitempty"`
	PickupDate          *string `xml:"pickupdate,omitempty"`
}

type operationCancelOrderResponse struct {
	CancelOrderResponse *cancelOrderResponse `xml:"cancelOrderResponse,omitempty"`
}

type cancelOrderResponse struct {
	Return []*OrderStatus `xml:"return,omitempty"`
}

//OrderStatus CancelOrder response item
type OrderStatus struct {
	OrderNumberInternal *string `xml:"orderNumberInternal,omitempty"`
	OrderNum            *string `xml:"orderNum,omitempty"`
	Status              *string `xml:"status,omitempty"`
	ErrorMessage        *string `xml:"errorMessage,omitempty"`
}

type operationCreateOrder struct {
	CreateOrder *createOrder `xml:"createOrder,omitempty"`
}

type createOrder struct {
	Orders *CreateOrderRequest `xml:"orders,omitempty"`
}

//CreateOrderRequest CreateOrder request body
type CreateOrderRequest struct {
	Auth   *Auth    `xml:"auth,omitempty"`
	Header *Header  `xml:"header,omitempty"`
	Order  []*Order `xml:"order,omitempty"`
}

//Order to creation
type Order struct {
	OrderNumberInternal *string            `xml:"orderNumberInternal,omitempty"`
	ServiceCode         *string            `xml:"serviceCode,omitempty"`
	ServiceVariant      *string            `xml:"serviceVariant,omitempty"`
	CargoNumPack        *int               `xml:"cargoNumPack,omitempty"`
	CargoWeight         *float64           `xml:"cargoWeight,omitempty"`
	CargoVolume         *float64           `xml:"cargoVolume,omitempty"`
	CargoRegistered     *bool              `xml:"cargoRegistered,omitempty"`
	CargoValue          *float64           `xml:"cargoValue,omitempty"`
	CargoCategory       *string            `xml:"cargoCategory,omitempty"`
	DeliveryTimePeriod  *string            `xml:"deliveryTimePeriod,omitempty"`
	PaymentType         *string            `xml:"paymentType,omitempty"`
	ExtraParam          []*ExtraParameter  `xml:"extraParam,omitempty"`
	DataInt             *DataInternational `xml:"dataInt,omitempty"`
	ReceiverAddress     *Address           `xml:"receiverAddress,omitempty"`
	ReturnAddress       *Address           `xml:"returnAddress,omitempty"`
	ExtraService        []*ExtraService    `xml:"extraService,omitempty"`
	Parcel              []*Parcel          `xml:"parcel,omitempty"`
	UnitLoad            []*UnitLoadOrder   `xml:"unitLoad,omitempty"`
}

//UnitLoadOrder array of units in cargo for 54 federal law (54-ФЗ)
type UnitLoadOrder struct {
	commonUnitLoad
	ClientCode *string `xml:"client_code,omitempty"`
}

//DataInternational data for international delivery orders
type DataInternational struct {
	Currency              *string `xml:"currency,omitempty"`
	CurrencyDeclaredValue *string `xml:"currencyDeclaredValue,omitempty"`
}

//ExtraService list of delivery options
type ExtraService struct {
	EsCode *string           `xml:"esCode,omitempty"`
	Param  []*ExtraParameter `xml:"param,omitempty"`
}

//ExtraParameter delivery option
type ExtraParameter struct {
	Name  *string `xml:"name,omitempty"`
	Value *string `xml:"value,omitempty"`
}

//Header common for all types of orders. Contains information about sender
type Header struct {
	DatePickup       *string  `xml:"datePickup,omitempty"`
	Payer            *int64   `xml:"payer,omitempty"`
	SenderAddress    *Address `xml:"senderAddress,omitempty"`
	PickupTimePeriod *string  `xml:"pickupTimePeriod,omitempty"`
	RegularNum       *string  `xml:"regularNum,omitempty"`
}

//Address ...
type Address struct {
	Code           *string `xml:"code,omitempty"`
	Name           *string `xml:"name,omitempty"`
	TerminalCode   *string `xml:"terminalCode,omitempty"`
	AddressString  *string `xml:"addressString,omitempty"`
	CountryName    *string `xml:"countryName,omitempty"`
	Index          *string `xml:"index,omitempty"`
	Region         *string `xml:"region,omitempty"`
	City           *string `xml:"city,omitempty"`
	Street         *string `xml:"street,omitempty"`
	StreetAbbr     *string `xml:"streetAbbr,omitempty"`
	House          *string `xml:"house,omitempty"`
	HouseKorpus    *string `xml:"houseKorpus,omitempty"`
	Str            *string `xml:"str,omitempty"`
	Vlad           *string `xml:"vlad,omitempty"`
	ExtraInfo      *string `xml:"extraInfo,omitempty"`
	Office         *string `xml:"office,omitempty"`
	Flat           *string `xml:"flat,omitempty"`
	WorkTimeFrom   *string `xml:"workTimeFrom,omitempty"`
	WorkTimeTo     *string `xml:"workTimeTo,omitempty"`
	DinnerTimeFrom *string `xml:"dinnerTimeFrom,omitempty"`
	DinnerTimeTo   *string `xml:"dinnerTimeTo,omitempty"`
	ContactFio     *string `xml:"contactFio,omitempty"`
	ContactPhone   *string `xml:"contactPhone,omitempty"`
	ContactEmail   *string `xml:"contactEmail,omitempty"`
	Instructions   *string `xml:"instructions,omitempty"`
	NeedPass       *bool   `xml:"needPass,omitempty"`
}

type operationCreateOrderResponse struct {
	CreateOrderResponse *createOrderResponse `xml:"createOrderResponse,omitempty"`
}

type createOrderResponse struct {
	Return []*OrderStatus `xml:"return,omitempty"`
}

type operationGetOrderStatus struct {
	GetOrderStatus *getOrderStatus `xml:"getOrderStatus,omitempty"`
}

type getOrderStatus struct {
	OrderStatus *getOrderStatusRequest `xml:"orderStatus,omitempty"`
}

type getOrderStatusRequest struct {
	Auth  *Auth                  `xml:"auth,omitempty"`
	Order []*InternalOrderNumber `xml:"order,omitempty"`
}

//InternalOrderNumber GetOrderStatus request body. Imply client side oder number
type InternalOrderNumber struct {
	OrderNumberInternal *string `xml:"orderNumberInternal,omitempty"`
	DatePickup          *string `xml:"datePickup,omitempty"`
}

type operationGetOrderStatusResponse struct {
	GetOrderStatusResponse *getOrderStatusResponse `xml:"getOrderStatusResponse,omitempty"`
}

type getOrderStatusResponse struct {
	Return []*OrderStatus `xml:"return,omitempty"`
}

type operationRemoveParcels struct {
	RemoveParcels *removeParcels `xml:"removeParcels,omitempty"`
}

type removeParcels struct {
	Parcels *UpdateOrderRequest `xml:"parcels,omitempty"`
}

type operationRemoveParcelsResponse struct {
	RemoveParcelsResponse *removeParcelsResponse `xml:"removeParcelsResponse,omitempty"`
}

type removeParcelsResponse struct {
	Return *OrderCorrectionStatus `xml:"return,omitempty"`
}
