 # SDK для интеграции с SOAP сервисами транспортной компании DPD
[![Build Status](https://travis-ci.com/seacomandor/dpd.svg?branch=master)](https://travis-ci.com/seacomandor/dpd) 
[![Coverage Status](https://coveralls.io/repos/github/seacomandor/dpd/badge.svg?branch=master)](https://coveralls.io/github/seacomandor/dpd?branch=master) 
[![GoDoc reference](https://godoc.org/github.com/seacomandor/dpd?status.svg)](https://godoc.org/github.com/seacomandor/dpd) 

Roadmap
- [ ] Web служба "Георгафия DPD"
  - [X] getCitiesCashPay
  - [X] getParcelShops
  - [X] getTerminalsSelfDelivery2
  - [ ] getStoragePeriod
  - [ ] getPossibleExtraService
- [ ] Web служба "Расчет стоимости"
  - [X] getServiceCost2
  - [ ] getServiceCostByParcels2
  - [ ] getServiceCostInternational  
- [ ] Web служба "Создание заказа"
  - [X] createOrder
  - [ ] getOrderStatus
  - [ ] createAddress
  - [ ] updateAddress
  - [ ] getInvoiceFile
  - [ ] getRegisterFile
  - [ ] addAirwayBill
  - [ ] changeUnitLoad
- [ ] Web служба "Изменение заказа"
  - [X] addParcels
  - [X] removeParcel
- [ ] Web служба "Отмена заказа"
  - [X] cancelOrder
- [ ] Web служба "Отслеживание статуса"
  - [X] getStatesByClient
  - [X] getStatesByClientOrder§
  - [ ] confirm
  - [ ] getStatesByClientParcel
  - [X] getStatesByDPDOrders
  - [ ] getEvents
  - [ ] getTrackingOrderLink
- [ ] Web служба "Отчеты" 
  - [ ] getNLAmount
  - [ ] getNLInvoice
  - [ ] getWaybill
- [ ] Web служба "Печать наклейки"
  - [ ] createLabelFile
  - [ ] createParcelLabel


## Пример использования

```go
dpdClient := dpdSdk.NewDPDClient(
    dpdSdk.DPDAuth{
        ClientNumber: client number,
        ClientKey: client key ,
    },
    dpdSdk.DPDUrls{
        GeographyUrl:  "http://wstest.dpd.ru/services/geography2",
        OrderUrl:      "http://wstest.dpd.ru/services/order",
        CalculatorUrl: "http://wstest.dpd.ru/services/calculator2",
        TrackingUrl:   "http://wstest.dpd.ru:80/services/tracing",
    },
    "RU",
)


calcRequest := dpdSdk.NewCalculateRequest().
    SetPickup(dpdSdk.NewCity(48951627)).
    SetDelivery(dpdSdk.NewCity(195595210)).
    SetWeight(2.34).
    SetSelfPickup(false).
    SetSelfDelivery(false)

res, err := dpdClient.GetServiceCost2(calcRequest)

```