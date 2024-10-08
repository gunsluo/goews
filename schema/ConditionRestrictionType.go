package schema

// The Condition element specifies the condition that is used to identify the end of a search for a FindItem or a FindConversation operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/condition-restrictiontype
import "encoding/xml"

type ConditionRestrictionType struct {
	XMLName xml.Name

	// The SearchExpression element is an abstract element that represents the substituted element within a restriction. All search expressions derive from this base type. This element is not used in an XML instance document.
	SearchExpression *SearchExpression `xml:"SearchExpression"`
}

func (C *ConditionRestrictionType) SetForMarshal() {
	C.XMLName.Local = "t:Condition"
}

func (C *ConditionRestrictionType) GetSchema() *Schema {
	return &SchemaTypes
}
