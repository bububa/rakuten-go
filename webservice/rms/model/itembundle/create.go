package itembundle

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// CreateRequest 创建/更新组合的请求体
type CreateRequest struct {
	model.JSONRequest
	// 亲商品管理编号 (Max 32)
	ParentItemManageNumber string `json:"parentItemManageNumber,omitempty"`
	// 组合管理编号 (唯一标识, Max 64)
	BundleManageNumber string `json:"bundleManageNumber,omitempty"`
	// 表示设定：ACTIVE, INACTIVE
	BundleState string `json:"bundleState,omitempty"`
	// 组合内商品列表 (包含亲商品本身)
	BundleItems []BundleItemToCreate `json:"bundleItems,omitempty"`
	// 组合管理名称 (Max 32)
	BundleName string `json:"bundleName,omitempty"`
	// 组合销售描述文 (Max 50)
	BundleDescription string `json:"bundleDescription,omitempty"`
}

// BundleItemToCreate 组合内的单项商品
type BundleItemToCreate struct {
	ItemManageNumber string `json:"itemManageNumber"`
	// Mandatory 是否必选
	Mandatory *bool `json:"mandatory,omitempty"`
	// Sequence 显示顺序
	Sequence *int `json:"sequence,omitempty"`
}

// Encode implements PostRequest
func (r *CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type CreateResponse struct {
	rmsModel.BaseResponse
	Bundle
}
