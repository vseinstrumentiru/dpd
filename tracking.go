package dpd

const trackingNamespace = "http://dpd.ru/ws/tracing/2011-11-18"

type operationGetStatesByClient struct {
	GetStatesByClient *getStatesByClient `xml:"getStatesByClient,omitempty"`
}

type getStatesByClient struct {
	Request *getStatesByClientRequest `xml:"request,omitempty"`
}

type getStatesByClientRequest struct {
	Auth *Auth `xml:"auth,omitempty"`
}

type operationGetStatesByClientResponse struct {
	GetStatesByClientResponse *getStatesByClientResponse `xml:"getStatesByClientResponse,omitempty"`
}

type getStatesByClientResponse struct {
	Return *ParcelsStates `xml:"return,omitempty"`
}

type ParcelsStates struct {
	DocID          *int64         `xml:"docId,omitempty"`
	DocDate        *string        `xml:"docDate,omitempty"`
	ClientNumber   *int64         `xml:"clientNumber,omitempty"`
	ResultComplete *bool          `xml:"resultComplete,omitempty"`
	States         []*ParcelState `xml:"states,omitempty"`
}

type ParcelState struct {
	ClientOrderNr    *string `xml:"clientOrderNr,omitempty"`
	ClientParcelNr   *string `xml:"clientParcelNr,omitempty"`
	DpdOrderNr       *string `xml:"dpdOrderNr,omitempty"`
	DpdParcelNr      *string `xml:"dpdParcelNr,omitempty"`
	PickupDate       *string `xml:"pickupDate,omitempty"`
	DpdOrderReNr     *string `xml:"dpdOrderReNr,omitempty"`
	DpdParcelReNr    *string `xml:"dpdParcelReNr,omitempty"`
	IsReturn         *bool   `xml:"isReturn,omitempty"`
	PlanDeliveryDate *string `xml:"planDeliveryDate,omitempty"`
	NewState         *string `xml:"newState,omitempty"`
	TransitionTime   *string `xml:"transitionTime,omitempty"`
	TerminalCode     *string `xml:"terminalCode,omitempty"`
	TerminalCity     *string `xml:"terminalCity,omitempty"`
	IncidentCode     *string `xml:"incidentCode,omitempty"`
	IncidentName     *string `xml:"incidentName,omitempty"`
	Consignee        *string `xml:"consignee,omitempty"`
}

type operationGetStatesByClientOrder struct {
	GetStatesByClientOrder *getStatesByClientOrder `xml:"getStatesByClientOrder,omitempty"`
}

type getStatesByClientOrder struct {
	Request *TrackByClientOrderRequest `xml:"request,omitempty"`
}

//TrackByClientOrderRequest GetStatesByClientOrder request body
type TrackByClientOrderRequest struct {
	Auth          *Auth   `xml:"auth,omitempty"`
	ClientOrderNr *string `xml:"clientOrderNr,omitempty"`
	PickupDate    *string `xml:"pickupDate,omitempty"`
}

type operationGetStatesByClientOrderResponse struct {
	GetStatesByClientOrderResponse *getStatesByClientOrderResponse `xml:"getStatesByClientOrderResponse,omitempty"`
}

type getStatesByClientOrderResponse struct {
	Return *ParcelsStates `xml:"return,omitempty"`
}

type operationGetStatesByDPDOrder struct {
	GetStatesByDPDOrder *getStatesByDPDOrder `xml:"getStatesByDPDOrder,omitempty"`
}

type getStatesByDPDOrder struct {
	Request *TrackByDPDOrderRequest `xml:"request,omitempty"`
}

//TrackByDPDOrderRequest GetStatesByDPDOrder request body
type TrackByDPDOrderRequest struct {
	Auth       *Auth   `xml:"auth,omitempty"`
	DpdOrderNr *string `xml:"dpdOrderNr,omitempty"`
	PickupYear *int    `xml:"pickupYear,omitempty"`
}

type operationGetStatesByDPDOrderResponse struct {
	GetStatesByDPDOrderResponse *getStatesByDPDOrderResponse `xml:"getStatesByDPDOrderResponse,omitempty"`
}

type getStatesByDPDOrderResponse struct {
	Return *ParcelsStates `xml:"return,omitempty"`
}
