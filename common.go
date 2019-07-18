package dpd

//Auth part of all requests. Credentials to DPD SOAP API
type Auth struct {
	Namespace    *string `xml:"xmlns,attr"`
	ClientNumber *int64  `xml:"clientNumber,omitempty"`
	ClientKey    *string `xml:"clientKey,omitempty"`
}

type commonUnitLoad struct {
	Article       *string `xml:"article,omitempty"`
	Description   *string `xml:"descript,omitempty"`
	DeclaredValue *string `xml:"declared_value,omitempty"`
	ParcelNum     *string `xml:"parcel_num,omitempty"`
	NppAmount     *string `xml:"npp_amount,omitempty"`
	VatPercent    *int    `xml:"vat_percent,omitempty"`
	WithoutVat    *bool   `xml:"without_vat,omitempty"`
	Count         *int    `xml:"count,omitempty"`
}
