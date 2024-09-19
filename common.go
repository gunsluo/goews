package goews

import (
	"fmt"
	"time"
)

type ResponseClass string

const (
	ResponseClassSuccess ResponseClass = "Success"
	ResponseClassWarning ResponseClass = "Warning"
	ResponseClassError   ResponseClass = "Error"
)

type Response struct {
	ResponseClass ResponseClass `xml:"ResponseClass,attr"`
	MessageText   string        `xml:"MessageText"`
	ResponseCode  string        `xml:"ResponseCode"`
	MessageXml    MessageXml    `xml:"MessageXml"`
}

type EmailAddress struct {
	Name         string `xml:"Name"`
	EmailAddress string `xml:"EmailAddress"`
	RoutingType  string `xml:"RoutingType"`
	MailboxType  string `xml:"MailboxType"`
	ItemId       ItemId `xml:"ItemId"`
}

type MessageXml struct {
	ExceptionType       string `xml:"ExceptionType"`
	ExceptionCode       string `xml:"ExceptionCode"`
	ExceptionServerName string `xml:"ExceptionServerName"`
	ExceptionMessage    string `xml:"ExceptionMessage"`
}

type DistinguishedFolderId struct {
	// List of values:
	// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedfolderid
	Id string `xml:"Id,attr"`
}

const (
	// Represents the Calendar folder.
	DistinguishedFolderIdcalendar = `calendar`
	// Represents the Contacts folder.
	DistinguishedFolderIdcontacts = `contacts`
	// Represents the Deleted Items folder.
	DistinguishedFolderIddeleteditems = `deleteditems`
	// Represents the Drafts folder.
	DistinguishedFolderIddrafts = `drafts`
	// Represents the Inbox folder.
	DistinguishedFolderIdinbox = `inbox`
	// Represents the Journal folder.
	DistinguishedFolderIdjournal = `journal`
	// Represents the Notes folder.
	DistinguishedFolderIdnotes = `notes`
	// Represents the Outbox folder.
	DistinguishedFolderIdoutbox = `outbox`
	// Represents the Sent Items folder.
	DistinguishedFolderIdsentitems = `sentitems`
	// Represents the Tasks folder.
	DistinguishedFolderIdtasks = `tasks`
	// Represents the message folder root.
	DistinguishedFolderIdmsgfolderroot = `msgfolderroot`
	// Represents the root of the mailbox.
	DistinguishedFolderIdroot = `root`
	// Represents the Junk Email folder.
	DistinguishedFolderIdjunkemail = `junkemail`
	// Represents the Search Folders folder.
	DistinguishedFolderIdsearchfolders = `searchfolders`
	// Represents the Voice Mail folder.
	DistinguishedFolderIdvoicemail = `voicemail`
	// Represents the dumpster root folder.
	DistinguishedFolderIdrecoverableitemsroot = `recoverableitemsroot`
	// Represents the dumpster deletions folder.
	DistinguishedFolderIdrecoverableitemsdeletions = `recoverableitemsdeletions`
	// Represents the dumpster versions folder.
	DistinguishedFolderIdrecoverableitemsversions = `recoverableitemsversions`
	// Represents the dumpster purges folder.
	DistinguishedFolderIdrecoverableitemspurges = `recoverableitemspurges`
	// Represents the root archive folder.
	DistinguishedFolderIdarchiveroot = `archiveroot`
	// Represents the root archive message folder.
	DistinguishedFolderIdarchivemsgfolderroot = `archivemsgfolderroot`
	// Represents the archive deleted items folder.
	DistinguishedFolderIdarchivedeleteditems = `archivedeleteditems`
	// Represents the archive Inbox folder. Versions of Exchange starting with build number 15.00.0913.09 include this value.
	DistinguishedFolderIdarchiveinbox = `archiveinbox`
	// Represents the archive recoverable items root folder.
	DistinguishedFolderIdarchiverecoverableitemsroot = `archiverecoverableitemsroot`
	// Represents the archive recoverable items deletions folder.
	DistinguishedFolderIdarchiverecoverableitemsdeletions = `archiverecoverableitemsdeletions`
	// Represents the archive recoverable items versions folder.
	DistinguishedFolderIdarchiverecoverableitemsversions = `archiverecoverableitemsversions`
	// Represents the archive recoverable items purges folder.
	DistinguishedFolderIdarchiverecoverableitemspurges = `archiverecoverableitemspurges`
	// Represents the sync issues folder.
	DistinguishedFolderIdsyncissues = `syncissues`
	// Represents the conflicts folder.
	DistinguishedFolderIdconflicts = `conflicts`
	// Represents the local failures folder.
	DistinguishedFolderIdlocalfailures = `localfailures`
	// Represents the server failures folder.
	DistinguishedFolderIdserverfailures = `serverfailures`
	// Represents the recipient cache folder.
	DistinguishedFolderIdrecipientcache = `recipientcache`
	// Represents the quick contacts folder.
	DistinguishedFolderIdquickcontacts = `quickcontacts`
	// Represents the conversation history folder.
	DistinguishedFolderIdconversationhistory = `conversationhistory`
	// Represents the admin audit logs folder.
	DistinguishedFolderIdadminauditlogs = `adminauditlogs`
	// Represents the todo search folder.
	DistinguishedFolderIdtodosearch = `todosearch`
	// Represents the My Contacts folder.
	DistinguishedFolderIdmycontacts = `mycontacts`
	// Represents the directory folder.
	DistinguishedFolderIddirectory = `directory`
	// Represents the IM contact list folder.
	DistinguishedFolderIdimcontactlist = `imcontactlist`
	// Represents the people connect folder.
	DistinguishedFolderIdpeopleconnect = `peopleconnect`
	// Represents the Favorites folder.
	DistinguishedFolderIdfavorites = `favorites`
)

type Persona struct {
	PersonaId            PersonaId            `xml:"PersonaId"`
	DisplayName          string               `xml:"DisplayName"`
	Title                string               `xml:"Title"`
	Department           string               `xml:"Department"`
	Departments          Departments          `xml:"Departments"`
	EmailAddress         EmailAddress         `xml:"EmailAddress"`
	RelevanceScore       int                  `xml:"RelevanceScore"`
	BusinessPhoneNumbers BusinessPhoneNumbers `xml:"BusinessPhoneNumbers"`
	MobilePhones         MobilePhones         `xml:"MobilePhones"`
	OfficeLocations      OfficeLocations      `xml:"OfficeLocations"`
}

type PersonaId struct {
	Id string `xml:"Id,attr"`
}

type BusinessPhoneNumbers struct {
	PhoneNumberAttributedValue PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

type MobilePhones struct {
	PhoneNumberAttributedValue PhoneNumberAttributedValue `xml:"PhoneNumberAttributedValue"`
}

type Value struct {
	Number string `json:"Number"`
	Type   string `json:"Type"`
}

type PhoneNumberAttributedValue struct {
	Value Value `json:"Value"`
}

type OfficeLocations struct {
	StringAttributedValue StringAttributedValue `xml:"StringAttributedValue"`
}

type Departments struct {
	StringAttributedValue StringAttributedValue `xml:"StringAttributedValue"`
}

type StringAttributedValue struct {
	Value string `json:"Value"`
}

type Time string

func (t Time) ToTime() (time.Time, error) {
	offset, err := getRFC3339Offset(time.Now())
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(time.RFC3339, string(t)+offset)

}

// return RFC3339 formatted offset, ex: +03:00 -03:30
func getRFC3339Offset(t time.Time) (string, error) {

	_, offset := t.Zone()
	i := int(float32(offset) / 36)

	sign := "+"
	if i < 0 {
		i = -i
		sign = "-"
	}
	hour := i / 100
	min := i % 100
	min = (60 * min) / 100

	return fmt.Sprintf("%s%02d:%02d", sign, hour, min), nil
}
