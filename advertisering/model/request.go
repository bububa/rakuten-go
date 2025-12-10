package model

import (
	"io"
)

// PostRequest post request interface
type PostRequest interface {
	ContentType() string
	// Encode encode request to bytes
	Encode() []byte
}

// GetRequest get request interface
type GetRequest interface {
	// Encode encode request to string
	Encode() string
}

type JSONRequest struct{}

func (r JSONRequest) ContentType() string {
	return "application/json"
}

type XMLRequest struct{}

func (r XMLRequest) ContentType() string {
	return "application/xml"
}

// UploadField multipart/form-data post request field struct
type UploadField struct {
	// Reader upload file reader
	Reader io.Reader
	// Key field key
	Key string
	// Value field value
	Value string
}

// UploadRequest multipart/form-data reqeust interface
type UploadRequest interface {
	// Encode encode request to UploadFields
	Encode() []UploadField
}
