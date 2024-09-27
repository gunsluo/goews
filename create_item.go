package goews

import (
	"time"
)

type CreateItem struct {
	XMLName                struct{}          `xml:"m:CreateItem"`
	MessageDisposition     string            `xml:"MessageDisposition,attr"`
	SendMeetingInvitations string            `xml:"SendMeetingInvitations,attr"`
	SavedItemFolderId      SavedItemFolderId `xml:"m:SavedItemFolderId"`
	Items                  Items             `xml:"m:Items"`
}

type Items struct {
	Message      []Message      `xml:"t:Message"`
	CalendarItem []CalendarItem `xml:"t:CalendarItem"`
}

type SavedItemFolderId struct {
	DistinguishedFolderId DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
}

type Message struct {
	ItemClass     string       `xml:"t:ItemClass,omitempty"`
	ItemId        *ItemId      `xml:"t:ItemId,omitempty"`
	Subject       string       `xml:"t:Subject"`
	Body          Body         `xml:"t:Body"`
	Sender        OneMailbox   `xml:"t:Sender"`
	ToRecipients  XMailbox     `xml:"t:ToRecipients"`
	CcRecipients  *XMailbox    `xml:"t:CcRecipients,omitempty"`
	BccRecipients *XMailbox    `xml:"t:BccRecipients,omitempty"`
	Attachments   *Attachments `xml:"t:Attachments,omitempty"`
}

type CalendarItem struct {
	Subject                    string      `xml:"t:Subject"`
	Body                       Body        `xml:"t:Body"`
	ReminderIsSet              bool        `xml:"t:ReminderIsSet"`
	ReminderMinutesBeforeStart int         `xml:"t:ReminderMinutesBeforeStart"`
	Start                      time.Time   `xml:"t:Start"`
	End                        time.Time   `xml:"t:End"`
	IsAllDayEvent              bool        `xml:"t:IsAllDayEvent"`
	LegacyFreeBusyStatus       string      `xml:"t:LegacyFreeBusyStatus"`
	Location                   string      `xml:"t:Location"`
	RequiredAttendees          []Attendees `xml:"t:RequiredAttendees"`
	OptionalAttendees          []Attendees `xml:"t:OptionalAttendees"`
	Resources                  []Attendees `xml:"t:Resources"`
}

type Body struct {
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

type Attachments struct {
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
