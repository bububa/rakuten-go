package item20

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
)

// UpsertRequest 商品管理番号を指定し、商品情報の登録・全項目の更新 API Request
type UpsertRequest struct {
	model.JSONRequest
	ManageNumber string      `json:"manageNumber,omitempty"`
	Item         *UpsertItem `json:"item,omitempty"`
}

// Encode implements PostRequest
func (r *UpsertRequest) Encode() []byte {
	return util.JSONMarshal(r.Item)
}
