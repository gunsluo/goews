package schema

// The ProxySecurityContext element is used by HTTP proxy of the computer that is running Microsoft Exchange Server 2007 that has the Client Access server role installed and is not used by Exchange Web Services operations. This element was introduced in Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proxysecuritycontext
import "encoding/xml"

type ProxySecurityContext struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (P *ProxySecurityContext) SetForMarshal() {
	P.XMLName.Local = "t:ProxySecurityContext"
}

func (P *ProxySecurityContext) GetSchema() *Schema {
	return &SchemaTypes
}
