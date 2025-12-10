package postback

import (
	"github.com/bububa/rakuten-go/advertisering/model"
	"github.com/bububa/rakuten-go/util"
)

// Postback postback
type Postback struct {
	SID      uint64     `json:"sid,omitempty"`
	IsActive model.Bool `json:"is_active,omitempty"`
	URL      string     `json:"url,omitempty"`
}

type PostbackRequest struct {
	model.JSONRequest
	Postback
}

// Encode implements Post Request
func (r PostbackRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type PostbackResponse struct {
	model.BaseResponse
	Data *Postback `json:"data,omitempty"`
}
