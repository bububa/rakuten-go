package model

import (
	"bytes"
	"encoding/xml"
	"errors"

	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
)

type XMLRequest[T any] struct {
	model.XMLRequest
	XMLName xml.Name `xml:"request" json:"-"`
	Content T        `xml:"Content"`
}

func NewXMLRequest[T any](content T) *XMLRequest[T] {
	return &XMLRequest[T]{Content: content}
}

// Encode implements PostRequest
func (r *XMLRequest[T]) Encode() []byte {
	buf := new(bytes.Buffer)
	buf.Write([]byte(xml.Header))
	encoder := xml.NewEncoder(buf)
	_ = encoder.Encode(r)
	return buf.Bytes()
}

type XMLError struct {
	Code    string `json:"code,omitempty" xml:"code"`
	Message string `json:"message,omitempty" xml:"message"`
}

func (e XMLError) IsError() bool {
	return e.Code != ""
}

func (e XMLError) Error() string {
	return util.StringsJoin(e.Code, ":", e.Message)
}

// XMLErrorResult mimics the base class with Error list
type XMLErrorResult struct {
	// The tag "errors>error" creates the <errors><error>...</error></errors> hierarchy
	Errors []XMLError `json:"errors,omitempty" xml:"errors>error,omitempty"`
}

func (r XMLErrorResult) IsError() bool {
	return len(r.Errors) > 0
}

func (r XMLErrorResult) Error() string {
	var err error
	for _, v := range r.Errors {
		err = errors.Join(err, v)
	}
	return err.Error()
}

type EmptyXMLResponse struct {
	XMLName xml.Name `xml:"result" json:"-"`
	XMLErrorResult
}

type XMLResponse[T any] struct {
	XMLName xml.Name `xml:"result" json:"-"`
	XMLErrorResult
	Content T `json:"content,omitempty" xml:"Content"`
}
