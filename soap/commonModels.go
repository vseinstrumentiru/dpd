package soap

type Date string

type DateTime string

type Auth struct {
	Ns           *string `xml:"xmlns,attr"`
	ClientNumber *int64  `xml:"clientNumber,omitempty"`
	ClientKey    *string `xml:"clientKey,omitempty"`
}

type WSFault struct {
	Code    *string `xml:"code,omitempty"`
	Message *string `xml:"message,omitempty"`
}
