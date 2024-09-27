package goews

import "encoding/xml"

type CreateItem struct {
	XMLName xml.Name `xml:"m:CreateItem"`

	// The Items element contains a set of items to create.
	Items *ItemsNonEmptyArrayOfAllItemsType `xml:"m:Items"`
	// The SavedItemFolderId element identifies the target folder for operations that update, send, and create items in a mailbox.
	SavedItemFolderId *SavedItemFolderId `xml:"m:SavedItemFolderId"`
	// Describes how the item will be handled after it is created. The attribute is required for e-mail messages. This attribute is only applicable to e-mail messages.
	MessageDisposition string `xml:"MessageDisposition,attr,omitempty"`
	// Describes how meeting requests are handled after they are created. This attribute is required for calendar items.
	SendMeetingInvitations string `xml:"SendMeetingInvitations,attr,omitempty"`
}

type ItemsNonEmptyArrayOfAllItemsType struct {
	// // The AcceptItem element represents an Accept reply to a meeting request.
	// AcceptItem *AcceptItem `xml:"AcceptItem"`
	// // The AcceptSharingInvitation element is used to accept an invitation that allows access to another user's calendar or contacts data.
	// AcceptSharingInvitation *AcceptSharingInvitation `xml:"AcceptSharingInvitation"`
	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem *CalendarItem `xml:"t:CalendarItem"`
	// // The CancelCalendarItem element represents the response object that is used to cancel a meeting.
	// CancelCalendarItem *CancelCalendarItem `xml:"CancelCalendarItem"`
	// // The Contact element represents a contact item in the Exchange store.
	// Contact *Contact `xml:"Contact"`
	// // The DeclineItem element represents a Decline reply to a meeting request.
	// DeclineItem *DeclineItem `xml:"DeclineItem"`
	// // The DistributionList element represents a distribution list.
	// DistributionList *DistributionList `xml:"DistributionList"`
	// // The ForwardItem element contains an Exchange store item to forward to recipients.
	// ForwardItem *ForwardItem `xml:"ForwardItem"`
	// // The Item element represents a generic item in the Exchange store.
	// Item *Item `xml:"Item"`
	// // The MeetingCancellation element represents a meeting cancellation in the Exchange store.
	// MeetingCancellation *MeetingCancellation `xml:"MeetingCancellation"`
	// // The MeetingMessage element represents a meeting in the Exchange store.
	// MeetingMessage *MeetingMessage `xml:"MeetingMessage"`
	// // The MeetingRequest element represents a meeting request in the Exchange store.
	// MeetingRequest *MeetingRequest `xml:"MeetingRequest"`
	// // The MeetingResponse element represents a meeting response in the Exchange store.
	// MeetingResponse *MeetingResponse `xml:"MeetingResponse"`
	// The Message element represents a Microsoft Exchange e-mail message.
	Message *TMessage `xml:"t:Message"`
	// // The PostReplyItem element contains a reply to a post item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	// PostReplyItem *PostReplyItem `xml:"PostReplyItem"`
	// // The RemoveItem element represents a response object that is used to remove a meeting item when a MeetingCancellation message is received.
	// RemoveItem *RemoveItem `xml:"RemoveItem"`
	// // The ReplyToAllItem element contains a reply to the sender and all identified recipients of an item in the Exchange store.
	// ReplyAllToItem *ReplyAllToItem `xml:"ReplyAllToItem"`
	// // The ReplyToItem element contains a reply to the sender of an item in the Exchange store.
	// ReplyToItem *ReplyToItem `xml:"ReplyToItem"`
	// // The SuppressReadReceipt element is used to suppress read receipts.
	// SuppressReadReceipt *SuppressReadReceipt `xml:"SuppressReadReceipt"`
	// // The Task element represents a task in the Exchange store.
	// Task *Task `xml:"Task"`
	// // The TentativelyAcceptItem element represents a Tentative reply to a meeting request.
	// TentativelyAcceptItem *TentativelyAcceptItem `xml:"TentativelyAcceptItem"`
}

type CalendarItem struct {
}

type SavedItemFolderId struct {
	DistinguishedFolderId DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
}

type TMessage struct {
	ItemClass     string        `xml:"t:ItemClass,omitempty"`
	ItemId        *ItemId       `xml:"t:ItemId,omitempty"`
	Subject       string        `xml:"t:Subject"`
	Body          TBody         `xml:"t:Body"`
	Sender        OneMailbox    `xml:"t:Sender"`
	ToRecipients  XMailbox      `xml:"t:ToRecipients"`
	CcRecipients  *XMailbox     `xml:"t:CcRecipients,omitempty"`
	BccRecipients *XMailbox     `xml:"t:BccRecipients,omitempty"`
	Attachments   *TAttachments `xml:"t:Attachments,omitempty"`
}

type TBody struct {
	BodyType string `xml:"BodyType,attr"`
	Body     []byte `xml:",chardata"`
}

type OneMailbox struct {
	Mailbox Mailbox `xml:"t:Mailbox"`
}

type XMailbox struct {
	Mailbox []Mailbox `xml:"t:Mailbox"`
}

type Mailbox struct {
	EmailAddress string `xml:"t:EmailAddress"`
}

type Attendee struct {
	Mailbox Mailbox `xml:"t:Mailbox"`
}

type Attendees struct {
	Attendee []Attendee `xml:"t:Attendee"`
}

type TAttachments struct {
	Items []ItemAttachment `xml:"t:ItemAttachment,omitempty"`
	Files []FileAttachment `xml:"t:FileAttachment,omitempty"`
}

type ItemAttachment struct {
	AttachmentId     AttachmentId `xml:"t:AttachmentId,omitempty"`
	Name             string       `xml:"t:Name"`
	ContentType      string       `xml:"t:ContentType,omitempty"`
	ContentId        string       `xml:"t:ContentId,omitempty"`
	ContentLocation  string       `xml:"t:ContentLocation,omitempty"`
	Size             int64        `xml:"t:Size,omitempty"`
	LastModifiedTime string       `xml:"t:LastModifiedTime,omitempty"`
	IsInline         bool         `xml:"t:IsInline"`

	// TODO:
	// Item
	// Message
	// CalendarItem
	// Contact
	// Task
	// MeetingMessage
	// MeetingResponse
	// MeetingCancellation
}

type FileAttachment struct {
	AttachmentId     AttachmentId `xml:"t:AttachmentId,omitempty"`
	Name             string       `xml:"t:Name"`
	ContentType      string       `xml:"t:ContentType,omitempty"`
	ContentId        string       `xml:"t:ContentId,omitempty"`
	ContentLocation  string       `xml:"t:ContentLocation,omitempty"`
	Size             int64        `xml:"t:Size,omitempty"`
	LastModifiedTime string       `xml:"t:LastModifiedTime,omitempty"`
	IsInline         bool         `xml:"t:IsInline"`
	IsContactPhoto   bool         `xml:"t:IsContactPhoto"`
	Content          string       `xml:"t:Content"`
}

type AttachmentId struct {
	Id                string `xml:"Id,attr"`
	RootItemId        string `xml:"RootItemId,attr"`
	RootItemChangeKey string `xml:"RootItemChangeKey,attr"`
}

type createItemResponseBodyEnvelop struct {
	XMLName struct{}               `xml:"Envelope"`
	Body    createItemResponseBody `xml:"Body"`
}
type createItemResponseBody struct {
	CreateItemResponse CreateItemResponse `xml:"CreateItemResponse"`
}

type CreateItemResponse struct {
	ResponseMessages ResponseMessages `xml:"ResponseMessages"`
}

type ResponseMessages struct {
	CreateItemResponseMessage Response `xml:"CreateItemResponseMessage"`
}
