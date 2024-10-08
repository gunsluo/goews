package schema

// The MailboxStatisticsSearchResult element contains the results of a keyword search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mailboxstatisticssearchresult
import "encoding/xml"

type MailboxStatisticsSearchResult struct {
	XMLName xml.Name

	// The KeywordStatisticsSearchResult element contains a single keyword search result.
	KeywordStatisticsSearchResult *KeywordStatisticsSearchResult `xml:"KeywordStatisticsSearchResult"`
	// The UserMailbox element identifies a user mailbox.
	UserMailbox *UserMailbox `xml:"UserMailbox"`
}
