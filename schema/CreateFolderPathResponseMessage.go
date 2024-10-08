package schema

// The CreateFolderPathResponseMessage element specifies the response message for a CreateFolderPath request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createfolderpathresponsemessage
import "encoding/xml"

type CreateFolderPathResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The Folders element contains an array of folders that are used in folder operations.
	Folders *Folders `xml:"Folders"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Indicates the class of the response.
	ResponseClass *string `xml:"ResponseClass,attr"`
}

func (C *CreateFolderPathResponseMessage) SetForMarshal() {
	C.XMLName.Local = "m:CreateFolderPathResponseMessage"
}

func (C *CreateFolderPathResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
