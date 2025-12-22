package model

import (
	"github.com/bububa/rakuten-go/util"
)

type ResponseError struct {
	// Error  error code
	ErrorCode string `json:"error_code,omitempty" xml:"-"`
	// Message error description
	Message string `json:"message ,omitempty" xml:"-"`
}

// Error implement Response interface
func (r ResponseError) Error() string {
	return util.StringsJoin(r.ErrorCode, ":", r.Message)
}

// BaseResponse shared api response data fields
type BaseResponse struct {
	Results *ResponseError `json:"Results,omitempty"`
}

// IsError implement Response interface
func (r BaseResponse) IsError() bool {
	return r.Results != nil && r.Results.ErrorCode != ""
}

// Error implement Response interface
func (r BaseResponse) Error() string {
	if r.Results == nil {
		return ""
	}
	return util.StringsJoin(r.Results.ErrorCode, ":", r.Results.Message)
}
