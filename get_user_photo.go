package goews

type GetUserPhotoRequest struct {
	XMLName       struct{} `xml:"m:GetUserPhoto"`
	Email         string   `xml:"m:Email"`
	SizeRequested string   `xml:"m:SizeRequested"`
}

type GetUserPhotoResponse struct {
	Response
	HasChanged  bool   `xml:"HasChanged"`
	PictureData string `xml:"PictureData"`
}

type getUserPhotoResponseEnvelop struct {
	XMLName struct{}                 `xml:"Envelope"`
	Body    getUserPhotoResponseBody `xml:"Body"`
}
type getUserPhotoResponseBody struct {
	GetUserPhotoResponse GetUserPhotoResponse `xml:"GetUserPhotoResponse"`
}
