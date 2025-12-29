package item20

import (
	"github.com/bububa/rakuten-go/model"
)

// DeleteRequest 商品管理番号を指定し、商品情報を削除 API Request
type DeleteRequest struct {
	model.JSONRequest
	// ManageNumber 商品管理番号
	ManageNumber string `json:"manageNumber,omitempty"`
}

// Encode implements GetRequest interface
func (r *DeleteRequest) Encode() []byte {
	return nil
}
