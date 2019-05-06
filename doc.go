//SDK для интеграции с транспортной компанией DPD
//
//Клиент и построители запросов
//
//Пример использования
// 	dpdClient := dpd_sandbox.NewDPDClient(
//		dpd_sandbox.DPDAuth{
//			ClientNumber: clientNumber(int64),
//			ClientKey:    clientKey(string),
//		},
//		dpd_sandbox.DPDUrls{
//			GeographyUrl:  "http://wstest.dpd.ru/services/geography2",
//			OrderUrl:      "http://wstest.dpd.ru/services/order",
//			CalculatorUrl: "http://wstest.dpd.ru/services/calculator2",
//			TrackingUrl:   "http://wstest.dpd.ru:80/services/tracing",
//		},
//		"RU",
//	)
//
//	calcRequest := dpd_sandbox.NewCalculateRequest().
//	SetPickup(dpd_sandbox.NewCity(48951627)).
//	SetDelivery(dpd_sandbox.NewCity(195595210)).
//	SetWeight(2.34).
//	SetSelfPickup(false).
//	SetSelfDelivery(false)
//
//	res, err := dpdClient.GetServiceCost2(calcRequest)
//
//	if err != nil {
//		log.Fatalf(err.Error())
//	}
//
package dpd_sdk
