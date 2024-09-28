package goews

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/gunsluo/goews/v3/ntlmssp"
	"github.com/gunsluo/goews/v3/schema"
)

type Client interface {
	CreateItem(*schema.CreateItem) error
	DoRequest(Envelope, Element) error

	// FindItem(*FindItem) (*FindItemResponse, error)
	// GetItem(*GetItem) (*GetItemResponse, error)
	// QueryMessage(QueryMessageParams) ([]Message, error)

	SendEmail(SendEmailParams) error
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

// SendEmail helper method to send Message
func (c *client) SendEmail(param SendEmailParams) error {
	var to []*schema.ToRecipients
	for _, addr := range param.To {
		to = append(to, &schema.ToRecipients{
			Mailbox: &schema.Mailbox{
				EmailAddress: &schema.EmailAddressNonEmptyStringType{
					TEXT: addr,
				},
			},
		})
	}

	var cc []*schema.CcRecipients
	for _, addr := range param.Cc {
		cc = append(cc, &schema.CcRecipients{
			Mailbox: &schema.Mailbox{
				EmailAddress: &schema.EmailAddressNonEmptyStringType{
					TEXT: addr,
				},
			},
		})
	}

	var bcc []*schema.BccRecipients
	for _, addr := range param.Bcc {
		bcc = append(bcc, &schema.BccRecipients{
			Mailbox: &schema.Mailbox{
				EmailAddress: &schema.EmailAddressNonEmptyStringType{
					TEXT: addr,
				},
			},
		})
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
		CcRecipients:  cc[0],
		BccRecipients: bcc[0],
	}

	/*
		if len(param.FileAttachments) > 0 {
			m.Attachments = &TAttachments{}
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
	*/

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
