package dpd_soap

import "github.com/fiorix/wsdl2go/soap"

const OrderNamespace = "http://dpd.ru/ws/order2/2012-04-04"

func NewDPDOrder(cli *soap.Client) DPDOrder {
	return &dPDOrder{cli}
}

type DPDOrder interface {
	AddAirwayBill(AddAirwayBill *AddAirwayBill) (*AddAirwayBillResponse, error)
	AddParcels(AddParcels *AddParcels) (*AddParcelsResponse, error)
	CancelOrder(CancelOrder *CancelOrder) (*CancelOrderResponse, error)
	ChangeDeliveryDate(ChangeDeliveryDate *ChangeDeliveryDate) (*ChangeDeliveryDateResponse, error)
	ChangeUnitLoad(ChangeUnitLoad *ChangeUnitLoad) (*ChangeUnitLoadResponse, error)
	CreateAddress(CreateAddress *CreateAddress) (*CreateAddressResponse, error)
	CreateOrder(CreateOrder *CreateOrder) (*CreateOrderResponse, error)
	CreateOrderAsync(CreateOrderAsync *CreateOrderAsync) (*CreateOrderAsyncResponse, error)
	CreateOrderResult(CreateOrderResult *CreateOrderResult) (*CreateOrderResultResponse, error)
	GetInvoiceFile(GetInvoiceFile *GetInvoiceFile) (*GetInvoiceFileResponse, error)
	GetOrderSMS(GetOrderSMS *GetOrderSMS) (*GetOrderSMSResponse, error)
	GetOrderStatus(GetOrderStatus *GetOrderStatus) (*GetOrderStatusResponse, error)
	GetRegisterFile(GetRegisterFile *GetRegisterFile) (*GetRegisterFileResponse, error)
	OrderReDirect(OrderReDirect *OrderReDirect) (*OrderReDirectResponse, error)
	RemoveParcels(RemoveParcels *RemoveParcels) (*RemoveParcelsResponse, error)
	UpdateAddress(UpdateAddress *UpdateAddress) (*UpdateAddressResponse, error)
}

func (p *dPDOrder) AddAirwayBill(AddAirwayBill *AddAirwayBill) (*AddAirwayBillResponse, error) {
	α := struct {
		OperationAddAirwayBill `xml:"tns:addAirwayBill"`
	}{
		OperationAddAirwayBill{
			AddAirwayBill,
		},
	}

	γ := struct {
		OperationAddAirwayBillResponse `xml:"addAirwayBillResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("AddAirwayBill", α, &γ); err != nil {
		return nil, err
	}
	return γ.AddAirwayBillResponse, nil
}

func (p *dPDOrder) AddParcels(AddParcels *AddParcels) (*AddParcelsResponse, error) {
	α := struct {
		OperationAddParcels `xml:"tns:addParcels"`
	}{
		OperationAddParcels{
			AddParcels,
		},
	}

	γ := struct {
		OperationAddParcelsResponse `xml:"addParcelsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("AddParcels", α, &γ); err != nil {
		return nil, err
	}
	return γ.AddParcelsResponse, nil
}

func (p *dPDOrder) CancelOrder(CancelOrder *CancelOrder) (*CancelOrderResponse, error) {
	α := struct {
		OperationCancelOrder `xml:"tns:cancelOrder"`
	}{
		OperationCancelOrder{
			CancelOrder,
		},
	}

	γ := struct {
		OperationCancelOrderResponse `xml:"cancelOrderResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CancelOrder", α, &γ); err != nil {
		return nil, err
	}
	return γ.CancelOrderResponse, nil
}

func (p *dPDOrder) ChangeDeliveryDate(ChangeDeliveryDate *ChangeDeliveryDate) (*ChangeDeliveryDateResponse, error) {
	α := struct {
		OperationChangeDeliveryDate `xml:"tns:changeDeliveryDate"`
	}{
		OperationChangeDeliveryDate{
			ChangeDeliveryDate,
		},
	}

	γ := struct {
		OperationChangeDeliveryDateResponse `xml:"changeDeliveryDateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ChangeDeliveryDate", α, &γ); err != nil {
		return nil, err
	}
	return γ.ChangeDeliveryDateResponse, nil
}

func (p *dPDOrder) ChangeUnitLoad(ChangeUnitLoad *ChangeUnitLoad) (*ChangeUnitLoadResponse, error) {
	α := struct {
		OperationChangeUnitLoad `xml:"tns:changeUnitLoad"`
	}{
		OperationChangeUnitLoad{
			ChangeUnitLoad,
		},
	}

	γ := struct {
		OperationChangeUnitLoadResponse `xml:"changeUnitLoadResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ChangeUnitLoad", α, &γ); err != nil {
		return nil, err
	}
	return γ.ChangeUnitLoadResponse, nil
}

func (p *dPDOrder) CreateAddress(CreateAddress *CreateAddress) (*CreateAddressResponse, error) {
	α := struct {
		OperationCreateAddress `xml:"tns:createAddress"`
	}{
		OperationCreateAddress{
			CreateAddress,
		},
	}

	γ := struct {
		OperationCreateAddressResponse `xml:"createAddressResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateAddress", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateAddressResponse, nil
}

func (p *dPDOrder) CreateOrder(CreateOrder *CreateOrder) (*CreateOrderResponse, error) {
	α := struct {
		OperationCreateOrder `xml:"tns:createOrder"`
	}{
		OperationCreateOrder{
			CreateOrder,
		},
	}

	γ := struct {
		OperationCreateOrderResponse `xml:"createOrderResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateOrder", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateOrderResponse, nil
}

func (p *dPDOrder) CreateOrderAsync(CreateOrderAsync *CreateOrderAsync) (*CreateOrderAsyncResponse, error) {
	α := struct {
		OperationCreateOrderAsync `xml:"tns:createOrderAsync"`
	}{
		OperationCreateOrderAsync{
			CreateOrderAsync,
		},
	}

	γ := struct {
		OperationCreateOrderAsyncResponse `xml:"createOrderAsyncResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateOrderAsync", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateOrderAsyncResponse, nil
}

func (p *dPDOrder) CreateOrderResult(CreateOrderResult *CreateOrderResult) (*CreateOrderResultResponse, error) {
	α := struct {
		OperationCreateOrderResult `xml:"tns:createOrderResult"`
	}{
		OperationCreateOrderResult{
			CreateOrderResult,
		},
	}

	γ := struct {
		OperationCreateOrderResultResponse `xml:"createOrderResultResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("CreateOrderResult", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateOrderResultResponse, nil
}

func (p *dPDOrder) GetInvoiceFile(GetInvoiceFile *GetInvoiceFile) (*GetInvoiceFileResponse, error) {
	α := struct {
		OperationGetInvoiceFile `xml:"tns:getInvoiceFile"`
	}{
		OperationGetInvoiceFile{
			GetInvoiceFile,
		},
	}

	γ := struct {
		OperationGetInvoiceFileResponse `xml:"getInvoiceFileResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetInvoiceFile", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetInvoiceFileResponse, nil
}

func (p *dPDOrder) GetOrderSMS(GetOrderSMS *GetOrderSMS) (*GetOrderSMSResponse, error) {
	α := struct {
		OperationGetOrderSMS `xml:"tns:getOrderSMS"`
	}{
		OperationGetOrderSMS{
			GetOrderSMS,
		},
	}

	γ := struct {
		OperationGetOrderSMSResponse `xml:"getOrderSMSResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetOrderSMS", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetOrderSMSResponse, nil
}

func (p *dPDOrder) GetOrderStatus(GetOrderStatus *GetOrderStatus) (*GetOrderStatusResponse, error) {
	α := struct {
		OperationGetOrderStatus `xml:"tns:getOrderStatus"`
	}{
		OperationGetOrderStatus{
			GetOrderStatus,
		},
	}

	γ := struct {
		OperationGetOrderStatusResponse `xml:"getOrderStatusResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetOrderStatus", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetOrderStatusResponse, nil
}

func (p *dPDOrder) GetRegisterFile(GetRegisterFile *GetRegisterFile) (*GetRegisterFileResponse, error) {
	α := struct {
		OperationGetRegisterFile `xml:"tns:getRegisterFile"`
	}{
		OperationGetRegisterFile{
			GetRegisterFile,
		},
	}

	γ := struct {
		OperationGetRegisterFileResponse `xml:"getRegisterFileResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("GetRegisterFile", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetRegisterFileResponse, nil
}

func (p *dPDOrder) OrderReDirect(OrderReDirect *OrderReDirect) (*OrderReDirectResponse, error) {
	α := struct {
		OperationOrderReDirect `xml:"tns:orderReDirect"`
	}{
		OperationOrderReDirect{
			OrderReDirect,
		},
	}

	γ := struct {
		OperationOrderReDirectResponse `xml:"orderReDirectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("OrderReDirect", α, &γ); err != nil {
		return nil, err
	}
	return γ.OrderReDirectResponse, nil
}

func (p *dPDOrder) RemoveParcels(RemoveParcels *RemoveParcels) (*RemoveParcelsResponse, error) {
	α := struct {
		OperationRemoveParcels `xml:"tns:removeParcels"`
	}{
		OperationRemoveParcels{
			RemoveParcels,
		},
	}

	γ := struct {
		OperationRemoveParcelsResponse `xml:"removeParcelsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("RemoveParcels", α, &γ); err != nil {
		return nil, err
	}
	return γ.RemoveParcelsResponse, nil
}

func (p *dPDOrder) UpdateAddress(UpdateAddress *UpdateAddress) (*UpdateAddressResponse, error) {
	α := struct {
		OperationUpdateAddress `xml:"tns:updateAddress"`
	}{
		OperationUpdateAddress{
			UpdateAddress,
		},
	}

	γ := struct {
		OperationUpdateAddressResponse `xml:"updateAddressResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("UpdateAddress", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateAddressResponse, nil
}

type AddAirwayBill struct {
	Request *DpdAirwayBillInput `xml:"request,omitempty"`
}

type AddAirwayBillResponse struct {
	Return *DpdAirwayBillOutput `xml:"return,omitempty"`
}

type AddParcels struct {
	Parcels *DpdOrderCorrection `xml:"parcels,omitempty"`
}

type AddParcelsResponse struct {
	Return *DpdOrderCorrectionStatus `xml:"return,omitempty"`
}

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

type AsyncRequest struct {
	Auth   *Auth  `xml:"auth,omitempty"`
	TaskId *int64 `xml:"taskId,omitempty"`
}

type AsyncResponse struct {
	TaskId   *int64    `xml:"taskId,omitempty"`
	TaskDate *DateTime `xml:"taskDate,omitempty"`
}

type CancelOrder struct {
	Orders *DpdOrderCancellation `xml:"orders,omitempty"`
}

type CancelOrderResponse struct {
	Return []*DpdOrderStatus `xml:"return,omitempty"`
}

type ChangeDeliveryDate struct {
	Request *DpdChangeDeliveryDate `xml:"request,omitempty"`
}

type ChangeDeliveryDateResponse struct {
	Return []*DpdOrderStatus `xml:"return,omitempty"`
}

type ChangeUnitLoad struct {
	Request *DpdUnitLoadChange `xml:"request,omitempty"`
}

type ChangeUnitLoadResponse struct {
	Return []*DpdOrderStatus `xml:"return,omitempty"`
}

type ClientAddress struct {
	Code           *string `xml:"code,omitempty"`
	Name           *string `xml:"name,omitempty"`
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

type CreateAddress struct {
	Address *DpdClientAddress `xml:"address,omitempty"`
}

type CreateAddressResponse struct {
	Return *DpdClientAddressStatus `xml:"return,omitempty"`
}

type CreateOrder struct {
	Orders *DpdOrdersData `xml:"orders,omitempty"`
}

type CreateOrderAsync struct {
	Orders *DpdOrdersData `xml:"orders,omitempty"`
}

type CreateOrderAsyncResponse struct {
	Return *AsyncResponse `xml:"return,omitempty"`
}

type CreateOrderResponse struct {
	Return []*DpdOrderStatus `xml:"return,omitempty"`
}

type CreateOrderResult struct {
	Orders *AsyncRequest `xml:"orders,omitempty"`
}

type CreateOrderResultResponse struct {
	Return []*DpdOrderStatus `xml:"return,omitempty"`
}

type DataInternational struct {
	Currency              *string `xml:"currency,omitempty"`
	CurrencyDeclaredValue *string `xml:"currencyDeclaredValue,omitempty"`
}

type DpdAirwayBillInput struct {
	Auth  *Auth                `xml:"auth,omitempty"`
	Order []*OrderAirbillInput `xml:"order,omitempty"`
}

type DpdAirwayBillOutput struct {
	Order []*OrderAirbillOutput `xml:"order,omitempty"`
}

type DpdChangeDeliveryDate struct {
	Auth  *Auth                `xml:"auth,omitempty"`
	Order []*OrderDeliveryDate `xml:"order,omitempty"`
}

type DpdClientAddress struct {
	Auth          *Auth          `xml:"auth,omitempty"`
	ClientAddress *ClientAddress `xml:"clientAddress,omitempty"`
}

type DpdClientAddressStatus struct {
	Code         *string `xml:"code,omitempty"`
	Status       *string `xml:"status,omitempty"`
	ErrorMessage *string `xml:"errorMessage,omitempty"`
}

type DpdGetInvoiceFile struct {
	Auth        *Auth    `xml:"auth,omitempty"`
	OrderNum    *string  `xml:"orderNum,omitempty"`
	ParcelCount *int     `xml:"parcelCount,omitempty"`
	CargoValue  *float64 `xml:"cargoValue,omitempty"`
}

type DpdGetOrderStatus struct {
	Auth  *Auth                  `xml:"auth,omitempty"`
	Order []*InternalOrderNumber `xml:"order,omitempty"`
}

type DpdGetRegisterFile struct {
	Auth         *Auth   `xml:"auth,omitempty"`
	DatePickup   *Date   `xml:"datePickup,omitempty"`
	RegularNum   *string `xml:"regularNum,omitempty"`
	CityPickupId *int64  `xml:"cityPickupId,omitempty"`
	AddressCode  *string `xml:"addressCode,omitempty"`
}

type DpdInvoiceFile struct {
	File *[]byte `xml:"file,omitempty"`
}

type DpdOrderCancellation struct {
	Auth   *Auth          `xml:"auth,omitempty"`
	Cancel []*OrderCancel `xml:"cancel,omitempty"`
}

type DpdOrderCorrection struct {
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

type DpdOrderCorrectionStatus struct {
	OrderNum     *string         `xml:"orderNum,omitempty"`
	Status       *string         `xml:"status,omitempty"`
	ErrorMessage *string         `xml:"errorMessage,omitempty"`
	ParcelStatus []*ParcelStatus `xml:"parcelStatus,omitempty"`
}

type DpdOrderResult struct {
	OrderNum   *string `xml:"orderNum,omitempty"`
	StatusCode *string `xml:"statusCode,omitempty"`
}

type DpdOrderSmsRequest struct {
	Auth  *Auth       `xml:"auth,omitempty"`
	Order []*OrderNum `xml:"order,omitempty"`
}

type DpdOrderSmsResult struct {
	Order []*DpdOrderSmsStatus `xml:"order,omitempty"`
}

type DpdOrderSmsStatus struct {
	OrderNum     *string `xml:"orderNum,omitempty"`
	Status       *string `xml:"status,omitempty"`
	ErrorMessage *string `xml:"errorMessage,omitempty"`
}

type DpdOrderStatus struct {
	OrderNumberInternal *string `xml:"orderNumberInternal,omitempty"`
	OrderNum            *string `xml:"orderNum,omitempty"`
	Status              *string `xml:"status,omitempty"`
	ErrorMessage        *string `xml:"errorMessage,omitempty"`
}

type DpdOrdersData struct {
	Auth   *Auth    `xml:"auth,omitempty"`
	Header *Header  `xml:"header,omitempty"`
	Order  []*Order `xml:"order,omitempty"`
}

type DpdOrdersReDirect struct {
	Auth  *Auth            `xml:"auth,omitempty"`
	Order []*ReDirectOrder `xml:"order,omitempty"`
}

type DpdRegisterFile struct {
	File []*[]byte `xml:"file,omitempty"`
}

type DpdUnitLoadChange struct {
	Auth  *Auth            `xml:"auth,omitempty"`
	Order []*OrderUnitLoad `xml:"order,omitempty"`
}

type OrderExtraService struct {
	EsCode *string      `xml:"esCode,omitempty"`
	Param  []*Parameter `xml:"param,omitempty"`
}

type GetInvoiceFile struct {
	Request *DpdGetInvoiceFile `xml:"request,omitempty"`
}

type GetInvoiceFileResponse struct {
	Return *DpdInvoiceFile `xml:"return,omitempty"`
}

type GetOrderSMS struct {
	Request *DpdOrderSmsRequest `xml:"request,omitempty"`
}

type GetOrderSMSResponse struct {
	Return *DpdOrderSmsResult `xml:"return,omitempty"`
}

type GetOrderStatus struct {
	OrderStatus *DpdGetOrderStatus `xml:"orderStatus,omitempty"`
}

type GetOrderStatusResponse struct {
	Return []*DpdOrderStatus `xml:"return,omitempty"`
}

type GetRegisterFile struct {
	Request *DpdGetRegisterFile `xml:"request,omitempty"`
}

type GetRegisterFileResponse struct {
	Return *DpdRegisterFile `xml:"return,omitempty"`
}

type Header struct {
	DatePickup       *Date    `xml:"datePickup,omitempty"`
	Payer            *int64   `xml:"payer,omitempty"`
	SenderAddress    *Address `xml:"senderAddress,omitempty"`
	PickupTimePeriod *string  `xml:"pickupTimePeriod,omitempty"`
	RegularNum       *string  `xml:"regularNum,omitempty"`
}

type InternalOrderNumber struct {
	OrderNumberInternal *string `xml:"orderNumberInternal,omitempty"`
	DatePickup          *Date   `xml:"datePickup,omitempty"`
}

type Order struct {
	OrderNumberInternal *string              `xml:"orderNumberInternal,omitempty"`
	ServiceCode         *string              `xml:"serviceCode,omitempty"`
	ServiceVariant      *string              `xml:"serviceVariant,omitempty"`
	CargoNumPack        *int                 `xml:"cargoNumPack,omitempty"`
	CargoWeight         *float64             `xml:"cargoWeight,omitempty"`
	CargoVolume         *float64             `xml:"cargoVolume,omitempty"`
	CargoRegistered     *bool                `xml:"cargoRegistered,omitempty"`
	CargoValue          *float64             `xml:"cargoValue,omitempty"`
	CargoCategory       *string              `xml:"cargoCategory,omitempty"`
	DeliveryTimePeriod  *string              `xml:"deliveryTimePeriod,omitempty"`
	PaymentType         *string              `xml:"paymentType,omitempty"`
	ExtraParam          []*Parameter         `xml:"extraParam,omitempty"`
	DataInt             *DataInternational   `xml:"dataInt,omitempty"`
	ReceiverAddress     *Address             `xml:"receiverAddress,omitempty"`
	ReturnAddress       *Address             `xml:"returnAddress,omitempty"`
	ExtraService        []*OrderExtraService `xml:"extraService,omitempty"`
	Parcel              []*Parcel            `xml:"parcel,omitempty"`
	UnitLoad            []*UnitLoad          `xml:"unitLoad,omitempty"`
}

type OrderAirbillInput struct {
	OrderNumberDPD      *string         `xml:"orderNumberDPD,omitempty"`
	OrderNumberInternal *string         `xml:"orderNumberInternal,omitempty"`
	Param               []*ParamAirbill `xml:"param,omitempty"`
}

type OrderAirbillOutput struct {
	OrderNumberDPD *string `xml:"orderNumberDPD,omitempty"`
	Result         *string `xml:"result,omitempty"`
	ErrorCode      *string `xml:"errorCode,omitempty"`
	ErrorMessage   *string `xml:"errorMessage,omitempty"`
}

type OrderCancel struct {
	OrderNumberInternal *string `xml:"orderNumberInternal,omitempty"`
	OrderNum            *string `xml:"orderNum,omitempty"`
	Pickupdate          *Date   `xml:"pickupdate,omitempty"`
}

type OrderDeliveryDate struct {
	OrderNum        *string `xml:"orderNum,omitempty"`
	NewDeliveryDate *Date   `xml:"newDeliveryDate,omitempty"`
}

type OrderNum struct {
	OrderNum *string `xml:"orderNum,omitempty"`
}

type OrderReDirect struct {
	Orders *DpdOrdersReDirect `xml:"orders,omitempty"`
}

type OrderReDirectResponse struct {
	Return []*DpdOrderResult `xml:"return,omitempty"`
}

type OrderUnitLoad struct {
	OrderNum          *string     `xml:"orderNum,omitempty"`
	DatePickup        *Date       `xml:"datePickup,omitempty"`
	DeliveryAmount    *float64    `xml:"deliveryAmount,omitempty"`
	GoodsReturnAmount *float64    `xml:"goodsReturnAmount,omitempty"`
	UnitLoad          []*UnitLoad `xml:"unitLoad,omitempty"`
}

type ParamAirbill struct {
	ParamName  *string `xml:"paramName,omitempty"`
	ParamValue *string `xml:"paramValue,omitempty"`
}

type Parameter struct {
	Name  *string `xml:"name,omitempty"`
	Value *string `xml:"value,omitempty"`
}

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

type ParcelStatus struct {
	Number       *string `xml:"number,omitempty"`
	Status       *string `xml:"status,omitempty"`
	ErrorMessage *string `xml:"errorMessage,omitempty"`
}

type ReDirectOrder struct {
	OrderNum        *string `xml:"orderNum,omitempty"`
	DatePickup      *Date   `xml:"datePickup,omitempty"`
	ReDirectionType *string `xml:"reDirectionType,omitempty"`
	PointCode       *string `xml:"pointCode,omitempty"`
}

type RemoveParcels struct {
	Parcels *DpdOrderCorrection `xml:"parcels,omitempty"`
}

type RemoveParcelsResponse struct {
	Return *DpdOrderCorrectionStatus `xml:"return,omitempty"`
}

type UnitLoad struct {
	Article       *string `xml:"article,omitempty"`
	Descript      *string `xml:"descript,omitempty"`
	ClientCode    *string `xml:"client_code,omitempty"`
	DeclaredValue *string `xml:"declared_value,omitempty"`
	ParcelNum     *string `xml:"parcel_num,omitempty"`
	NppAmount     *string `xml:"npp_amount,omitempty"`
	VatPercent    *int    `xml:"vat_percent,omitempty"`
	WithoutVat    *bool   `xml:"without_vat,omitempty"`
	Count         *int    `xml:"count,omitempty"`
}

type UpdateAddress struct {
	Address *DpdClientAddress `xml:"address,omitempty"`
}

type UpdateAddressResponse struct {
	Return *DpdClientAddressStatus `xml:"return,omitempty"`
}

type OperationAddAirwayBill struct {
	AddAirwayBill *AddAirwayBill `xml:"addAirwayBill,omitempty"`
}

type OperationAddAirwayBillResponse struct {
	AddAirwayBillResponse *AddAirwayBillResponse `xml:"addAirwayBillResponse,omitempty"`
}

type OperationAddParcels struct {
	AddParcels *AddParcels `xml:"addParcels,omitempty"`
}

type OperationAddParcelsResponse struct {
	AddParcelsResponse *AddParcelsResponse `xml:"addParcelsResponse,omitempty"`
}

type OperationCancelOrder struct {
	CancelOrder *CancelOrder `xml:"cancelOrder,omitempty"`
}

type OperationCancelOrderResponse struct {
	CancelOrderResponse *CancelOrderResponse `xml:"cancelOrderResponse,omitempty"`
}

type OperationChangeDeliveryDate struct {
	ChangeDeliveryDate *ChangeDeliveryDate `xml:"changeDeliveryDate,omitempty"`
}

type OperationChangeDeliveryDateResponse struct {
	ChangeDeliveryDateResponse *ChangeDeliveryDateResponse `xml:"changeDeliveryDateResponse,omitempty"`
}

type OperationChangeUnitLoad struct {
	ChangeUnitLoad *ChangeUnitLoad `xml:"changeUnitLoad,omitempty"`
}

type OperationChangeUnitLoadResponse struct {
	ChangeUnitLoadResponse *ChangeUnitLoadResponse `xml:"changeUnitLoadResponse,omitempty"`
}

type OperationCreateAddress struct {
	CreateAddress *CreateAddress `xml:"createAddress,omitempty"`
}

type OperationCreateAddressResponse struct {
	CreateAddressResponse *CreateAddressResponse `xml:"createAddressResponse,omitempty"`
}

type OperationCreateOrder struct {
	CreateOrder *CreateOrder `xml:"createOrder,omitempty"`
}

type OperationCreateOrderResponse struct {
	CreateOrderResponse *CreateOrderResponse `xml:"createOrderResponse,omitempty"`
}

type OperationCreateOrderAsync struct {
	CreateOrderAsync *CreateOrderAsync `xml:"createOrderAsync,omitempty"`
}

type OperationCreateOrderAsyncResponse struct {
	CreateOrderAsyncResponse *CreateOrderAsyncResponse `xml:"createOrderAsyncResponse,omitempty"`
}

type OperationCreateOrderResult struct {
	CreateOrderResult *CreateOrderResult `xml:"createOrderResult,omitempty"`
}

type OperationCreateOrderResultResponse struct {
	CreateOrderResultResponse *CreateOrderResultResponse `xml:"createOrderResultResponse,omitempty"`
}

type OperationGetInvoiceFile struct {
	GetInvoiceFile *GetInvoiceFile `xml:"getInvoiceFile,omitempty"`
}

type OperationGetInvoiceFileResponse struct {
	GetInvoiceFileResponse *GetInvoiceFileResponse `xml:"getInvoiceFileResponse,omitempty"`
}

type OperationGetOrderSMS struct {
	GetOrderSMS *GetOrderSMS `xml:"getOrderSMS,omitempty"`
}

type OperationGetOrderSMSResponse struct {
	GetOrderSMSResponse *GetOrderSMSResponse `xml:"getOrderSMSResponse,omitempty"`
}

type OperationGetOrderStatus struct {
	GetOrderStatus *GetOrderStatus `xml:"getOrderStatus,omitempty"`
}

type OperationGetOrderStatusResponse struct {
	GetOrderStatusResponse *GetOrderStatusResponse `xml:"getOrderStatusResponse,omitempty"`
}

type OperationGetRegisterFile struct {
	GetRegisterFile *GetRegisterFile `xml:"getRegisterFile,omitempty"`
}

type OperationGetRegisterFileResponse struct {
	GetRegisterFileResponse *GetRegisterFileResponse `xml:"getRegisterFileResponse,omitempty"`
}

type OperationOrderReDirect struct {
	OrderReDirect *OrderReDirect `xml:"orderReDirect,omitempty"`
}

type OperationOrderReDirectResponse struct {
	OrderReDirectResponse *OrderReDirectResponse `xml:"orderReDirectResponse,omitempty"`
}

type OperationRemoveParcels struct {
	RemoveParcels *RemoveParcels `xml:"removeParcels,omitempty"`
}

type OperationRemoveParcelsResponse struct {
	RemoveParcelsResponse *RemoveParcelsResponse `xml:"removeParcelsResponse,omitempty"`
}

type OperationUpdateAddress struct {
	UpdateAddress *UpdateAddress `xml:"updateAddress,omitempty"`
}

type OperationUpdateAddressResponse struct {
	UpdateAddressResponse *UpdateAddressResponse `xml:"updateAddressResponse,omitempty"`
}

type dPDOrder struct {
	cli *soap.Client
}
