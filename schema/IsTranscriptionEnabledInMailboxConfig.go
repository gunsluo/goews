package schema

// The IsTranscriptionEnabledInMailboxConfig element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/istranscriptionenabledinmailboxconfig
import "encoding/xml"

type IsTranscriptionEnabledInMailboxConfig struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
