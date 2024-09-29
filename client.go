package goews

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"time"

	"github.com/gunsluo/goews/v3/ntlmssp"
	"github.com/gunsluo/goews/v3/schema"
)

type Client interface {
	CreateItem(*schema.CreateItem) error
	FindItem(*schema.FindItem) (*schema.FindItemResponse, error)
	GetItem(*schema.GetItem) (*schema.GetItemResponse, error)
	DoRequest(Envelope, Element) error

	SendEmail(SendEmailParams) error
	QueryMessage(QueryMessageParams) ([]*schema.Message, error)
}

type Option func(*options)

func SetAddress(address string) Option {
	return func(o *options) {
		o.address = address
	}
}

func SetCredentials(username, password string) Option {
	return func(o *options) {
		o.username = username
		o.password = password
	}
}

func SetDomain(domain string) Option {
	return func(o *options) {
		o.domain = domain
	}
}

func EnabledNTLM() Option {
	return func(o *options) {
		o.ntlm = true
	}
}

func SkipTLS() Option {
	return func(o *options) {
		o.skipTLS = true
	}
}

func Debug() Option {
	return func(o *options) {
		o.dump = true
	}
}

type options struct {
	address  string
	username string
	password string
	ntlm     bool
	domain   string
	dump     bool

	skipTLS bool
	caPath  string
}

type client struct {
	httpClient *http.Client
	opts       options
}

func NewClient(opts ...Option) (Client, error) {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

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

	if o.caPath != "" {
		caCert, err := os.ReadFile(o.caPath)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		transport.TLSClientConfig = &tls.Config{
			RootCAs: caCertPool,
		}
	} else if o.skipTLS {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	if o.ntlm {
		httpClient.Transport = ntlmssp.NtlmTransport{
			Domain:       o.domain,
			User:         o.username,
			Password:     o.password,
			RoundTripper: transport,
		}
	} else {
		httpClient.Transport = transport
	}

	return &client{
		opts:       o,
		httpClient: httpClient,
	}, nil
}

func (c *client) DoRequest(e Envelope, oe Element) error {
	respBytes, err := c.SendAndReceive(e)
	if err != nil {
		return err
	}
	return Unmarshal(respBytes, oe)
}

func (c *client) SendAndReceive(e Envelope) ([]byte, error) {
	bb, err := e.GetEnvelopeBytes()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.opts.address, bytes.NewReader(bb))
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

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, err
}

func logRequest(c *client, req *http.Request) {
	if c.opts.dump {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Request:\n%v\n----\n", string(dump))
	}
}

func logResponse(c *client, resp *http.Response) {
	if c.opts.dump {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Response:\n%v\n----\n", string(dump))
	}
}

// CreateItem
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation-email-message
func (c *client) CreateItem(item *schema.CreateItem) error {
	envelope, err := NewEnvelopeMarshal(item)
	if err != nil {
		return err
	}

	resp := &schema.CreateItemResponse{}
	if err := c.DoRequest(envelope, resp); err != nil {
		return err
	}

	if resp.ResponseMessages != nil &&
		resp.ResponseMessages.CreateItemResponseMessage != nil &&
		resp.ResponseMessages.CreateItemResponseMessage.ResponseClass != nil &&
		*resp.ResponseMessages.CreateItemResponseMessage.ResponseClass == schema.CreateItemResponseMessageError {
		if resp.ResponseMessages.CreateItemResponseMessage.MessageText != nil {
			return errors.New(resp.ResponseMessages.CreateItemResponseMessage.MessageText.TEXT)
		}

		return errors.New("Unknow Error")
	}

	return nil
}

// FindItem
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/finditem-operation1
func (c *client) FindItem(item *schema.FindItem) (*schema.FindItemResponse, error) {
	envelope, err := NewEnvelopeMarshal(item)
	if err != nil {
		return nil, err
	}

	resp := &schema.FindItemResponse{}
	if err := c.DoRequest(envelope, resp); err != nil {
		return nil, err
	}

	if resp.ResponseMessages != nil &&
		resp.ResponseMessages.FindItemResponseMessage != nil &&
		resp.ResponseMessages.FindItemResponseMessage.ResponseClass != nil &&
		*resp.ResponseMessages.FindItemResponseMessage.ResponseClass == schema.FindItemResponseMessageError {
		if resp.ResponseMessages.FindItemResponseMessage.MessageText != nil {
			return nil, errors.New(resp.ResponseMessages.FindItemResponseMessage.MessageText.TEXT)
		}

		return nil, errors.New("Unknow Error")
	}

	return resp, nil
}

// GetItem
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getitem-operation-email-message
func (c *client) GetItem(item *schema.GetItem) (*schema.GetItemResponse, error) {
	envelope, err := NewEnvelopeMarshal(item)
	if err != nil {
		return nil, err
	}

	resp := &schema.GetItemResponse{}
	if err := c.DoRequest(envelope, resp); err != nil {
		return nil, err
	}

	if resp.ResponseMessages != nil &&
		resp.ResponseMessages.GetItemResponseMessage != nil &&
		resp.ResponseMessages.GetItemResponseMessage.ResponseClass != nil &&
		*resp.ResponseMessages.GetItemResponseMessage.ResponseClass == schema.FindItemResponseMessageError {
		if resp.ResponseMessages.GetItemResponseMessage.MessageText != nil {
			return nil, errors.New(resp.ResponseMessages.GetItemResponseMessage.MessageText.TEXT)
		}

		return nil, errors.New("Unknow Error")
	}

	return resp, nil
}

// SendEmail helper method to send Message
func (c *client) SendEmail(param SendEmailParams) error {
	var to *schema.ToRecipients
	if len(param.To) > 0 {
		to = &schema.ToRecipients{}
		for _, addr := range param.To {
			to.Mailbox = append(to.Mailbox, &schema.Mailbox{
				EmailAddress: &schema.EmailAddressNonEmptyStringType{
					TEXT: addr,
				},
			})
		}
	}

	var cc *schema.CcRecipients
	if len(param.Cc) > 0 {
		cc = &schema.CcRecipients{}
		for _, addr := range param.Cc {
			cc.Mailbox = append(cc.Mailbox, &schema.Mailbox{
				EmailAddress: &schema.EmailAddressNonEmptyStringType{
					TEXT: addr,
				},
			})
		}
	}

	var bcc *schema.BccRecipients
	if len(param.Bcc) > 0 {
		bcc = &schema.BccRecipients{}
		for _, addr := range param.Bcc {
			bcc.Mailbox = append(bcc.Mailbox, &schema.Mailbox{
				EmailAddress: &schema.EmailAddressNonEmptyStringType{
					TEXT: addr,
				},
			})
		}
	}

	var attachments *schema.Attachments
	if len(param.FileAttachments) > 0 {
		attachments = &schema.Attachments{}
		for _, fa := range param.FileAttachments {
			var attachmentId *schema.AttachmentId
			if fa.AttachmentId != nil {
				attachmentId = &schema.AttachmentId{
					Id:                getPTR[string](fa.AttachmentId.Id),
					RootItemChangeKey: getPTR[string](fa.AttachmentId.RootItemChangeKey),
					RootItemId:        getPTR[string](fa.AttachmentId.RootItemId),
				}
			}

			var contentType *schema.ContentType
			if fa.ContentType != "" {
				contentType = &schema.ContentType{TEXT: fa.ContentType}
			}

			content := base64.StdEncoding.EncodeToString(fa.Content)
			attachments.FileAttachment = append(
				attachments.FileAttachment,
				&schema.FileAttachment{
					AttachmentId:   attachmentId,
					Name:           &schema.NameAttachmentType{TEXT: fa.Name},
					ContentType:    contentType,
					Size:           &schema.Size{TEXT: fa.Size},
					IsInline:       &schema.IsInline{TEXT: false},
					IsContactPhoto: &schema.IsContactPhoto{TEXT: false},
					Content:        &schema.Content{TEXT: content},
				})
		}
	}

	message := &schema.Message{
		Subject: &schema.Subject{TEXT: param.Subject},
		Body: &schema.Body{
			BodyType: getPTR[string](param.BodyType),
			TEXT:     param.Body,
		},
		Sender: &schema.Sender{
			Mailbox: &schema.Mailbox{
				EmailAddress: &schema.EmailAddressNonEmptyStringType{
					TEXT: param.From,
				},
			},
		},
		ToRecipients:  to,
		CcRecipients:  cc,
		BccRecipients: bcc,
		Attachments:   attachments,
	}

	item := &schema.CreateItem{
		MessageDisposition: getPTR[string]("SendAndSaveCopy"),
		SavedItemFolderId: &schema.SavedItemFolderId{
			DistinguishedFolderId: &schema.DistinguishedFolderId{
				Id: getPTR[string]("sentitems"),
			},
		},
		Items: &schema.ItemsNonEmptyArrayOfAllItemsType{
			Message: message,
		},
	}

	return c.CreateItem(item)
}

func (c *client) QueryMessage(param QueryMessageParams) ([]*schema.Message, error) {
	var restriction *schema.Restriction
	if !param.StartTime.IsZero() {
		var endTime time.Time
		if param.EndTime.IsZero() {
			endTime = time.Now()
		} else {
			endTime = param.EndTime
		}

		restriction = &schema.Restriction{And: &schema.And{
			IsGreaterThanOrEqualTo: &schema.IsGreaterThanOrEqualTo{
				FieldURI: &schema.FieldURI{
					FieldURI: getPTR[string]("item:DateTimeReceived"),
				},
				FieldURIOrConstant: &schema.FieldURIOrConstant{
					Constant: &schema.Constant{
						Value: getPTR[string](param.StartTime.UTC().Format(time.RFC3339)),
					},
				},
			},
			IsLessThanOrEqualTo: &schema.IsLessThanOrEqualTo{
				FieldURI: &schema.FieldURI{
					FieldURI: getPTR[string]("item:DateTimeReceived"),
				},
				FieldURIOrConstant: &schema.FieldURIOrConstant{
					Constant: &schema.Constant{
						Value: getPTR[string](endTime.UTC().Format(time.RFC3339)),
					},
				},
			},
		}}
	}

	item := &schema.FindItem{
		Traversal: getPTR[string](schema.FindItemShallow),
		ItemShape: &schema.ItemShape{
			BaseShape: &schema.BaseShape{TEXT: schema.BaseShapeIdOnly},
			AdditionalProperties: &schema.AdditionalProperties{
				FieldURI: []*schema.FieldURI{
					{FieldURI: getPTR[string]("item:Subject")},
					{FieldURI: getPTR[string]("item:DateTimeReceived")},
					{FieldURI: getPTR[string]("message:Sender")},
				},
			},
		},
		IndexedPageItemView: &schema.IndexedPageItemView{
			MaxEntriesReturned: getPTR[string](strconv.Itoa(param.Limit)),
			Offset:             getPTR[string](strconv.Itoa(param.Offset)),
			BasePoint:          getPTR[string](schema.BasePointBeginning),
		},
		SortOrder: &schema.SortOrder{
			FieldOrder: &schema.FieldOrder{
				Order:    getPTR[string](schema.OrderAscending),
				FieldURI: &schema.FieldURI{FieldURI: getPTR[string]("item:DateTimeReceived")},
			},
		},
		ParentFolderIds: &schema.ParentFolderIds{
			DistinguishedFolderId: &schema.DistinguishedFolderId{
				Id: getPTR[string](param.FolderId),
			},
		},
		Restriction: restriction,
	}

	resp, err := c.FindItem(item)
	if err != nil {
		return nil, err
	}

	var messages []*schema.Message
	if resp.ResponseMessages != nil &&
		resp.ResponseMessages.FindItemResponseMessage != nil &&
		resp.ResponseMessages.FindItemResponseMessage.RootFolder != nil &&
		resp.ResponseMessages.FindItemResponseMessage.RootFolder.Items != nil {
		var bodyType *schema.BodyType
		if param.BodyType != "" {
			bodyType = &schema.BodyType{TEXT: param.BodyType}
		}

		for _, message := range resp.ResponseMessages.FindItemResponseMessage.RootFolder.Items.Message {
			if message.ItemId == nil {
				return nil, errors.New("missing item id")
			}

			itemmore, err := c.GetItem(&schema.GetItem{
				ItemShape: &schema.ItemShape{
					BaseShape: &schema.BaseShape{TEXT: schema.BaseShapeIdOnly},
					AdditionalProperties: &schema.AdditionalProperties{
						FieldURI: []*schema.FieldURI{
							{FieldURI: getPTR[string]("item:Body")},
						},
					},
					BodyType: bodyType,
				},
				ItemIds: &schema.ItemIds{
					ItemId: &schema.ItemId{
						Id:        message.ItemId.Id,
						ChangeKey: message.ItemId.ChangeKey,
					},
				},
			})
			if err != nil {
				return nil, err
			}

			if itemmore.ResponseMessages != nil &&
				itemmore.ResponseMessages.GetItemResponseMessage != nil &&
				itemmore.ResponseMessages.GetItemResponseMessage.Items != nil {
				for _, itemMessage := range itemmore.ResponseMessages.GetItemResponseMessage.Items.Message {
					if itemMessage.Body != nil {
						message.Body = itemMessage.Body
						break
					}
				}
			}

			messages = append(messages, message)
		}
	}

	return messages, nil
}
