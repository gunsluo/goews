package goews

type Schema string

const (
	SchemaSOAP     Schema = `http://schemas.xmlsoap.org/soap/envelope/`
	SchemaXsi      Schema = `http://www.w3.org/2001/XMLSchema-instance`
	SchemaTypes    Schema = `http://schemas.microsoft.com/exchange/services/2006/types`
	SchemaMessages Schema = `http://schemas.microsoft.com/exchange/services/2006/messages`
)
