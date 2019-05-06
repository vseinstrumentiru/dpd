package dpdlib



type Confirm struct {
	Request *RequestConfirm `xml:"request,omitempty"`
}

type ConfirmResponse struct {
	Return *string `xml:"return,omitempty"`
}

type GetStatesByClient struct {
	Request *RequestClient `xml:"request,omitempty"`
}

type GetStatesByClientOrder struct {
	Request *RequestClientOrder `xml:"request,omitempty"`
}

type GetStatesByClientOrderResponse struct {
	Return *StateParcels `xml:"return,omitempty"`
}

type GetStatesByClientParcel struct {
	Request *RequestClientParcel `xml:"request,omitempty"`
}

type GetStatesByClientParcelResponse struct {
	Return *StateParcels `xml:"return,omitempty"`
}

type GetStatesByClientResponse struct {
	Return *StateParcels `xml:"return,omitempty"`
}

type GetStatesByDPDOrder struct {
	Request *RequestDpdOrder `xml:"request,omitempty"`
}

type GetStatesByDPDOrderResponse struct {
	Return *StateParcels `xml:"return,omitempty"`
}

type RequestClient struct {
	Auth *Auth `xml:"auth,omitempty"`
}

type RequestClientOrder struct {
	Auth          *Auth   `xml:"auth,omitempty"`
	ClientOrderNr *string `xml:"clientOrderNr,omitempty"`
	PickupDate    *Date   `xml:"pickupDate,omitempty"`
}

type RequestClientParcel struct {
	Auth           *Auth   `xml:"auth,omitempty"`
	ClientParcelNr *string `xml:"clientParcelNr,omitempty"`
	PickupDate     *Date   `xml:"pickupDate,omitempty"`
}

type RequestConfirm struct {
	Auth  *Auth  `xml:"auth,omitempty"`
	DocId *int64 `xml:"docId,omitempty"`
}

type RequestDpdOrder struct {
	Auth       *Auth   `xml:"auth,omitempty"`
	DpdOrderNr *string `xml:"dpdOrderNr,omitempty"`
	PickupYear *int    `xml:"pickupYear,omitempty"`
}

type StateParcel struct {
	ClientOrderNr    *string   `xml:"clientOrderNr,omitempty"`
	ClientParcelNr   *string   `xml:"clientParcelNr,omitempty"`
	DpdOrderNr       *string   `xml:"dpdOrderNr,omitempty"`
	DpdParcelNr      *string   `xml:"dpdParcelNr,omitempty"`
	PickupDate       *Date     `xml:"pickupDate,omitempty"`
	DpdOrderReNr     *string   `xml:"dpdOrderReNr,omitempty"`
	DpdParcelReNr    *string   `xml:"dpdParcelReNr,omitempty"`
	IsReturn         *bool     `xml:"isReturn,omitempty"`
	PlanDeliveryDate *Date     `xml:"planDeliveryDate,omitempty"`
	NewState         *string   `xml:"newState,omitempty"`
	TransitionTime   *DateTime `xml:"transitionTime,omitempty"`
	TerminalCode     *string   `xml:"terminalCode,omitempty"`
	TerminalCity     *string   `xml:"terminalCity,omitempty"`
	IncidentCode     *string   `xml:"incidentCode,omitempty"`
	IncidentName     *string   `xml:"incidentName,omitempty"`
	Consignee        *string   `xml:"consignee,omitempty"`
}

type StateParcels struct {
	DocId          *int64         `xml:"docId,omitempty"`
	DocDate        *DateTime      `xml:"docDate,omitempty"`
	ClientNumber   *int64         `xml:"clientNumber,omitempty"`
	ResultComplete *bool          `xml:"resultComplete,omitempty"`
	States         []*StateParcel `xml:"states,omitempty"`
}

type OperationConfirm struct {
	Confirm *Confirm `xml:"confirm,omitempty"`
}

type OperationConfirmResponse struct {
	ConfirmResponse *ConfirmResponse `xml:"confirmResponse,omitempty"`
}

type OperationGetStatesByClient struct {
	GetStatesByClient *GetStatesByClient `xml:"getStatesByClient,omitempty"`
}

type OperationGetStatesByClientResponse struct {
	GetStatesByClientResponse *GetStatesByClientResponse `xml:"getStatesByClientResponse,omitempty"`
}

type OperationGetStatesByClientOrder struct {
	GetStatesByClientOrder *GetStatesByClientOrder `xml:"getStatesByClientOrder,omitempty"`
}

type OperationGetStatesByClientOrderResponse struct {
	GetStatesByClientOrderResponse *GetStatesByClientOrderResponse `xml:"getStatesByClientOrderResponse,omitempty"`
}

type OperationGetStatesByClientParcel struct {
	GetStatesByClientParcel *GetStatesByClientParcel `xml:"getStatesByClientParcel,omitempty"`
}

type OperationGetStatesByClientParcelResponse struct {
	GetStatesByClientParcelResponse *GetStatesByClientParcelResponse `xml:"getStatesByClientParcelResponse,omitempty"`
}

type OperationGetStatesByDPDOrder struct {
	GetStatesByDPDOrder *GetStatesByDPDOrder `xml:"getStatesByDPDOrder,omitempty"`
}

type OperationGetStatesByDPDOrderResponse struct {
	GetStatesByDPDOrderResponse *GetStatesByDPDOrderResponse `xml:"getStatesByDPDOrderResponse,omitempty"`
}