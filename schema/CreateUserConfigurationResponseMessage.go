package schema

// The CreateUserConfigurationResponseMessage element contains the status and result of a single CreateUserConfiguration request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createuserconfigurationresponsemessage
import "encoding/xml"

type CreateUserConfigurationResponseMessage struct {
	XMLName xml.Name

	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey"`
	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml"`
	// The ResponseCode element provides status information about the request.
	ResponseCode *ResponseCode `xml:"ResponseCode"`
	// Describes the status of the response.The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	CreateUserConfigurationResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline during the batch.  - Active Directory Domain Services (AD DS) is offline.  - Mailboxes were moved.  - The message database (MDB) is offline.  - A password is expired.  - A quota has been exceeded.
	CreateUserConfigurationResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources of errors:  - Invalid attributes or elements.  - Attributes or elements that are out of range.  - An unknown tag.  - The attribute or element is not valid in the context.  - An unauthorized access attempt by any client.  - A server-side failure in response to a valid client-side call.  Information about the error can be found in the ResponseCode and MessageText elements.
	CreateUserConfigurationResponseMessageError = `Error`
)

func (C *CreateUserConfigurationResponseMessage) SetForMarshal() {
	C.XMLName.Local = "m:CreateUserConfigurationResponseMessage"
}

func (C *CreateUserConfigurationResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
