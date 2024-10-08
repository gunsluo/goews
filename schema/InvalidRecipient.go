package schema

// The InvalidRecipient element contains the SMTP address of the invalid recipient and information about why the recipient is invalid.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/invalidrecipient
import "encoding/xml"

type InvalidRecipient struct {
	XMLName xml.Name

	// The MessageText element provides a text description of the status of the response.
	MessageText *MessageText `xml:"MessageText"`
	// The ResponseCode element provides information about why the recipient is invalid.
	ResponseCode *ResponseCodeInvalidRecipientResponseCodeType `xml:"ResponseCode"`
	// The SmtpAddress element represents the Simple Mail Transfer Protocol (SMTP) address of an account to be used for impersonation or a Simple Mail Transfer Protocol (SMTP) recipient address of a calendar or contact sharing request.
	SmtpAddress *SmtpAddress `xml:"SmtpAddress"`
}

func (I *InvalidRecipient) SetForMarshal() {
	I.XMLName.Local = "t:InvalidRecipient"
}

func (I *InvalidRecipient) GetSchema() *Schema {
	return &SchemaTypes
}
