package schema

// The AddDelegateResponse element contains the status and result of an AddDelegate operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/adddelegateresponse
import "encoding/xml"

type AddDelegateResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// The ResponseMessages element contains the response messages for an Exchange Web Services delegate management request.
	ResponseMessages *ResponseMessagesArrayOfDelegateUserResponseMessageType `xml:"ResponseMessages"`
}

func (A *AddDelegateResponse) SetForMarshal() {
	A.XMLName.Local = "m:AddDelegateResponse"
}

func (A *AddDelegateResponse) GetSchema() *Schema {
	return &SchemaMessages
}
