package item20

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
)

// PatchRequest 商品管理番号を指定し、商品情報の部分更新 API Request
type PatchRequest struct {
	model.JSONRequest
	ManageNumber string     `json:"manageNumber,omitempty"`
	Item         *PatchItem `json:"item,omitempty"`
}

// Encode implements PostRequest
func (r *PatchRequest) Encode() []byte {
	return util.JSONMarshal(r.Item)
}
