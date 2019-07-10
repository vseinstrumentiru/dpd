 # SDK для интеграции с SOAP сервисами транспортной компании DPD
[![GoDoc reference](https://godoc.org/github.com/vseinstrumentiru/dpd?status.svg)](https://godoc.org/github.com/vseinstrumentiru/dpd) 

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
  - [x] getOrderStatus
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
dpdClient := dpdSdk.NewDPDClient(clinetNumber, clientKey
    ServiceUrls{
        GeographyUrl:  "http://wstest.dpd.ru/services/geography2",
        OrderUrl:      "http://wstest.dpd.ru/services/order2",
        CalculatorUrl: "http://wstest.dpd.ru/services/calculator2",
        TrackingUrl:   "http://wstest.dpd.ru/services/tracing",
    },
)


calcRequest := dpdSdk.NewCalculateRequest().
    SetPickup(dpdSdk.NewCity(48951627)).
    SetDelivery(dpdSdk.NewCity(195595210)).
    SetWeight(2.34).
    SetSelfPickup(false).
    SetSelfDelivery(false)

res, err := dpdClient.GetServiceCost2(calcRequest)

```