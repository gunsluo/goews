package goews

const (
	BodyTypeHtml = "HTML"
	BodyTypeText = "Text"
)

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
