package goews

import "time"

type Envelope interface {
	GetEnvelopeBytes() ([]byte, error)
}

type SendEmailParams struct {
	From     string
	To       []string
	Cc       []string
	Bcc      []string
	Subject  string
	Body     string
	BodyType string

	FileAttachments []FileAttachment
}

type FileAttachment struct {
	AttachmentId *AttachmentId
	Name         string
	ContentType  string
	Size         int64
	Content      []byte
}

type AttachmentId struct {
	Id                string
	RootItemId        string
	RootItemChangeKey string
}

type QueryMessageParams struct {
	FolderId  string
	StartTime time.Time
	EndTime   time.Time
	Limit     int
	Offset    int
	BodyType  string
}
