package dpdlib

import "github.com/fiorix/wsdl2go/soap"

type GeographyAddress struct {
	CityId      *int64  `xml:"cityId,omitempty"`
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
	Descript    *string `xml:"descript,omitempty"`
}

type City struct {
	CityId       *int64  `xml:"cityId,omitempty"`
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

type DpdCitiesCashPayRequest struct {
	Auth        *Auth   `xml:"auth,omitempty"`
	CountryCode *string `xml:"countryCode,omitempty"`
}

type DpdParcelShopRequest struct {
	Ns			string	`xml:"xmlns,attr"`
	Auth        *Auth   `xml:"auth,omitempty"`
	CountryCode *string `xml:"countryCode,omitempty"`
	RegionCode  *string `xml:"regionCode,omitempty"`
	CityCode    *string `xml:"cityCode,omitempty"`
	CityName    *string `xml:"cityName,omitempty"`
}

type DpdParcelShops struct {
	ParcelShop []*ParcelShop `xml:"parcelShop,omitempty"`
}

type DpdPossibleESOption struct {
	Option []*string `xml:"option,omitempty"`
}

type DpdPossibleESPickupDelivery struct {
	CityId       *float64 `xml:"cityId,omitempty"`
	TerminalCode *string  `xml:"terminalCode,omitempty"`
	Index        *float64 `xml:"index,omitempty"`
	CityName     *string  `xml:"cityName,omitempty"`
	RegionCode   *string  `xml:"regionCode,omitempty"`
	CountryCode  *string  `xml:"countryCode,omitempty"`
}

type DpdPossibleESRequest struct {
	Auth         *Auth                        `xml:"auth,omitempty"`
	Pickup       *DpdPossibleESPickupDelivery `xml:"pickup,omitempty"`
	Delivery     *DpdPossibleESPickupDelivery `xml:"delivery,omitempty"`
	SelfPickup   *bool                        `xml:"selfPickup,omitempty"`
	SelfDelivery *bool                        `xml:"selfDelivery,omitempty"`
	ServiceCode  *string                      `xml:"serviceCode,omitempty"`
	PickupDate   *Date                        `xml:"pickupDate,omitempty"`
	Options      *DpdPossibleESOption         `xml:"options,omitempty"`
}

type DpdPossibleESResult struct {
	ResultCode    *string                 `xml:"resultCode,omitempty"`
	ResultMessage *string                 `xml:"resultMessage,omitempty"`
	ExtraService  []*PossibleExtraService `xml:"extraService,omitempty"`
}

type DpdTerminals struct {
	Terminal []*TerminalStoragePeriods `xml:"terminal,omitempty"`
}

type ExtraService struct {
	EsCode *string              `xml:"esCode,omitempty"`
	Params []*ExtraServiceParam `xml:"params,omitempty"`
}

type ExtraServiceParam struct {
	Name  *string `xml:"name,omitempty"`
	Value *string `xml:"value,omitempty"`
}

type GeoCoordinates struct {
	Latitude  *float64 `xml:"latitude,omitempty"`
	Longitude *float64 `xml:"longitude,omitempty"`
}

type GetCitiesCashPay struct {
	Request *DpdCitiesCashPayRequest `xml:"request,omitempty"`
}

type GetCitiesCashPayResponse struct {
	Return []*City `xml:"return,omitempty"`
}

type GetParcelShops struct {
	Request *DpdParcelShopRequest `xml:"request,omitempty"`
	Ns 			string  `xml:"xmlns,attr"`
}

type GetParcelShopsResponse struct {
	Return *DpdParcelShops `xml:"return,omitempty"`
}

type GetPossibleExtraService struct {
	Request *DpdPossibleESRequest `xml:"request,omitempty"`
}

type GetPossibleExtraServiceResponse struct {
	Return *DpdPossibleESResult `xml:"return,omitempty"`
}

type GetStoragePeriod struct {
	Request *StoragePeriodRequest `xml:"request,omitempty"`
}

type GetStoragePeriodResponse struct {
	Return *DpdTerminals `xml:"return,omitempty"`
}

type GetTerminalsSelfDelivery2 struct {
	Ns			string	`xml:"xmlns,attr"`
	Auth *Auth `xml:"auth,omitempty"`
}

type GetTerminalsSelfDelivery2Response struct {
	Return *SelfTerminals `xml:"return,omitempty"`
}

type Limits struct {
	MaxShipmentWeight *float64 `xml:"maxShipmentWeight,omitempty"`
	MaxWeight         *float64 `xml:"maxWeight,omitempty"`
	MaxLength         *float64 `xml:"maxLength,omitempty"`
	MaxWidth          *float64 `xml:"maxWidth,omitempty"`
	MaxHeight         *float64 `xml:"maxHeight,omitempty"`
	DimensionSum      *float64 `xml:"dimensionSum,omitempty"`
}

type ParcelShop struct {
	Code                *string           `xml:"code,omitempty"`
	ParcelShopType      *string           `xml:"parcelShopType,omitempty"`
	State               *string           `xml:"state,omitempty"`
	Address             *GeographyAddress `xml:"address,omitempty"`
	Brand               *string           `xml:"brand,omitempty"`
	ClientDepartmentNum *string           `xml:"clientDepartmentNum,omitempty"`
	GeoCoordinates      *GeoCoordinates   `xml:"geoCoordinates,omitempty"`
	Limits              *Limits           `xml:"limits,omitempty"`
	Schedule            []*Schedule     `xml:"schedule,omitempty"`
	ExtraService        []*ExtraService `xml:"extraService,omitempty"`
	Services            []*string       `xml:"services>serviceCode,omitempty"`
}

type PossibleExtraService struct {
	Code   *string `xml:"code,omitempty"`
	Name   *string `xml:"name,omitempty"`
	IsPaid *bool   `xml:"isPaid,omitempty"`
}

type Schedule struct {
	Operation *string      `xml:"operation,omitempty"`
	Timetable []*Timetable `xml:"timetable,omitempty"`
}

type SelfTerminals struct {
	Terminal []*TerminalSelf `xml:"terminal,omitempty"`
}

type StoragePeriod struct {
	ServiceCode *string `xml:"serviceCode,omitempty"`
	Days        *int    `xml:"days,omitempty"`
}

type StoragePeriodRequest struct {
	Auth         *Auth   `xml:"auth,omitempty"`
	TerminalCode *string `xml:"terminalCode,omitempty"`
	ServiceCode  *string `xml:"serviceCode,omitempty"`
}

type TerminalSelf struct {
	TerminalCode   *string           `xml:"terminalCode,omitempty"`
	TerminalName   *string           `xml:"terminalName,omitempty"`
	Address        *GeographyAddress `xml:"address,omitempty"`
	GeoCoordinates *GeoCoordinates   `xml:"geoCoordinates,omitempty"`
	Schedule       []*Schedule       `xml:"schedule,omitempty"`
	ExtraService   []*ExtraService   `xml:"extraService,omitempty"`
	Services       []*string         `xml:"services>serviceCode,omitempty"`
}

type TerminalStoragePeriods struct {
	TerminalCode *string          `xml:"terminalCode,omitempty"`
	Services     []*StoragePeriod `xml:"services,omitempty"`
}

type Timetable struct {
	WeekDays *string `xml:"weekDays,omitempty"`
	WorkTime *string `xml:"workTime,omitempty"`
}

type OperationGetCitiesCashPay struct {
	GetCitiesCashPay *GetCitiesCashPay `xml:"getCitiesCashPay,omitempty"`
}

type OperationGetCitiesCashPayResponse struct {
	GetCitiesCashPayResponse *GetCitiesCashPayResponse `xml:"getCitiesCashPayResponse,omitempty"`
}

type OperationGetParcelShops struct {
	GetParcelShops *GetParcelShops `xml:"getParcelShops,omitempty"`
}

type OperationGetParcelShopsResponse struct {
	GetParcelShopsResponse *GetParcelShopsResponse `xml:"getParcelShopsResponse,omitempty"`
}

type OperationGetPossibleExtraService struct {
	GetPossibleExtraService *GetPossibleExtraService `xml:"getPossibleExtraService,omitempty"`
}

type OperationGetPossibleExtraServiceResponse struct {
	GetPossibleExtraServiceResponse *GetPossibleExtraServiceResponse `xml:"getPossibleExtraServiceResponse,omitempty"`
}

type OperationGetStoragePeriod struct {
	GetStoragePeriod *GetStoragePeriod `xml:"getStoragePeriod,omitempty"`
}

type OperationGetStoragePeriodResponse struct {
	GetStoragePeriodResponse *GetStoragePeriodResponse `xml:"getStoragePeriodResponse,omitempty"`
}

type OperationGetTerminalsSelfDelivery2 struct {
	GetTerminalsSelfDelivery2 *GetTerminalsSelfDelivery2 `xml:"getTerminalsSelfDelivery2,omitempty"`
}

type OperationGetTerminalsSelfDelivery2Response struct {
	GetTerminalsSelfDelivery2Response *GetTerminalsSelfDelivery2Response `xml:"getTerminalsSelfDelivery2Response,omitempty"`
}

type dPDGeography2 struct {
	cli *soap.Client
}