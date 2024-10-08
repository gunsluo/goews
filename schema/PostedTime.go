package schema

// The PostedTime element represents the time at which a PostItem was posted. This element is read-only. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/postedtime
import "time"

import "encoding/xml"

type PostedTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (P *PostedTime) SetForMarshal() {
	P.XMLName.Local = "t:PostedTime"
}

func (P *PostedTime) GetSchema() *Schema {
	return &SchemaTypes
}
