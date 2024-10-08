package goews

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

func getPTR[T any](s T) *T { return &s }

var (
	startBytes = []byte(`<?xml version="1.0" encoding="UTF-8"?>`)
)

type EnvelopeRequest struct {
	XMLName  xml.Name
	Soap     *Schema `xml:"xmlns:soap,attr"`
	Type     *Schema `xml:"xmlns:t,attr"`
	Messages *Schema `xml:"xmlns:m,attr"`
	Header   *Header
	Body     *Body
}

func (e *EnvelopeRequest) SetForMarshal() {
	e.XMLName.Local = `soap:Envelope`
}

type Body struct {
	XMLName xml.Name
	Body    interface{}
}

func (b *Body) SetForMarshal() {
	b.XMLName.Local = `soap:Body`
}

type Header struct {
	XMLName              xml.Name
	RequestServerVersion *RequestServerVersion `xml:"RequestServerVersion"`
}

func (h *Header) SetForMarshal() {
	h.XMLName.Local = `soap:Header`
}

type RequestServerVersion struct {
	XMLName xml.Name
	Version string `xml:"Version,attr,omitempty"`
}

func (r *RequestServerVersion) SetForMarshal() {
	r.XMLName.Local = `t:RequestServerVersion`
}

type Element interface {
	SetForMarshal()
}

func NewEnvelopeMarshal(body interface{}, schemas ...*Schema) (*EnvelopeRequest, error) {
	reTagXMLElement(body)
	res := &EnvelopeRequest{
		Soap: getPTR(SchemaSOAP),
		Header: &Header{
			RequestServerVersion: &RequestServerVersion{Version: "Exchange2013_SP1"},
		},
		Body: &Body{Body: body},
	}
	for _, schema := range schemas {
		switch *schema {
		case SchemaTypes:
			res.Type = schema
		case SchemaMessages:
			res.Messages = schema
		default:
			return nil, fmt.Errorf("unsupported schema name %v", *schema)
		}
	}
	if len(schemas) == 0 {
		res.Type = getPTR(SchemaTypes)
		res.Messages = getPTR(SchemaMessages)
	}
	return res, nil
}

func reTagXMLElement(target interface{}) {
	addrValue := reflect.ValueOf(target)
	if addrValue.Kind() != reflect.Ptr {
		return
	}
	targetValue := addrValue.Elem()
	if !targetValue.IsValid() {
		return
	}
	targetType := targetValue.Type()

	if targetType.Kind() == reflect.Ptr && !targetValue.IsNil() {
		targetValue = targetValue.Elem()
		if !targetValue.IsValid() {
			return
		}
		targetType = targetValue.Type()
	}

	if targetType.Kind() == reflect.Struct {
		for i := 0; i < targetType.NumField(); i++ {
			fValue := targetValue.Field(i)

			if !fValue.IsValid() {
				continue
			}

			if !fValue.CanAddr() {
				continue
			}

			if !fValue.Addr().CanInterface() {
				continue
			}
			if elIface, ok := fValue.Interface().(Element); ok && !fValue.IsNil() {
				elIface.SetForMarshal()
				fValue.Set(reflect.ValueOf(elIface))
			}
			reTagXMLElement(fValue.Addr().Interface())
		}
		return
	}

	if targetType.Kind() == reflect.Slice {
		for i := 0; i < targetValue.Len(); i++ {
			fValue := targetValue.Index(i)

			if !fValue.IsValid() {
				continue
			}

			if !fValue.CanAddr() {
				continue
			}

			if !fValue.Addr().CanInterface() {
				continue
			}

			if elIface, ok := fValue.Interface().(Element); ok && !fValue.IsNil() {
				elIface.SetForMarshal()
				fValue.Set(reflect.ValueOf(elIface))
			}
			reTagXMLElement(fValue.Addr().Interface())
		}
	}
}

func (e *EnvelopeRequest) GetEnvelopeBytes() ([]byte, error) {
	e.SetForMarshal()
	reTagXMLElement(e)
	res, err := xml.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("cant marshal %#v, err %v", *e, err.Error())
	}
	return append(startBytes, res...), nil
}
