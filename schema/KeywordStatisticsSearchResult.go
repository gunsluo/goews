package schema

// The KeywordStatisticsSearchResult element contains a single keyword search result.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/keywordstatisticssearchresult
import "encoding/xml"

type KeywordStatisticsSearchResult struct {
	XMLName xml.Name

	// The ItemHits element identifies how many times a keyword was found.
	ItemHits *ItemHits `xml:"ItemHits"`
	// The Keyword element specifies a single keyword.
	Keyword *Keyword `xml:"Keyword"`
	// The Size element specifies the total size of one or more mailbox items.
	Size *Sizelong `xml:"Size"`
}

func (K *KeywordStatisticsSearchResult) SetForMarshal() {
	K.XMLName.Local = "t:KeywordStatisticsSearchResult"
}

func (K *KeywordStatisticsSearchResult) GetSchema() *Schema {
	return &SchemaTypes
}
