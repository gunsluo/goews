package schema

// The DelegateUserResponseMessageType element contains the response message for a single delegate user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/delegateuserresponsemessagetype
import "encoding/xml"

type DelegateUserResponseMessageType struct {
	XMLName xml.Name

	// The DelegateUser element identifies a single delegate to add or update in a mailbox or a delegate returned in a delegate management response. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	DelegateUser *DelegateUser `xml:"DelegateUser"`
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (D *DelegateUserResponseMessageType) SetForMarshal() {
	D.XMLName.Local = "m:DelegateUserResponseMessageType"
}

func (D *DelegateUserResponseMessageType) GetSchema() *Schema {
	return &SchemaMessages
}
