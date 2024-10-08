package schema

// The JoinedBy element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/joinedby
import "encoding/xml"

type JoinedBy struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
