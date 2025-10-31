package model

import (
	"io"
)

type Request interface {
	ContentType() string
}

// PostRequest post request interface
type PostRequest interface {
	Request
	// Encode encode request to bytes
	Encode() []byte
}

// GetRequest get request interface
type GetRequest interface {
	Request
	// Encode encode request to string
	Encode() string
}

type JSONRequest struct{}

func (r JSONRequest) ContentType() string {
	return "application/json"
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
