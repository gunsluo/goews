package schema

// The BaseFolderIds element represents the collection of folders that will be mined to determine the contents of a search folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/basefolderids
import "encoding/xml"

type BaseFolderIds struct {
	XMLName xml.Name

	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"DistinguishedFolderId"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
}

func (B *BaseFolderIds) SetForMarshal() {
	B.XMLName.Local = "t:BaseFolderIds"
}

func (B *BaseFolderIds) GetSchema() *Schema {
	return &SchemaTypes
}
