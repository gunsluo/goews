package goews

import "time"

type SendEmailParams struct {
	From     string
	To       []string
	Cc       []string
	Bcc      []string
	Subject  string
	Body     string
	BodyType string

	FileAttachments []AttachmentParams
}

type AttachmentParams struct {
	AttachmentId AttachmentId
	Name         string
	ContentType  string
	Size         int64
	Content      []byte
}

type GetFolderParams struct {
	FolderId  string
	BaseShape BaseShape
}

type QueryMessageParams struct {
	FolderId  string
	StartTime time.Time
	EndTime   time.Time
	Limit     int
	Offset    int
}
