package schema

// The GetNonIndexableItemDetails element specifies a request to retrieve nonindexable item details.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getnonindexableitemdetails
import "encoding/xml"

type GetNonIndexableItemDetails struct {
	XMLName xml.Name

	// The Mailboxes element specifies an array of mailboxes identified by legacy distinguished name.
	Mailboxes *MailboxesNonEmptyArrayOfLegacyDNsType `xml:"Mailboxes"`
	// The PageDirection element contains the direction for pagination in the search result. The value is Previous or Next.
	PageDirection *PageDirection `xml:"PageDirection"`
	// The PageItemReference element specifies the reference for a page item.
	PageItemReference *PageItemReference `xml:"PageItemReference"`
	// The PageSize element contains the number of items to be returned in a single page for a search result.
	PageSize *PageSize `xml:"PageSize"`
}

func (G *GetNonIndexableItemDetails) SetForMarshal() {
	G.XMLName.Local = "m:GetNonIndexableItemDetails"
}

func (G *GetNonIndexableItemDetails) GetSchema() *Schema {
	return &SchemaMessages
}
