package schema

// The ItemClasses element contains a list of item classes that represents all the item classes of the conversation items in the current folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemclasses-arrayofitemclasstype
import "encoding/xml"

type ItemClassesArrayOfItemClassType struct {
	XMLName xml.Name

	// The ItemClass element represents the message class of an item.
	ItemClass *ItemClass `xml:"ItemClass"`
}

func (I *ItemClassesArrayOfItemClassType) SetForMarshal() {
	I.XMLName.Local = "t:ItemClasses"
}

func (I *ItemClassesArrayOfItemClassType) GetSchema() *Schema {
	return &SchemaTypes
}
