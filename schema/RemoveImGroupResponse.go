package schema

// The RemoveImGroupResponse element represents a response to a RemoveImGroup request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/removeimgroupresponse
import "encoding/xml"

type RemoveImGroupResponse struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
}

func (R *RemoveImGroupResponse) SetForMarshal() {
	R.XMLName.Local = "m:RemoveImGroupResponse"
}

func (R *RemoveImGroupResponse) GetSchema() *Schema {
	return &SchemaMessages
}
