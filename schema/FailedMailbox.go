package schema

// The FailedMailbox element specifies the error message for a mailbox that failed on search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/failedmailbox
import "encoding/xml"

type FailedMailbox struct {
	XMLName xml.Name

	// The ErrorCode element specifies the error code of a failed search performed against a mailbox.
	ErrorCode *ErrorCodeint `xml:"ErrorCode"`
	// The ErrorMessage element represents the reason for the validation error.
	ErrorMessage *ErrorMessage `xml:"ErrorMessage"`
	// The IsArchive element specifies a Boolean value that indicates whether the mailbox is an archive mailbox.
	IsArchive *IsArchive `xml:"IsArchive"`
	// The Mailbox element contains an identifier for a mailbox.
	Mailbox *Mailboxstring `xml:"Mailbox"`
}

func (F *FailedMailbox) SetForMarshal() {
	F.XMLName.Local = "t:FailedMailbox"
}

func (F *FailedMailbox) GetSchema() *Schema {
	return &SchemaTypes
}
