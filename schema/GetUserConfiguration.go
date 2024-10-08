package schema

// The GetUserConfiguration element represent a request to get a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserconfiguration
import "encoding/xml"

type GetUserConfiguration struct {
	XMLName xml.Name

	// The UserConfigurationName element represents the name of a user configuration object. The user configuration object name is the identifier for a user configuration object.
	UserConfigurationName *UserConfigurationName `xml:"UserConfigurationName"`
	// The UserConfigurationProperties element specifies the property types to get in a GetUserConfiguration operation.
	UserConfigurationProperties *UserConfigurationProperties `xml:"UserConfigurationProperties"`
}

func (G *GetUserConfiguration) SetForMarshal() {
	G.XMLName.Local = "m:GetUserConfiguration"
}

func (G *GetUserConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}
