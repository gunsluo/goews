package schema

// The FirstMatchingRowIndex element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/firstmatchingrowindex
import "encoding/xml"

type FirstMatchingRowIndex struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
