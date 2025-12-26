package item20

import (
	"github.com/bububa/rakuten-go/model"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// GetRequest 商品管理番号を指定し、商品情報を取得 API Request
type GetRequest struct {
	model.JSONRequest
	// ManageNumber 商品管理番号
	ManageNumber string `json:"manageNumber,omitempty"`
}

// Encode implements GetRequest interface
func (r *GetRequest) Encode() string {
	return ""
}

// GetResponse は楽天ペイ受注APIの注文情報取得時の検索条件です API Response
type GetResponse struct {
	rmsModel.BaseResponse
	Item
}
