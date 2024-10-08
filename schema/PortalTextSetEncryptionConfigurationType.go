package schema

// The PortalText (SetEncryptionConfigurationType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/portaltext-setencryptionconfigurationtype
import "encoding/xml"

type PortalTextSetEncryptionConfigurationType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
