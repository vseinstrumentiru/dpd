package dpdlib

import (
	"fmt"
	"github.com/fiorix/wsdl2go/soap"
	"log"
	"strings"
)

type commonPoint struct {
	code         string
	latitude     string
	longitude    string
	maxWeight    float32
	maxWidth     float32
	maxHeight    float32
	maxLength    float32
	dimensionSum float32
	schedule     string
	cityCode     uint32
	state        string
}

type normalizedLimits struct {
	maxWeight float32
	maxWidth float32
	maxHeight float32
	maxLength float32
	dimensionSum float32
}

func NewDpdSdk(clientNumber int64, clientKey, countryCode string) dpdSdk {
	return dpdSdk{
		clientNumber: clientNumber,
		clientKey:    clientKey,
		countryCode:  countryCode,
	}
}

type dpdSdk struct {
	clientNumber int64
	clientKey    string
	countryCode  string
}

type DpdSdk interface {
	GetOffers() []*ServiceCost
	GetPoints()
	getLimitedPoints(geography2 DPDGeography2)
	//getUnlimitedPoints(geography2 DPDGeography2)

}

func (dpdSdk) GetOffers() []*ServiceCost {
	client := soap.Client{
		URL:       "http://wstest.dpd.ru:80/services/calculator",
		Namespace: CalculatorNamespace,
	}

	service := NewDPDCalculator(&client)

	getServiceCost2Req := &GetServiceCost2{
		//TODO: implement req body
	}

	result, err := service.GetServiceCost2(getServiceCost2Req)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	return result.Return
}

func (sdk dpdSdk) GetPoints() []*commonPoint {
	client := soap.Client{
		URL:       "http://wstest.dpd.ru/services/geography2",
		Namespace: GeographyNamespace,
	}

	service := NewDPDGeography2(&client)

	limitedPoints := sdk.getLimitedPoints(service)
	unlimitedPoints := sdk.getUnlimitedPoints(service)
	result := append(limitedPoints, unlimitedPoints...)

	return result
}

func (sdk dpdSdk) getLimitedPoints(service DPDGeography2) []*commonPoint {

	limitedPointsReq := &GetParcelShops{
		Ns: GeographyNamespace,
		Request: &DpdParcelShopRequest{
			Auth: &Auth{
				ClientNumber: &sdk.clientNumber,
				ClientKey:    &sdk.clientKey,
			},
			CountryCode: &sdk.countryCode,
		},
	}

	result, err := service.GetParcelShops(limitedPointsReq)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	return convertDpdLimitedPoints2Common(result.Return.ParcelShop)
}

func (sdk dpdSdk) getUnlimitedPoints(service DPDGeography2) []*commonPoint {
	unlimitedPointsReq := &GetTerminalsSelfDelivery2{
		Ns: GeographyNamespace,
		Auth: &Auth{
			ClientNumber: &sdk.clientNumber,
			ClientKey:    &sdk.clientKey,
		},
	}

	result, err := service.GetTerminalsSelfDelivery2(unlimitedPointsReq)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	commonPoints := convertDpdUnlimitedPoints2Common(result.Return.Terminal)

	return commonPoints
}

func convertDpdUnlimitedPoints2Common(terminals []*TerminalSelf) []*commonPoint {
	var commonPoints []*commonPoint

	for _, terminal := range terminals {
		commonPoints = append(commonPoints, &commonPoint{
			code:      *terminal.TerminalCode,
			latitude:  fmt.Sprintf("%g", *terminal.GeoCoordinates.Latitude),
			longitude: fmt.Sprintf("%g", *terminal.GeoCoordinates.Longitude),
			schedule:  getStrSchedule(terminal.Schedule),
			cityCode:  uint32(*terminal.Address.CityId),
			state:     "Open",
		})
	}

	return commonPoints
}

func convertDpdLimitedPoints2Common(shops []*ParcelShop) []*commonPoint {
	var commonPoints []*commonPoint

	for _, shop := range shops {

		weight, width, length, height, dimensionSum := getValidLimits(shop.Limits)

		commonPoints = append(commonPoints, &commonPoint{
			code:         *shop.Code,
			latitude:     fmt.Sprintf("%g", *shop.GeoCoordinates.Latitude),
			longitude:    fmt.Sprintf("%g", *shop.GeoCoordinates.Longitude),
			maxWeight:    weight,
			maxHeight:    height,
			maxLength:    length,
			maxWidth:     width,
			dimensionSum: dimensionSum,
			schedule:     getStrSchedule(shop.Schedule),
			cityCode:     uint32(*shop.Address.CityId),
			state:        *shop.State,
		})
	}

	return commonPoints
}

func getValidLimits(limits *Limits) (float32, float32, float32, float32, float32) {

	var weight, width, length, height, dimensionSum float32

	if limits == nil {
		return weight, width, length, height, dimensionSum
	}

	if limits.MaxWeight != nil {
		weight = float32(*limits.MaxWeight)
	}

	if limits.MaxWidth != nil {
		width = float32(*limits.MaxWidth)
	}

	if limits.MaxLength != nil {
		length = float32(*limits.MaxLength)
	}

	if limits.MaxHeight != nil {
		height = float32(*limits.MaxHeight)
	}

	if limits.DimensionSum != nil {
		dimensionSum = float32(*limits.DimensionSum)
	}

	return weight, width, length, height, dimensionSum
}

func getStrSchedule(schedule []*Schedule) string {

	var timetable []*Timetable
	var normalizedSchedule string

	for _, item := range schedule {
		if strings.Compare(*item.Operation, "SelfPickup") == 0 {
			timetable = item.Timetable
			break
		}
	}

	for _, schedule := range timetable {
		normalizedSchedule += *schedule.WeekDays + " " + *schedule.WorkTime
	}

	return normalizedSchedule
}
