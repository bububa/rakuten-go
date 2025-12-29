package itembundle

import "github.com/bububa/rakuten-go/model"

// Bundle 组合详情
type Bundle struct {
	BundleManageNumber     string         `json:"bundleManageNumber,omitempty"`
	BundleName             string         `json:"bundleName,omitempty"`
	BundleDescription      string         `json:"bundleDescription,omitempty"`
	BundleState            string         `json:"bundleState,omitempty"`
	ParentItemManageNumber string         `json:"parentItemManageNumber,omitempty"`
	BundleItems            []BundleItem   `json:"bundleItems,omitempty"`
	CreatedDate            model.DateTime `json:"createdDate,omitzero"`
	UpdatedDate            model.DateTime `json:"updatedDate,omitzero"`
}

// BundleItem 响应中的单项商品详情
type BundleItem struct {
	ItemManageNumber string `json:"itemManageNumber,omitempty"`
	IsDeletedItem    bool   `json:"isDeletedItem,omitempty"`
	Mandatory        bool   `json:"mandatory,omitempty"`
	Sequence         int    `json:"sequence,omitempty"`
}
