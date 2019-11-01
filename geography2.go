package dpd

const (
	geographyNamespace = "http://dpd.ru/ws/geography/2015-05-20"
	cityCODPaymentLimit = 250000
)

type operationGetCitiesCashPay struct {
	GetCitiesCashPay *getCitiesCashPay `xml:"getCitiesCashPay,omitempty"`
}

type getCitiesCashPay struct {
	Namespace string                `xml:"xmlns,attr"`
	Request   *citiesCashPayRequest `xml:"request,omitempty"`
}

type citiesCashPayRequest struct {
	Namespace   string  `xml:"xmlns,attr"`
	Auth        *Auth   `xml:"auth,omitempty"`
	CountryCode *string `xml:"countryCode,omitempty"`
}

type operationGetCitiesCashPayResponse struct {
	GetCitiesCashPayResponse *GetCitiesCashPayResponse `xml:"getCitiesCashPayResponse,omitempty"`
}

//GetCitiesCashPayResponse list of DPD cities with C.O.D available
type GetCitiesCashPayResponse struct {
	Return []*City `xml:"return,omitempty"`
}

//City with C.O.D available
type City struct {
	CityID       *int64  `xml:"cityId,omitempty"`
	CountryCode  *string `xml:"countryCode,omitempty"`
	CountryName  *string `xml:"countryName,omitempty"`
	RegionCode   *int    `xml:"regionCode,omitempty"`
	RegionName   *string `xml:"regionName,omitempty"`
	CityCode     *string `xml:"cityCode,omitempty"`
	CityName     *string `xml:"cityName,omitempty"`
	Abbreviation *string `xml:"abbreviation,omitempty"`
	IndexMin     *string `xml:"indexMin,omitempty"`
	IndexMax     *string `xml:"indexMax,omitempty"`
}

type operationGetParcelShops struct {
	GetParcelShops *getParcelShops `xml:"getParcelShops,omitempty"`
}

type getParcelShops struct {
	Request *ParcelShopRequest `xml:"request,omitempty"`
	Ns      string             `xml:"xmlns,attr"`
}

//ParcelShopRequest GetParcelShops request body
type ParcelShopRequest struct {
	Namespace   string  `xml:"xmlns,attr"`
	Auth        *Auth   `xml:"auth,omitempty"`
	CountryCode *string `xml:"countryCode,omitempty"`
	RegionCode  *string `xml:"regionCode,omitempty"`
	CityCode    *string `xml:"cityCode,omitempty"`
	CityName    *string `xml:"cityName,omitempty"`
}

type operationGetParcelShopsResponse struct {
	GetParcelShopsResponse *getParcelShopsResponse `xml:"getParcelShopsResponse,omitempty"`
}

type getParcelShopsResponse struct {
	Return *parcelShops `xml:"return,omitempty"`
}

type parcelShops struct {
	ParcelShop []*ParcelShop `xml:"parcelShop,omitempty"`
}

//ParcelShop self delivery point with limits on dimensions and weight
type ParcelShop struct {
	Code                *string                   `xml:"code,omitempty"`
	ParcelShopType      *string                   `xml:"parcelShopType,omitempty"`
	State               *string                   `xml:"state,omitempty"`
	Address             *GeographyAddress         `xml:"address,omitempty"`
	Brand               *string                   `xml:"brand,omitempty"`
	ClientDepartmentNum *string                   `xml:"clientDepartmentNum,omitempty"`
	GeoCoordinates      *GeoCoordinates           `xml:"geoCoordinates,omitempty"`
	Limits              *Limits                   `xml:"limits,omitempty"`
	Schedule            []*Schedule               `xml:"schedule,omitempty"`
	ExtraService        []*ExtraServiceParameters `xml:"extraService,omitempty"`
	Services            []*string                 `xml:"services>serviceCode,omitempty"`
}

//GeographyAddress of ParcelShop or Terminal
type GeographyAddress struct {
	CityID      *int64  `xml:"cityId,omitempty"`
	CountryCode *string `xml:"countryCode,omitempty"`
	RegionCode  *string `xml:"regionCode,omitempty"`
	RegionName  *string `xml:"regionName,omitempty"`
	CityCode    *string `xml:"cityCode,omitempty"`
	CityName    *string `xml:"cityName,omitempty"`
	Index       *string `xml:"index,omitempty"`
	Street      *string `xml:"street,omitempty"`
	StreetAbbr  *string `xml:"streetAbbr,omitempty"`
	HouseNo     *string `xml:"houseNo,omitempty"`
	Building    *string `xml:"building,omitempty"`
	Structure   *string `xml:"structure,omitempty"`
	Ownership   *string `xml:"ownership,omitempty"`
	Description *string `xml:"descript,omitempty"`
}

//GeoCoordinates of ParcelShop or Terminal
type GeoCoordinates struct {
	Latitude  *float64 `xml:"latitude,omitempty"`
	Longitude *float64 `xml:"longitude,omitempty"`
}

//Limits of ParcelShop
type Limits struct {
	MaxShipmentWeight *float64 `xml:"maxShipmentWeight,omitempty"`
	MaxWeight         *float64 `xml:"maxWeight,omitempty"`
	MaxLength         *float64 `xml:"maxLength,omitempty"`
	MaxWidth          *float64 `xml:"maxWidth,omitempty"`
	MaxHeight         *float64 `xml:"maxHeight,omitempty"`
	DimensionSum      *float64 `xml:"dimensionSum,omitempty"`
}

//Schedule of ParcelShop or Terminal
type Schedule struct {
	Operation *string      `xml:"operation,omitempty"`
	Timetable []*Timetable `xml:"timetable,omitempty"`
}

//Timetable for operations
type Timetable struct {
	WeekDays *string `xml:"weekDays,omitempty"`
	WorkTime *string `xml:"workTime,omitempty"`
}

//ExtraServiceParameters for ParcelShop or Terminal
type ExtraServiceParameters struct {
	EsCode *string                  `xml:"esCode,omitempty"`
	Params []*ExtraServiceParameter `xml:"params,omitempty"`
}

//ExtraServiceParameter ...
type ExtraServiceParameter struct {
	Name  *string `xml:"name,omitempty"`
	Value *string `xml:"value,omitempty"`
}

type operationGetTerminalsSelfDelivery2 struct {
	GetTerminalsSelfDelivery2 *getTerminalsSelfDelivery2Request `xml:"getTerminalsSelfDelivery2,omitempty"`
}

type getTerminalsSelfDelivery2Request struct {
	Namespace string `xml:"xmlns,attr"`
	Auth      *Auth  `xml:"auth,omitempty"`
}

type operationGetTerminalsSelfDelivery2Response struct {
	Response *getTerminalsSelfDelivery2Response `xml:"getTerminalsSelfDelivery2Response,omitempty"`
}

type getTerminalsSelfDelivery2Response struct {
	Return *Terminals `xml:"return,omitempty"`
}

//Terminals GetTerminalsSelfDelivery2 response body
type Terminals struct {
	Terminal []*Terminal `xml:"terminal,omitempty"`
}

//Terminal ...
type Terminal struct {
	TerminalCode   *string                   `xml:"terminalCode,omitempty"`
	TerminalName   *string                   `xml:"terminalName,omitempty"`
	Address        *GeographyAddress         `xml:"address,omitempty"`
	GeoCoordinates *GeoCoordinates           `xml:"geoCoordinates,omitempty"`
	Schedule       []*Schedule               `xml:"schedule,omitempty"`
	ExtraService   []*ExtraServiceParameters `xml:"extraService,omitempty"`
	Services       []*string                 `xml:"services>serviceCode,omitempty"`
}
