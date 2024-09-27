package goews

import "encoding/xml"

type GetItem struct {
	XMLName xml.Name `xml:"m:GetItem"`

	// The ItemIds element contains the unique identities of items, occurrence items, and recurring master items that are used to delete, send, get, move, or copy items in the Exchange store.
	ItemIds *ItemIds `xml:"m:ItemIds,omitempty"`
	// The ItemShape element identifies a set of properties to return in a GetItem operation, FindItem operation, or SyncFolderItems operation response.
	ItemShape *ItemShape `xml:"m:ItemShape,omitempty"`
}

type ItemIds struct {
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId,omitempty"`
	// The OccurrenceItemId element identifies a single occurrence of a recurring item.
	OccurrenceItemId *OccurrenceItemId `xml:"t:OccurrenceItemId,omitempty"`
	// The RecurringMasterItemId element identifies a recurrence master item by identifying the identifiers of one of its related occurrence items.
	RecurringMasterItemId *RecurringMasterItemId `xml:"t:RecurringMasterItemId,omitempty"`
}

type ItemId struct {
	// Identifies a specific version of an item. A ChangeKey is required for the following scenarios:  - The UpdateItem element requires a ChangeKey if the ConflictResolution attribute is set to AutoResolve. AutoResolve is a default value. If the ChangeKey attribute is not included, the response will return a ResponseCode value equal to ErrorChangeKeyRequired.  - The SendItem element requires a ChangeKey to test whether the attempted operation will act upon the most recent version of an item. If the ChangeKey attribute is not included in the ItemId or if the ChangeKey is empty, the response will return a ResponseCode value equal to ErrorStaleObject.
	ChangeKey string `xml:"ChangeKey,attr,omitempty"`
	// Identifies a specific item in the Exchange store. Id is case-sensitive; therefore, comparisons between Ids must be case-sensitive or binary.
	Id string `xml:"Id,attr,omitempty"`
}

type OccurrenceItemId struct {
	// Identifies a specific version of the recurring master or an item occurrence. If either the recurring master or any of its occurrences change, the ChangeKey changes. The ChangeKey is the same for the recurring master and all occurrences.
	ChangeKey string `xml:"ChangeKey,attr,omitempty"`
	// Identifies the index of the item occurrence. This attribute is required. This value represents an integer.
	InstanceIndex string `xml:"InstanceIndex,attr,omitempty"`
	// Identifies the recurring master of a recurring item. This attribute is required.
	RecurringMasterId string `xml:"RecurringMasterId,attr,omitempty"`
}

type RecurringMasterItemId struct {
	// Identifies a specific version of a single occurrence of a recurring master item. Additionally, the recurring master item is also identified because it and the single occurrence will contain the same change key. This attribute is optional.
	ChangeKey string `xml:"ChangeKey,attr,omitempty"`
	// Identifies a single occurrence of a recurring master item. This attribute is required.
	OccurrenceId string `xml:"OccurrenceId,attr,omitempty"`
}

type getGetItemResponseEnvelop struct {
	XMLName struct{}               `xml:"Envelope"`
	Body    getGetItemResponseBody `xml:"Body"`
}
type getGetItemResponseBody struct {
	GetItemResponse GetItemResponse `xml:"GetItemResponse"`
}

type GetItemResponse struct {
	// The GetItemResponseMessage element contains the status and result of a single GetItem operation request.
	ResponseMessages *GetItemResponseMessage `xml:"GetItemResponseMessage"`
}

type GetItemResponseMessage struct {
	// The DescriptiveLinkKey element is currently unused and is reserved for future use. It contains a value of 0.
	DescriptiveLinkKey *DescriptiveLinkKey `xml:"DescriptiveLinkKey,omitempty"`
	// The Items element contains an array of items.
	Items *Items `xml:"Items,omitempty"`
	// The MessageText element provides a text description of the status of the response.
	MessageText MessageText `xml:"MessageText,omitempty"`
	// The MessageXml element provides additional error response information.
	MessageXml *MessageXml `xml:"MessageXml,omitempty"`
	// The ResponseCode element provides status information about the request.
	ResponseCode ResponseCode `xml:"ResponseCode"`
	// Describes the status of a GetItem operation response. The following values are valid for this attribute:- Success- Warning- Error
	ResponseClass string `xml:"ResponseClass,attr"`
}
