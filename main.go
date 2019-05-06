package main

import "dpd-sandbox/dpdlib"

func main()  {

	dpdSdk := dpdlib.NewDpdSdk(1001027554, "01DA4E5140C3FBBAB52A55262C2FB231E5EE499E", "RU")

	dpdSdk.GetPoints()
}