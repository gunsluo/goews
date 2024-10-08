package schema

// The ConvertIdResponse element contains a response to a ConvertId request. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/convertidresponse
import "encoding/xml"

type ConvertIdResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (C *ConvertIdResponse) SetForMarshal() {
	C.XMLName.Local = "m:ConvertIdResponse"
}

func (C *ConvertIdResponse) GetSchema() *Schema {
	return &SchemaMessages
}
