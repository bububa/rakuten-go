package model

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util/query"
	"github.com/bububa/rakuten-go/webservice/ichiba"
)

// TagSearchRequest Rakuten Tag Search API Request
type TagSearchRequest struct {
	ichiba.BaseRequest
	// TagID tag id
	// Tag IDs can be obtained from the Rakuten Ichiba product search API and Rakuten Ichiba genre search API.
	// 0 cannot be specified.
	// Up to 10 tag IDs can be specified, separated by commas (,).
	TagID int64 `json:"tagId,omitempty"`
}

// Encode implements GetRequest interface
func (r *TagSearchRequest) Encode() string {
	values, err := query.Values(r)
	if err != nil {
		return ""
	}
	return values.Encode()
}

// TagSearchResponse Rakuten Tag Search API Response
type TagSearchResponse struct {
	model.BaseResponse
	// TagGroups tag groups
	TagGroups []TagGroup `json:"tagGroups,omitempty"`
}

// TagGroup tag group
type TagGroup struct {
	// TagGroupName Tag Group Name
	TagGroupName string `json:"tagGroupName,omitempty"`
	// TagGroupID Tag Group ID
	TagGroupID model.Int64 `json:"tagGroupId,omitempty"`
	// Tags tags
	Tags []Tag `json:"tags,omitempty"`
}

// Tag Tag Information
type Tag struct {
	// TagID tag ID
	TagID model.Int64 `json:"tagId,omitempty"`
	// TagName tag name
	TagName string `json:"tagName,omitempty"`
	// ParentTagID parent tag ID
	ParentTagID model.Int64 `json:"parentTagId,omitempty"`
}
