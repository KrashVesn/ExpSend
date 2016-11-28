package get_str

import "encoding/xml"

//import "text/template"

type ApiResponse struct {
	XMLName xml.Name `xml:"ApiResponse"`
	Data    []*List  `xml:"Data>Lists>List"`
}

type List struct {
	Id           string
	Name         string
	FriendlyName string
	Language     string
	OptInMode    string
}

type ApiResponse2 struct {
	XMLName xml.Name `xml:"ApiResponse"`
	Data    string   `xml:"Data"`
}
