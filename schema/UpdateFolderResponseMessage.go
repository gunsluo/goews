package schema

// The UpdateFolderResponseMessage element contains the status and result of updates defined by the FolderChange element of an UpdateFolder operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatefolderresponsemessage
import "encoding/xml"

type UpdateFolderResponseMessage struct {
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
	// Describes the status of an UpdateFolder operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass *string `xml:"ResponseClass,attr"`
}

const (
	// Describes a request that is fulfilled.
	UpdateFolderResponseMessageSuccess = `Success`
	// Describes a request that was not processed. A warning may be returned if an error occurred while an item in the request was processing and subsequent items could not be processed. The following are examples of sources of warnings:  - The Exchange store is offline.  - Active Directory Domain Services (AD DS) is offline.  - A mailbox is moved.  - A password is expired.  - A quota is exceeded.
	UpdateFolderResponseMessageWarning = `Warning`
	// Describes a request that cannot be fulfilled. The following are examples of sources for errors:  - Invalid attributes or elements  - Attributes or elements out of range  - Unknown tag  - Attribute or element not valid in the context  - Unauthorized access attempt by any client  - Server-side failure in response to a valid client-side call    Information about the error can be found in the ResponseCode and MessageText elements.
	UpdateFolderResponseMessageError = `Error`
)

func (U *UpdateFolderResponseMessage) SetForMarshal() {
	U.XMLName.Local = "m:UpdateFolderResponseMessage"
}

func (U *UpdateFolderResponseMessage) GetSchema() *Schema {
	return &SchemaMessages
}
