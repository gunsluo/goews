package schema

// The Notification element contains information about the subscription and the events that have occurred since the last notification.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/notification-ex15websvcsotherref
import "encoding/xml"

type Notification struct {
	XMLName xml.Name

	// The CopiedEvent element represents an event in which an item or folder is copied.
	CopiedEvent *CopiedEvent `xml:"CopiedEvent"`
	// The CreatedEvent element represents an event in which an item or folder is created.
	CreatedEvent *CreatedEvent `xml:"CreatedEvent"`
	// The DeletedEvent element represents an event in which an item or folder is deleted.
	DeletedEvent *DeletedEvent `xml:"DeletedEvent"`
	// The FreeBusyChangedEvent element represents an event in which an item's free/busy time has changed.
	FreeBusyChangedEvent *FreeBusyChangedEvent `xml:"FreeBusyChangedEvent"`
	// The ModifiedEvent element represents an event in which an item or folder is modified.
	ModifiedEvent *ModifiedEvent `xml:"ModifiedEvent"`
	// The MoreEvents element indicates whether there are more events in the queue to be delivered to the client.
	MoreEvents *MoreEvents `xml:"MoreEvents"`
	// The MovedEvent element represents an event in which an item or folder is moved from one parent folder to another parent folder.
	MovedEvent *MovedEvent `xml:"MovedEvent"`
	// The NewMailEvent element represents an event that is triggered by a new mail item in a mailbox.
	NewMailEvent *NewMailEvent `xml:"NewMailEvent"`
	// The PreviousWatermark element represents the watermark of the latest event that was successfully communicated to the client for the subscription.
	PreviousWatermark *PreviousWatermark `xml:"PreviousWatermark"`
	// The StatusEvent element represents a notification that no new activity has occurred in the mailbox.
	StatusEvent *StatusEvent `xml:"StatusEvent"`
	// The SubscriptionId element represents the identifier for a subscription.
	SubscriptionId *SubscriptionIdGetEvents `xml:"SubscriptionId"`
}

func (N *Notification) SetForMarshal() {
	N.XMLName.Local = "t:Notification"
}

func (N *Notification) GetSchema() *Schema {
	return &SchemaTypes
}
