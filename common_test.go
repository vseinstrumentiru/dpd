package dpd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	anyString           = "any string"
	anyKey              = "any key"
	anyNumber           = 1
	internalOrderNumber = "1901-200100-42389"
	dpdOrderNumber      = "RU019202656"
	thisCountry         = "Россия"
)

func makeTestHTTPServer(t *testing.T, status int, body []byte) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/xml")

		if _, err := w.Write(body); err != nil {
			t.Errorf("Error writing response %s", err.Error())
			t.FailNow()
		}

		//if !bytes.Contains(body, []byte("clientNumber")) || !bytes.Contains(body, []byte("clientKey")) {
		//	t.Errorf("Auth tag is necessary")
		//	t.FailNow()
		//}
	}))

	return ts
}
