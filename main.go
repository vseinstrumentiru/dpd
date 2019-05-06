package main

import (
	"dpd-sandbox/dpdlib"
	"time"
)

func main() {
	defer dpdlib.PrintTime("main", time.Now())

	dpdSdk := dpdlib.NewDpdSdk(1001027554, "01DA4E5140C3FBBAB52A55262C2FB231E5EE499E", "RU")
	dpdSdk.GetPoints()
}
