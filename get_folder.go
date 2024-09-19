package goews

import "encoding/xml"

type GetFolder struct {
	XMLName xml.Name `xml:"m:GetFolder"`

	FolderIds   *FolderIds   `xml:"m:FolderIds,omitempty"`
	FolderShape *FolderShape `xml:"m:FolderShape,omitempty"`
}

type FolderIds struct {
	DistinguishedFolderId *DistinguishedFolderId `xml:"t:DistinguishedFolderId,omitempty"`
	FolderId              *FolderId              `xml:"t:FolderId,omitempty"`
}

type FolderId struct {
	ChangeKey *string `xml:"ChangeKey,attr"`
	Id        *string `xml:"Id,attr"`
}

type FolderShape struct {
	// The AdditionalProperties element identifies additional properties for use in GetItem, UpdateItem, CreateItem, FindItem, or FindFolder requests.
	AdditionalProperties *AdditionalProperties `xml:"t:AdditionalProperties,omitempty"`
	// The BaseShape element identifies the set of properties to return in an item or folder response.
	BaseShape BaseShape `xml:"t:BaseShape,omitempty"`
}

type getFolderResponseEnvelop struct {
	XMLName struct{}              `xml:"Envelope"`
	Body    getFolderResponseBody `xml:"Body"`
}
type getFolderResponseBody struct {
	GetFolderResponse GetFolderResponse `xml:"GetFolderResponse"`
}

type GetFolderResponse struct {
	ResponseMessages *GetFolderResponseMessages `xml:"m:ResponseMessages"`
}

type GetFolderResponseMessages struct {
	GetFolderResponseMessage *GetFolderResponseMessage `xml:"m:GetFolderResponseMessage"`
}

type GetFolderResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"m:DescriptiveLinkKey,omitempty"`
	// The Folders element contains an array of folders that are used in folder operations.
	Folders *Folders `xml:"t:Folders,omitempty"`
	// The MessageText element provides a text description of the status of the response.
	MessageText MessageText `xml:"m:MessageText,omitempty"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"m:MessageXml,omitempty"`
	// The ResponseCode element provides status information about the request.
	ResponseCode ResponseCode `xml:"m:ResponseCode"`
	// Describes the status of a GetFolder operation response. The following values are valid for this attribute:  - Success  - Warning  - Error
	ResponseClass string `xml:"ResponseClass,attr"`
}

type DescriptiveLinkKey struct {
	TEXT int64 `xml:",chardata"`
}

type MessageText string

type ResponseCode string

type Folders struct {
	// The CalendarFolder element represents a folder that primarily contains calendar items.
	CalendarFolder *CalendarFolder `xml:"t:CalendarFolder,omitempty"`
	// The ContactsFolder element represents a contacts folder that is contained in a mailbox.
	ContactsFolder *ContactsFolder `xml:"t:ContactsFolder,omitempty"`
	// The Folder element defines a folder to create, get, find, synchronize, or update.
	Folder *Folder `xml:"t:Folder,omitempty"`
	// The SearchFolder element represents a search folder that is contained in a mailbox.
	SearchFolder *SearchFolder `xml:"t:SearchFolder,omitempty"`
	// The TasksFolder element represents a Tasks folder that is contained in a mailbox.
	TasksFolder *TasksFolder `xml:"t:TasksFolder,omitempty"`
}

type CalendarFolder struct {
	// The ChildFolderCount element represents the number of immediate child folders that are contained within a folder. This property is read-only.
	ChildFolderCount ChildFolderCount `xml:"t:ChildFolderCount,omitempty"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName DisplayNamestring `xml:"t:DisplayName,omitempty"`
	// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	EffectiveRights *EffectiveRights `xml:"t:EffectiveRights,omitempty"`
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"t:ExtendedProperty,omitempty"`
	// The FolderClass element represents the folder class for a folder.
	FolderClass FolderClass `xml:"t:FolderClass,omitempty"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId,omitempty"`
	// The ManagedFolderInformation element contains information about a managed custom folder.
	ManagedFolderInformation *ManagedFolderInformation `xml:"t:ManagedFolderInformation,omitempty"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId,omitempty"`
	// The PermissionSet element contains all the permissions that are configured for a calendar folder.
	PermissionSet *PermissionSetCalendarPermissionSetType `xml:"t:PermissionSet,omitempty"`
	// The SharingEffectiveRights element indicates the permissions that the user has for the calendar data that is being shared.
	SharingEffectiveRights SharingEffectiveRightsCalendarPermissionReadAccessType `xml:"t:SharingEffectiveRights,omitempty"`
	// The TotalCount element represents the total count of items within a given folder.
	TotalCount TotalCount `xml:"t:TotalCount"`
}

type ChildFolderCount int64

type DisplayNamestring string

type EffectiveRights struct {
	// // The CreateAssociated element indicates whether a client can create an associated contents table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// CreateAssociated *CreateAssociated `xml:"CreateAssociated"`
	// // The CreateContents element indicates whether a client can create a contents table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// CreateContents *CreateContents `xml:"CreateContents"`
	// // The CreateHierarchy element indicates whether a client can create a hierarchy table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// CreateHierarchy *CreateHierarchy `xml:"CreateHierarchy"`
	// // The Delete element indicates whether a client can delete a folder or item.
	// Delete *Delete `xml:"Delete"`
	// // The Modify element indicates whether a client can modify a folder or item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// Modify *Modify `xml:"Modify"`
	// // The Read element indicates whether a client can read a folder or item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// Read *Read `xml:"Read"`
	// // The ViewPrivateItems element indicates whether a delegate user or client application has permission to view private items in the principal's mailbox.
	// ViewPrivateItems *ViewPrivateItems `xml:"ViewPrivateItems"`
}

type ExtendedProperty struct {
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI,omitempty"`
	// The Value element contains the value of an extended property.
	Value string `xml:"t:Value,omitempty"`
	// The Values element contains a collection of values for an extended property.
	Values *Values `xml:"t:Values,omitempty"`
}

type ExtendedFieldURI struct {
	// Defines the well-known property set IDs for extended MAPI properties.If this attribute is used, the PropertySetId and PropertyTag attributes cannot be used. This attribute must be used with either the PropertyId or PropertyName attribute, and the PropertyType attribute.The DistinguishedPropertySetId Attribute table later in this topic lists the possible values for this attribute.This attribute is optional.
	DistinguishedPropertySetId string `xml:"DistinguishedPropertySetId,attr"`
	// Identifies an extended property by its dispatch ID. The dispatch ID can be identified in either decimal or hexadecimal formats. This property must be coupled with either DistinguishedPropertySetId or PropertySetId.If this attribute is used, the PropertyName and PropertyTag attributes cannot be used.This attribute is optional.
	PropertyId string `xml:"PropertyId,attr"`
	// Identifies an extended property by its name. This property must be coupled with either DistinguishedPropertySetId or PropertySetId.If this attribute is used, the PropertyId and PropertyTag attributes cannot be used.This attribute is optional.
	PropertyName string `xml:"PropertyName,attr"`
	// Identifies a MAPI extended property set or namespace by its identifying GUID.If this attribute is used, the DistinguishedPropertySetId and PropertyTag attribute cannot be used. This attribute must be used with either the PropertyId or PropertyName attribute, and the PropertyType attribute.This attribute is optional.
	PropertySetId string `xml:"PropertySetId,attr"`
	// Identifies the property tag without the type part of the tag. The PropertyTag can be represented as either a hexadecimal or a short integer.The range between 0x8000 and 0xFFFE represents the custom range of properties. When a mailbox database encounters a custom property for the first time, it assigns that custom property a property tag within the custom property range of 0x8000-0xFFFE. A given custom property tag will most likely differ across databases. Therefore, a custom property request by property tag can return different properties on different databases. The use of the PropertyTag attribute is prohibited for custom properties. Instead, use the PropertySetId attribute and the PropertyName or PropertyId attribute.IMPORTANT: Access any custom property between 0x8000 and 0xFFFE by using the GUID + name/ID. If the PropertyTag attribute is used, the DistinguishedPropertySetId, PropertySetId, PropertyName, and PropertyId attributes cannot be used.This attribute is optional.NOTE: You cannot use a property tag attribute for properties within the custom range 0x8000-0xFFFE. You must use a named property in this case.
	PropertyTag string `xml:"PropertyTag,attr"`
	// Represents the property type of a property tag. This corresponds to the least significant word in a property tag.The PropertyType Attribute table later in this topic contains the possible values for this attribute.This attribute is required.
	PropertyType string `xml:"PropertyType,attr"`
}

type Values struct {
	// The Value element contains the value of an extended property.
	Value string `xml:"t:Value,omitempty"`
}

type FolderClass string

type ManagedFolderInformation struct {
	// // The CanDelete element indicates whether a managed folder can be deleted by a customer.
	// CanDelete *CanDelete `xml:"CanDelete"`
	// // The CanRenameOrMove element indicates whether a managed folder can be renamed or moved by the customer.
	// CanRenameOrMove *CanRenameOrMove `xml:"CanRenameOrMove"`
	// // The Comment element contains the comment that is associated with a managed folder.
	// Comment *Comment `xml:"Comment"`
	// // The FolderSize element describes the total size of all the contents of a managed folder.
	// FolderSize *FolderSize `xml:"FolderSize"`
	// // The HasQuota element indicates whether the managed folder has a quota.
	// HasQuota *HasQuota `xml:"HasQuota"`
	// // The HomePage element specifies the URL that will be the default home page for the managed folder.
	// HomePage *HomePage `xml:"HomePage"`
	// // The IsManagedFoldersRoot element indicates whether the managed folder is the root for all managed folders.
	// IsManagedFoldersRoot *IsManagedFoldersRoot `xml:"IsManagedFoldersRoot"`
	// // The ManagedFolderId element contains the folder ID of the managed folder.
	// ManagedFolderId *ManagedFolderId `xml:"ManagedFolderId"`
	// // The MustDisplayComment element indicates whether the managed folder comment must be displayed.
	// MustDisplayComment *MustDisplayComment `xml:"MustDisplayComment"`
	// // The StorageQuota element describes the storage quota for the managed folder.
	// StorageQuota *StorageQuota `xml:"StorageQuota"`
}

type PermissionSetCalendarPermissionSetType struct {
	// The CalendarPermissions element contains an array of calendar permissions for a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CalendarPermissions *CalendarPermissions `xml:"t:CalendarPermissions,omitempty"`
	// The UnknownEntries element contains an array of unknown permission entries that cannot be resolved against the Active Directory directory service. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	UnknownEntries *UnknownEntries `xml:"t:UnknownEntries,omitempty"`
}

type CalendarPermissions struct {
	// The CalendarPermission element defines the access that a user has to a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CalendarPermission *CalendarPermission `xml:"t:CalendarPermission,omitempty"`
}

type CalendarPermission struct {
	// // The CalendarPermissionLevel element represents the permission level that a user has on a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// CalendarPermissionLevel *CalendarPermissionLevel `xml:"CalendarPermissionLevel"`
	// // The CanCreateItems element indicates whether a user has permission to create items in a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// CanCreateItems *CanCreateItems `xml:"CanCreateItems"`
	// // The CanCreateSubFolders element indicates whether a user has permission to create subfolders in a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// CanCreateSubFolders *CanCreateSubFolders `xml:"CanCreateSubFolders"`
	// // The DeleteItems element indicates which items in a folder a user has permission to delete. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// DeleteItems *DeleteItems `xml:"DeleteItems"`
	// // The EditItems element indicates which items in a folder a user has permission to edit. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// EditItems *EditItems `xml:"EditItems"`
	// // The IsFolderContact element indicates whether a user is a contact for a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// IsFolderContact *IsFolderContact `xml:"IsFolderContact"`
	// // The IsFolderOwner element indicates whether a user is the owner of a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// IsFolderOwner *IsFolderOwner `xml:"IsFolderOwner"`
	// // The IsFolderVisible element indicates whether a user can view a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// IsFolderVisible *IsFolderVisible `xml:"IsFolderVisible"`
	// // The ReadItems element indicates whether a user has permission to read items within a Calendar folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// ReadItems *ReadItemsCalendarPermissionType `xml:"ReadItems"`
	// // The UserId element identifies a delegate user or a user who has folder access permissions.
	// UserId *UserId `xml:"UserId"`
}

type UnknownEntries struct {
	// The UnknownEntry element represents a single unknown permission entry that cannot be resolved against the Active Directory directory service. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	UnknownEntry UnknownEntry `xml:"t:UnknownEntry,omitempty"`
}

type UnknownEntry string

type SharingEffectiveRightsCalendarPermissionReadAccessType string

const (
	// Indicates that the user has permission to view all items in the calendar, including free/busy time and subject, location, and details of appointments.
	SharingEffectiveRightsCalendarPermissionReadAccessTypeFullDetails string = `FullDetails`
	// Indicates that the user does not have permission to view items in the calendar.
	SharingEffectiveRightsCalendarPermissionReadAccessTypeNone string = `None`
	// Indicates that the user has permission to view free/busy time in the calendar and the subject and location of appointments.
	SharingEffectiveRightsCalendarPermissionReadAccessTypeTimeAndSubjectAndLocation string = `TimeAndSubjectAndLocation`
	// Indicates that the user has permission to view only free/busy time in the calendar.
	SharingEffectiveRightsCalendarPermissionReadAccessTypeTimeOnly string = `TimeOnly`
)

type TotalCount int64

type ContactsFolder struct {

	// The ChildFolderCount element represents the number of immediate child folders that are contained within a folder. This property is read-only.
	ChildFolderCount ChildFolderCount `xml:"t:ChildFolderCount,omitempty"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName DisplayNamestring `xml:"t:DisplayName,omitempty"`
	// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	EffectiveRights *EffectiveRights `xml:"t:EffectiveRights,omitempty"`
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"t:ExtendedProperty,omitempty"`
	// The FolderClass element represents the folder class for a folder.
	FolderClass FolderClass `xml:"t:FolderClass,omitempty"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId,omitempty"`
	// The ManagedFolderInformation element contains information about a managed custom folder.
	ManagedFolderInformation *ManagedFolderInformation `xml:"t:ManagedFolderInformation,omitempty"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId,omitempty"`
	// The PermissionSet element contains all the permissions that are configured for a folder.
	PermissionSet *PermissionSetPermissionSetType `xml:"t:PermissionSet,omitempty"`
	// The SharingEffectiveRights element indicates the permissions that the user has for the contact data that is being shared.
	SharingEffectiveRights SharingEffectiveRightsPermissionReadAccessType `xml:"t:SharingEffectiveRights,omitempty"`
	// The TotalCount element represents the total count of items within a given folder.
	TotalCount TotalCount `xml:"t:TotalCount"`
}

type PermissionSetPermissionSetType struct {
	// The Permissions element contains the collection of permissions for a folder.
	Permissions *Permissions `xml:"t:Permissions,omitempty"`
	// The UnknownEntries element contains an array of unknown permission entries that cannot be resolved against the Active Directory directory service. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	UnknownEntries *UnknownEntries `xml:"t:UnknownEntries,omitempty"`
}

type SharingEffectiveRightsPermissionReadAccessType string

const (
	// FullDetails
	SharingEffectiveRightsPermissionReadAccessTypeFullDetails string = `FullDetails`
	// None
	SharingEffectiveRightsPermissionReadAccessTypeNone string = `None`
)

type Permissions struct {
	// The Permission element defines the access that a user has to a folder.
	Permission *Permission `xml:"t:Permission,omitempty"`
}

type Permission struct {
	// // The CanCreateItems element indicates whether a user has permission to create items in a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// CanCreateItems *CanCreateItems `xml:"CanCreateItems"`
	// // The CanCreateSubFolders element indicates whether a user has permission to create subfolders in a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// CanCreateSubFolders *CanCreateSubFolders `xml:"CanCreateSubFolders"`
	// // The DeleteItems element indicates which items in a folder a user has permission to delete. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// DeleteItems *DeleteItems `xml:"DeleteItems"`
	// // The EditItems element indicates which items in a folder a user has permission to edit. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// EditItems *EditItems `xml:"EditItems"`
	// // The IsFolderContact element indicates whether a user is a contact for a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// IsFolderContact *IsFolderContact `xml:"IsFolderContact"`
	// // The IsFolderOwner element indicates whether a user is the owner of a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// IsFolderOwner *IsFolderOwner `xml:"IsFolderOwner"`
	// // The IsFolderVisible element indicates whether a user can view a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// IsFolderVisible *IsFolderVisible `xml:"IsFolderVisible"`
	// // The PermissionLevel element represents the permission level that a user has on a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// PermissionLevel *PermissionLevel `xml:"PermissionLevel"`
	// // The ReadItems element indicates whether a user has permission to read items within a folder. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// ReadItems *ReadItemsPermissionType `xml:"ReadItems"`
	// // The UserId element identifies a delegate user or a user who has folder access permissions.
	// UserId *UserId `xml:"UserId"`
}

type Folder struct {
	// The ChildFolderCount element represents the number of immediate child folders that are contained within a folder. This property is read-only.
	ChildFolderCount ChildFolderCount `xml:"t:ChildFolderCount,omitempty"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName DisplayNamestring `xml:"t:DisplayName,omitempty"`
	// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	EffectiveRights *EffectiveRights `xml:"t:EffectiveRights,omitempty"`
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"t:ExtendedProperty,omitempty"`
	// The FolderClass element represents the folder class for a folder.
	FolderClass FolderClass `xml:"t:FolderClass,omitempty"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId,omitempty"`
	// The ManagedFolderInformation element contains information about a managed custom folder.
	ManagedFolderInformation *ManagedFolderInformation `xml:"t:ManagedFolderInformation,omitempty"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId,omitempty"`
	// The PermissionSet element contains all the permissions that are configured for a folder.
	PermissionSet *PermissionSetPermissionSetType `xml:"t:PermissionSet,omitempty"`
	// The TotalCount element represents the total count of items within a given folder.
	TotalCount TotalCount `xml:"t:TotalCount"`
	// The UnreadCount element contains the count of unread items within a folder.
	UnreadCount UnreadCount `xml:"t:UnreadCount"`
}

type UnreadCount int64

type SearchFolder struct {
	// The ChildFolderCount element represents the number of immediate child folders that are contained within a folder. This property is read-only.
	ChildFolderCount ChildFolderCount `xml:"t:ChildFolderCount,omitempty"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName DisplayNamestring `xml:"t:DisplayName,omitempty"`
	// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	EffectiveRights *EffectiveRights `xml:"t:EffectiveRights,omitempty"`
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"t:ExtendedProperty,omitempty"`
	// The FolderClass element represents the folder class for a folder.
	FolderClass FolderClass `xml:"t:FolderClass,omitempty"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId,omitempty"`
	// The ManagedFolderInformation element contains information about a managed custom folder.
	ManagedFolderInformation *ManagedFolderInformation `xml:"t:ManagedFolderInformation,omitempty"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId,omitempty"`
	// The PermissionSet element contains all the permissions that are configured for a folder.
	PermissionSet *PermissionSetPermissionSetType `xml:"t:PermissionSet,omitempty"`
	// The SearchParameters element represents the parameters that define a search folder.
	SearchParameters *SearchParameters `xml:"t:SearchParameters,omitempty"`
	// The TotalCount element represents the total count of items within a given folder.
	TotalCount TotalCount `xml:"t:TotalCount"`
	// The UnreadCount element contains the count of unread items within a folder.
	UnreadCount UnreadCount `xml:"t:UnreadCount"`
}

type SearchParameters struct {
	// The BaseFolderIds element represents the collection of folders that will be mined to determine the contents of a search folder.
	BaseFolderIds *BaseFolderIds `xml:"t:BaseFolderIds,omitempty"`
	// The Restriction element represents the restriction or query that is used to filter items or folders in FindItem/FindFolder and search folder operations.
	Restriction *Restriction `xml:"m:Restriction,omitempty"`
	// Describes how a search folder traverses the folder hierarchy. The options are for either a Deep or a Shallow search.
	Traversal string `xml:"Traversal,attr"`
}

type BaseFolderIds struct {
	// The DistinguishedFolderId element identifies folders that can be referenced by name. If you do not use this element, you must use the FolderId element to identify a folder.
	DistinguishedFolderId *DistinguishedFolderId `xml:"t:DistinguishedFolderId,omitempty"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId,omitempty"`
}

type TasksFolder struct {
	// The ChildFolderCount element represents the number of immediate child folders that are contained within a folder. This property is read-only.
	ChildFolderCount ChildFolderCount `xml:"t:ChildFolderCount,omitempty"`
	// The DisplayName element defines the display name of a folder, contact, distribution list, delegate user, location, or rule.
	DisplayName DisplayNamestring `xml:"t:DisplayName,omitempty"`
	// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
	EffectiveRights *EffectiveRights `xml:"t:EffectiveRights,omitempty"`
	// The ExtendedProperty element identifies extended MAPI properties on folders and items.
	ExtendedProperty *ExtendedProperty `xml:"t:ExtendedProperty,omitempty"`
	// The FolderClass element represents the folder class for a folder.
	FolderClass FolderClass `xml:"t:FolderClass,omitempty"`
	// The FolderId element contains the identifier and change key of a folder.
	FolderId *FolderId `xml:"t:FolderId,omitempty"`
	// The ManagedFolderInformation element contains information about a managed custom folder.
	ManagedFolderInformation *ManagedFolderInformation `xml:"t:ManagedFolderInformation,omitempty"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId,omitempty"`
	// The PermissionSet element contains all the permissions that are configured for a folder.
	PermissionSet *PermissionSetPermissionSetType `xml:"t:PermissionSet,omitempty"`
	// The TotalCount element represents the total count of items within a given folder.
	TotalCount TotalCount `xml:"t:TotalCount"`
	// The UnreadCount element contains the count of unread items within a folder.
	UnreadCount UnreadCount `xml:"t:UnreadCount"`
}
