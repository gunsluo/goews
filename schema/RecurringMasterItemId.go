package schema

// The RecurringMasterItemId element identifies a recurrence master item by identifying the identifiers of one of its related occurrence items.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurringmasteritemid
import "encoding/xml"

type RecurringMasterItemId struct {
	XMLName xml.Name

	// Identifies a specific version of a single occurrence of a recurring master item. Additionally, the recurring master item is also identified because it and the single occurrence will contain the same change key. This attribute is optional.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Identifies a single occurrence of a recurring master item. This attribute is required.
	OccurrenceId *string `xml:"OccurrenceId,attr"`
}

func (R *RecurringMasterItemId) SetForMarshal() {
	R.XMLName.Local = "t:RecurringMasterItemId"
}

func (R *RecurringMasterItemId) GetSchema() *Schema {
	return &SchemaTypes
}
