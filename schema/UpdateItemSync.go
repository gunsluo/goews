package schema

// The Update element identifies a single item to update in the local client store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/update-itemsync
import "encoding/xml"

type UpdateItemSync struct {
	XMLName xml.Name

	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem *CalendarItem `xml:"CalendarItem"`
	// The Contact element represents a contact item in the Exchange store.
	Contact *Contact `xml:"Contact"`
	// The DistributionList element represents a distribution list.
	DistributionList *DistributionList `xml:"DistributionList"`
	// The Item element represents a generic item in the Exchange store.
	Item *Item `xml:"Item"`
	// The MeetingCancellation element represents a meeting cancellation in the Exchange store.
	MeetingCancellation *MeetingCancellation `xml:"MeetingCancellation"`
	// The MeetingMessage element represents a meeting in the Exchange store.
	MeetingMessage *MeetingMessage `xml:"MeetingMessage"`
	// The MeetingRequest element represents a meeting request in the Exchange store.
	MeetingRequest *MeetingRequest `xml:"MeetingRequest"`
	// The MeetingResponse element represents a meeting response in the Exchange store.
	MeetingResponse *MeetingResponse `xml:"MeetingResponse"`
	// The Message element represents a Microsoft Exchange e-mail message.
	Message *Message `xml:"Message"`
	// The Task element represents a task in the Exchange store.
	Task *Task `xml:"Task"`
}

func (U *UpdateItemSync) SetForMarshal() {
	U.XMLName.Local = "t:Update"
}

func (U *UpdateItemSync) GetSchema() *Schema {
	return &SchemaTypes
}
