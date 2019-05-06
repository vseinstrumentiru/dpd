package dpdlib

import "github.com/fiorix/wsdl2go/soap"

var OrderNamespace = "http://dpd.ru/ws/order2/2012-04-04"

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
