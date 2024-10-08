package schema

// The AggregateOn element represents the property that is used to determine the order of grouped items for a grouped FindItem result set.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/aggregateon
import "encoding/xml"

type AggregateOn struct {
	XMLName xml.Name

	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
	// Indicates the maximum or minimum value of the property identified by the FieldURI element that is used for ordering the groups of items.The following are the possible values:  - Minimum  - Maximum
	Aggregate *string `xml:"Aggregate,attr"`
}

func (A *AggregateOn) SetForMarshal() {
	A.XMLName.Local = "t:AggregateOn"
}

func (A *AggregateOn) GetSchema() *Schema {
	return &SchemaTypes
}
