package schema

// The ContactsFolder element represents a contacts folder that is contained in a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contactsfolder
import "encoding/xml"

type ContactsFolder struct {
	XMLName xml.Name

	// The ChildFolderCount element represents the number of immediate child folders that are contained within a folder. This property is read-only.
	ChildFolderCount *ChildFolderCount `xml:"ChildFolderCount"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName *DisplayNamestring `xml:"DisplayName"`
	// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	EffectiveRights *EffectiveRights `xml:"EffectiveRights"`
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"ExtendedProperty"`
	// The FolderClass element represents the folder class for a folder.
	FolderClass *FolderClass `xml:"FolderClass"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"FolderId"`
	// The ManagedFolderInformation element contains information about a managed custom folder.
	ManagedFolderInformation *ManagedFolderInformation `xml:"ManagedFolderInformation"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"ParentFolderId"`
	// The PermissionSet element contains all the permissions that are configured for a folder.
	PermissionSet *PermissionSetPermissionSetType `xml:"PermissionSet"`
	// The SharingEffectiveRights element indicates the permissions that the user has for the contact data that is being shared.
	SharingEffectiveRights *SharingEffectiveRightsPermissionReadAccessType `xml:"SharingEffectiveRights"`
	// The TotalCount element represents the total count of items within a given folder.
	TotalCount *TotalCount `xml:"TotalCount"`
}

func (C *ContactsFolder) SetForMarshal() {
	C.XMLName.Local = "t:ContactsFolder"
}

func (C *ContactsFolder) GetSchema() *Schema {
	return &SchemaTypes
}
