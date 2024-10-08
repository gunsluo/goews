package schema

// The AlternatePublicFolderId element describes a public folder identifier to convert to another identifier format. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/alternatepublicfolderid
import "encoding/xml"

type AlternatePublicFolderId struct {
	XMLName xml.Name

	// Contains the public folder identifier to convert. This attribute is required.
	FolderId *string `xml:"FolderId,attr"`
	// Identifies the format that describes the public folder identifier to convert. This attribute is required.
	Format *string `xml:"Format,attr"`
}

func (A *AlternatePublicFolderId) SetForMarshal() {
	A.XMLName.Local = "t:AlternatePublicFolderId"
}

func (A *AlternatePublicFolderId) GetSchema() *Schema {
	return &SchemaTypes
}
