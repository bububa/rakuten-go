package item20

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// BulkGetRequest 商品管理番号を指定し、商品情報を取得 API Request
type BulkGetRequest struct {
	model.JSONRequest
	// ManageNumber 商品管理番号
	ManageNumbers []string `json:"manageNumbers,omitempty"`
}

// Encode implements PostRequest interface
func (r *BulkGetRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BulkGetResponse は楽天ペイ受注APIの注文情報取得時の検索条件です API Response
type BulkGetResponse struct {
	rmsModel.BaseResponse
	Result []Item `json:"result,omitempty"`
}
