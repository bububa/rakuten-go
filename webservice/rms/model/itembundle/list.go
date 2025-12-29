package itembundle

import (
	"strconv"

	"github.com/bububa/rakuten-go/util"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// BundleState 状态：ACTIVE (有效), INACTIVE (无效)
type BundleState string

const (
	// BundleStateActive 有效
	BundleStateActive BundleState = "ACTIVE"
	// BundleStateInactive 无效
	BundleStateInactive BundleState = "INACTIVE"
)

// BundleType 类型：parent (亲商品), child (子商品)
type BundleType string

const (
	// BundleTypeParent 亲商品
	BundleTypeParent BundleType = "parent"
	// BundleTypeChild 子商品
	BundleTypeChild BundleType = "child"
)

// ListRequest 组合查询请求
type ListRequest struct {
	// BundleState 状态：ACTIVE (有效), INACTIVE (无效)
	BundleState BundleState `json:"bundleState,omitempty"`
	PageSize    int         `json:"pageSize,omitempty"`
	PageNumber  int         `json:"pageNumber,omitempty"`
	// ItemManageNumber 商品管理编号
	ItemManageNumber string `json:"itemManageNumber,omitempty"`
	// Type 类型：parent (亲商品), child (子商品)
	Type       BundleType `json:"type,omitempty"`
	BundleName string     `json:"bundleName,omitempty"`
}

// Encode implements GetRequest
func (r *ListRequest) Encode() string {
	values := util.NewURLValues()
	if r.BundleState != "" {
		values.Set("bundleState", string(r.BundleState))
	}
	if r.PageSize > 0 {
		values.Set("pageSize", strconv.Itoa(r.PageSize))
	}
	if r.PageNumber > 0 {
		values.Set("pageNumber", strconv.Itoa(r.PageNumber))
	}
	if r.ItemManageNumber != "" {
		values.Set("itemManageNumber", r.ItemManageNumber)
	}
	if r.Type != "" {
		values.Set("type", string(r.Type))
	}
	if r.BundleName != "" {
		values.Set("bundleName", r.BundleName)
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

// ListResponse 组合列表响应包装
type ListResponse struct {
	rmsModel.BaseResponse
	NumberOfBundles int      `json:"numberOfBundles"`
	Bundles         []Bundle `json:"bundles"`
}
