package model

import (
	"io"

	"github.com/bububa/rakuten-go/util"
)

// Response api response interface
type Response interface {
	IsError() bool
	Error() string
}

type Metadata struct {
	// APINameVersion api_name_version
	APINameVersion string `json:"api_name_version,omitempty"`
	// Links _links
	Links struct {
		Prev string `json:"prev,omitempty"`
		Self string `json:"self,omitempty"`
		Next string `json:"next,omitempty"`
	} `json:"_links"`
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	Total int `json:"total,omitempty"`
}

type ResponseError struct {
	// Error  error code
	ErrorCode string `json:"error,omitempty"`
	// ErrorDescription error description
	ErrorDescription string `json:"error_description,omitempty"`
}

// Error implement Response interface
func (r ResponseError) Error() string {
	return util.StringsJoin(r.ErrorCode, ":", r.ErrorDescription)
}

// BaseResponse shared api response data fields
type BaseResponse struct {
	ResponseError
	Metadata *Metadata `json:"_metadata,omitempty"`
}

// IsError implement Response interface
func (r BaseResponse) IsError() bool {
	return r.ErrorCode != ""
}

type Unmarshaler interface {
	Unmarshal(r io.Reader) error
}
