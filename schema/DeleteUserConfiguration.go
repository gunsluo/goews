package schema

// The DeleteUserConfiguration element represents a request to delete a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteuserconfiguration
import "encoding/xml"

type DeleteUserConfiguration struct {
	XMLName xml.Name

	// The UserConfigurationName element represents the name of a user configuration object. The user configuration object name is the identifier for a user configuration object.
	UserConfigurationName *UserConfigurationName `xml:"UserConfigurationName"`
}

func (D *DeleteUserConfiguration) SetForMarshal() {
	D.XMLName.Local = "m:DeleteUserConfiguration"
}

func (D *DeleteUserConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}
