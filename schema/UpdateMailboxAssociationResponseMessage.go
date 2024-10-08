package schema

// The UpdateMailboxAssociationResponseMessage element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updatemailboxassociationresponsemessage
import "encoding/xml"

type UpdateMailboxAssociationResponseMessage struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
