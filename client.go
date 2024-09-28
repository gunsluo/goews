package goews

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
	"time"

	"github.com/gunsluo/goews/v3/ntlmssp"
)

type Client interface {
	// SendEmail(SendEmailParams) error
	// GetPersonaById(personaID string) (*Persona, error)
	// GetPersona(r *GetPersonaRequest) (*GetPersonaResponse, error)
	// GetUserPhoto(email string) (string, error)
	// GetDecodingUserPhoto(email string) ([]byte, error)
	// GetUserPhotoURL(email string) string
	// FindPeopleByCondition(q string) ([]Persona, error)
	// FindPeople(r *FindPeopleRequest) (*FindPeopleResponse, error)
	// GetRoomLists() (*GetRoomListsResponse, error)
	// ListUsersEvents(eventUsers []EventUser, from time.Time, duration time.Duration) (map[EventUser][]Event, error)
	// SendAndReceive(body []byte) ([]byte, error)
	// GetFolder(GetFolderParams) (*GetFolderResponse, error)
	// FindItem(*FindItem) (*FindItemResponse, error)
	// GetItem(*GetItem) (*GetItemResponse, error)
	// QueryMessage(QueryMessageParams) ([]Message, error)
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
	// bArr, err := c.SendAndReceive(e)
	// if err != nil {
	// 	return err
	// }
	// return operations.Unmarshal(bArr, oe)
	return nil
}
