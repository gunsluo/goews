package goews

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gunsluo/goews/v2/ntlmssp"
)

const (
	soapStart = `<?xml version="1.0" encoding="utf-8" ?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
		xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages"
		xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types"
		xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  		<soap:Header>
    		<t:RequestServerVersion Version="Exchange2013_SP1" />
  		</soap:Header>
  		<soap:Body>
`
	soapEnd = `
</soap:Body></soap:Envelope>`
)

type Client interface {
	SendEmail(SendEmailParams) error
	CreateHTMLEvent(to, optional []string, subject, body, location string, from time.Time, duration time.Duration) error
	CreateEvent(to, optional []string, subject, body, location string, from time.Time, duration time.Duration) error
	GetPersonaById(personaID string) (*Persona, error)
	GetPersona(r *GetPersonaRequest) (*GetPersonaResponse, error)
	GetUserPhoto(email string) (string, error)
	GetDecodingUserPhoto(email string) ([]byte, error)
	GetUserPhotoURL(email string) string
	FindPeopleByCondition(q string) ([]Persona, error)
	FindPeople(r *FindPeopleRequest) (*FindPeopleResponse, error)
	GetRoomLists() (*GetRoomListsResponse, error)
	ListUsersEvents(eventUsers []EventUser, from time.Time, duration time.Duration) (map[EventUser][]Event, error)
	SendAndReceive(body []byte) ([]byte, error)
}

type Config struct {
	Address  string
	Username string
	Password string
	NTLM     bool
	Domain   string
	Dump     bool

	SkipTLS bool
	CaPath  string
}

type client struct {
	httpClient *http.Client
	config     Config
}

func NewClient(config Config) (Client, error) {
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	transport := &http.Transport{
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if config.CaPath != "" {
		caCert, err := ioutil.ReadFile(config.CaPath)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		transport.TLSClientConfig = &tls.Config{
			RootCAs: caCertPool,
		}
	} else if config.SkipTLS {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	if config.NTLM {
		httpClient.Transport = ntlmssp.NtlmTransport{
			Domain:       config.Domain,
			User:         config.Username,
			Password:     config.Password,
			RoundTripper: transport,
		}
	} else {
		httpClient.Transport = transport
	}

	return &client{
		config:     config,
		httpClient: httpClient,
	}, nil
}

func (c *client) SendAndReceive(body []byte) ([]byte, error) {
	bb := []byte(soapStart)
	bb = append(bb, body...)
	bb = append(bb, soapEnd...)

	req, err := http.NewRequest("POST", c.config.Address, bytes.NewReader(bb))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	logRequest(c, req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	logResponse(c, resp)

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(resp)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, err
}

func logRequest(c *client, req *http.Request) {
	if c.config.Dump {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Request:\n%v\n----\n", string(dump))
	}
}

func logResponse(c *client, resp *http.Response) {
	if c.config.Dump {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Response:\n%v\n----\n", string(dump))
	}
}

// SendEmail helper method to send Message
func (c *client) SendEmail(param SendEmailParams) error {
	m := Message{
		//ItemClass: "IPM.Note",
		Subject: param.Subject,
		Body: Body{
			BodyType: param.BodyType,
			Body:     []byte(param.Body),
		},
		Sender: OneMailbox{
			Mailbox: Mailbox{
				EmailAddress: param.From,
			},
		},
	}
	mb := make([]Mailbox, len(param.To))
	for i, addr := range param.To {
		mb[i].EmailAddress = addr
	}
	m.ToRecipients.Mailbox = append(m.ToRecipients.Mailbox, mb...)

	if len(param.Cc) > 0 {
		m.CcRecipients = &XMailbox{}
		for _, addr := range param.Cc {
			m.CcRecipients.Mailbox = append(m.CcRecipients.Mailbox,
				Mailbox{
					EmailAddress: addr,
				})
		}
	}

	if len(param.Bcc) > 0 {
		m.BccRecipients = &XMailbox{}
		for _, addr := range param.Bcc {
			m.BccRecipients.Mailbox = append(m.BccRecipients.Mailbox,
				Mailbox{
					EmailAddress: addr,
				})
		}
	}

	if len(param.FileAttachments) > 0 {
		m.Attachments = &Attachments{}
		for _, attachment := range param.FileAttachments {
			content := base64.StdEncoding.EncodeToString(attachment.Content)
			m.Attachments.Files = append(m.Attachments.Files,
				FileAttachment{
					AttachmentId:     attachment.AttachmentId,
					Name:             attachment.Name,
					ContentType:      attachment.ContentType,
					ContentId:        "",
					ContentLocation:  "",
					Size:             attachment.Size,
					LastModifiedTime: "",
					IsInline:         false,
					IsContactPhoto:   false,
					Content:          content,
				})
		}
	}

	return c.CreateMessageItem(m)
}

// CreateMessageItem
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation-email-message
func (c *client) CreateMessageItem(m ...Message) error {

	item := &CreateItem{
		MessageDisposition: "SendAndSaveCopy",
		SavedItemFolderId:  SavedItemFolderId{DistinguishedFolderId{Id: "sentitems"}},
	}
	item.Items.Message = append(item.Items.Message, m...)

	xmlBytes, err := xml.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return err
	}

	if err := checkCreateItemResponseForErrors(bb); err != nil {
		return err
	}

	return nil
}

func checkCreateItemResponseForErrors(bb []byte) error {
	var soapResp createItemResponseBodyEnvelop
	if err := xml.Unmarshal(bb, &soapResp); err != nil {
		return err
	}

	resp := soapResp.Body.CreateItemResponse.ResponseMessages.CreateItemResponseMessage
	if resp.ResponseClass == ResponseClassError {
		return errors.New(resp.MessageText)
	}
	return nil
}

// CreateCalendarItem
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation-calendar-item
func (c *client) CreateCalendarItem(ci ...CalendarItem) error {

	item := &CreateItem{
		SendMeetingInvitations: "SendToAllAndSaveCopy",
		SavedItemFolderId:      SavedItemFolderId{DistinguishedFolderId{Id: "calendar"}},
	}
	item.Items.CalendarItem = append(item.Items.CalendarItem, ci...)

	xmlBytes, err := xml.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return err
	}

	if err := checkCreateItemResponseForErrors(bb); err != nil {
		return err
	}

	return nil
}

func (c *client) CreateHTMLEvent(
	to, optional []string, subject, body, location string, from time.Time, duration time.Duration,
) error {
	return c.createEvent(to, optional, subject, body, location, "HTML", from, duration)
}

// CreateEvent helper method to send Message
func (c *client) CreateEvent(
	to, optional []string, subject, body, location string, from time.Time, duration time.Duration,
) error {
	return c.createEvent(to, optional, subject, body, location, "Text", from, duration)
}

func (c *client) createEvent(
	to, optional []string, subject, body, location, bodyType string, from time.Time, duration time.Duration,
) error {

	requiredAttendees := make([]Attendee, len(to))
	for i, tt := range to {
		requiredAttendees[i] = Attendee{Mailbox: Mailbox{EmailAddress: tt}}
	}

	optionalAttendees := make([]Attendee, len(optional))
	for i, tt := range optional {
		optionalAttendees[i] = Attendee{Mailbox: Mailbox{EmailAddress: tt}}
	}

	room := make([]Attendee, 1)
	room[0] = Attendee{Mailbox: Mailbox{EmailAddress: location}}

	m := CalendarItem{
		Subject: subject,
		Body: Body{
			BodyType: bodyType,
			Body:     []byte(body),
		},
		ReminderIsSet:              true,
		ReminderMinutesBeforeStart: 15,
		Start:                      from,
		End:                        from.Add(duration),
		IsAllDayEvent:              false,
		LegacyFreeBusyStatus:       BusyTypeBusy,
		Location:                   location,
		RequiredAttendees:          []Attendees{{Attendee: requiredAttendees}},
		OptionalAttendees:          []Attendees{{Attendee: optionalAttendees}},
		Resources:                  []Attendees{{Attendee: room}},
	}

	return c.CreateCalendarItem(m)
}

// GetPersonaById find persona slice by query string
func (c *client) GetPersonaById(personaID string) (*Persona, error) {

	resp, err := c.GetPersona(&GetPersonaRequest{
		PersonaId: PersonaId{Id: personaID},
	})

	if err != nil {
		return nil, err
	}

	return &resp.Persona, nil
}

// GetPersona
//https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getpersona-operation
func (c *client) GetPersona(r *GetPersonaRequest) (*GetPersonaResponse, error) {

	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp getPersonaResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	if soapResp.Body.FindPeopleResponse.ResponseClass == ResponseClassError {
		return nil, errors.New(soapResp.Body.FindPeopleResponse.MessageText)
	}

	return &soapResp.Body.FindPeopleResponse, nil
}

// GetUserPhoto
//https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuserphoto-operation
func (c *client) GetUserPhoto(email string) (string, error) {
	resp, err := c.getUserPhoto(&GetUserPhotoRequest{
		Email:         email,
		SizeRequested: "HR48x48",
	})

	if err != nil {
		return "", err
	}

	return resp.PictureData, nil
}

func (c *client) GetDecodingUserPhoto(email string) ([]byte, error) {
	s, err := c.GetUserPhoto(email)
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(s)
}

func (c *client) GetUserPhotoURL(email string) string {
	return fmt.Sprintf("%s/s/GetUserPhoto?email=%s&size=HR48x48", c.config.Address, email)
}

func (c *client) getUserPhoto(r *GetUserPhotoRequest) (*GetUserPhotoResponse, error) {

	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp getUserPhotoResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	if soapResp.Body.GetUserPhotoResponse.ResponseClass == ResponseClassError {
		return nil, errors.New(soapResp.Body.GetUserPhotoResponse.MessageText)
	}

	return &soapResp.Body.GetUserPhotoResponse, nil
}

// FindPeopleByCondition find persona slice by query string
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/findpeople-operation
func (c *client) FindPeopleByCondition(q string) ([]Persona, error) {
	req := &FindPeopleRequest{IndexedPageItemView: IndexedPageItemView{
		MaxEntriesReturned: math.MaxInt32,
		Offset:             0,
		BasePoint:          BasePointBeginning,
	}, ParentFolderId: ParentFolderId{
		DistinguishedFolderId: DistinguishedFolderId{Id: "directory"}},
		PersonaShape: &PersonaShape{BaseShape: BaseShapeIdOnly,
			AdditionalProperties: AdditionalProperties{
				FieldURI: []FieldURI{
					{FieldURI: "persona:DisplayName"},
					{FieldURI: "persona:Title"},
					{FieldURI: "persona:EmailAddress"},
					{FieldURI: "persona:Departments"},
				},
			}},
		QueryString: q,
	}

	resp, err := c.FindPeople(req)
	if err != nil {
		return nil, err
	}

	return resp.People.Persona, nil
}

// FindPeople
func (c *client) FindPeople(r *FindPeopleRequest) (*FindPeopleResponse, error) {

	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp findPeopleResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	if soapResp.Body.FindPeopleResponse.ResponseClass == ResponseClassError {
		return nil, errors.New(soapResp.Body.FindPeopleResponse.MessageText)
	}

	return &soapResp.Body.FindPeopleResponse, nil
}

func (c *client) ListUsersEvents(
	eventUsers []EventUser, from time.Time, duration time.Duration,
) (map[EventUser][]Event, error) {
	req := buildGetUserAvailabilityRequest(eventUsers, from, duration)

	resp, err := c.getUserAvailability(req)
	if err != nil {
		return nil, err
	}

	events, err := traverseGetUserAvailabilityResponse(eventUsers, resp)
	if err != nil {
		return nil, err
	}

	return events, nil
}

//https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseravailability-operation
func (c *client) getUserAvailability(r *GetUserAvailabilityRequest) (*GetUserAvailabilityResponse, error) {
	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp getUserAvailabilityResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	resp := soapResp.Body.GetUserAvailabilityResponse

	err = checkForFunctionalError(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func checkForFunctionalError(resp *GetUserAvailabilityResponse) error {

	if len(resp.FreeBusyResponseArray.FreeBusyResponse) > 0 {
		for _, rr := range resp.FreeBusyResponseArray.FreeBusyResponse {
			if rr.ResponseMessage.ResponseClass == ResponseClassError {
				return errors.New(rr.ResponseMessage.MessageText)
			}
		}
	}
	return nil
}

func buildGetUserAvailabilityRequest(
	eventUsers []EventUser, from time.Time, duration time.Duration,
) *GetUserAvailabilityRequest {

	mb := make([]MailboxData, 0)
	for _, mm := range eventUsers {
		mb = append(mb, MailboxData{
			Email: Email{
				Name:        "",
				Address:     mm.Email,
				RoutingType: "SMTP",
			},
			AttendeeType:     mm.AttendeeType,
			ExcludeConflicts: false,
		})
	}
	_, offset := time.Now().Zone()
	req := &GetUserAvailabilityRequest{
		//https://github.com/MicrosoftDocs/office-developer-exchange-docs/issues/61
		TimeZone: TimeZone{
			Bias: -offset / 60,
			StandardTime: TimeZoneTime{ // I don't have much clue about the values here
				Bias:      0,
				Time:      "02:00:00",
				DayOrder:  5,
				Month:     10,
				DayOfWeek: "Sunday",
			},
			DaylightTime: TimeZoneTime{ // I don't have much clue about the values here
				Bias:      0,
				Time:      "02:00:00",
				DayOrder:  1,
				Month:     4,
				DayOfWeek: "Sunday",
			},
		},
		MailboxDataArray: MailboxDataArray{MailboxData: mb},
		FreeBusyViewOptions: FreeBusyViewOptions{
			TimeWindow: TimeWindow{
				StartTime: from,
				EndTime:   from.Add(duration),
			},
			RequestedView: RequestedViewFreeBusy,
		},
	}
	return req
}

func traverseGetUserAvailabilityResponse(
	eventUsers []EventUser, resp *GetUserAvailabilityResponse,
) (map[EventUser][]Event, error) {

	m := make(map[EventUser][]Event)
	for i, rr := range resp.FreeBusyResponseArray.FreeBusyResponse {

		ce := make([]Event, 0)
		for _, cc := range rr.FreeBusyView.CalendarEventArray.CalendarEvent {

			start, err := cc.StartTime.ToTime()
			if err != nil {
				return nil, err
			}

			end, err := cc.EndTime.ToTime()
			if err != nil {
				return nil, err
			}

			ce = append(ce, Event{
				Start:    start,
				End:      end,
				BusyType: cc.BusyType,
			})
		}
		m[eventUsers[i]] = ce
	}
	return m, nil
}

func (c *client) GetRoomLists() (*GetRoomListsResponse, error) {
	xmlBytes, err := xml.MarshalIndent(&GetRoomListsRequest{}, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp getRoomListsResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	return &soapResp.Body.GetRoomListsResponse, nil
}
