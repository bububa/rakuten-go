package item20

import (
	"github.com/bububa/rakuten-go/util/query"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// ItemType 商品类型枚举
type ItemType string

const (
	// NORMAL 通常商品
	ItemTypeNormal ItemType = "NORMAL"

	// PRE_ORDER 予約商品
	ItemTypePreOrder ItemType = "PRE_ORDER"

	// BUYING_CLUB 定期購入/頒布会商品
	ItemTypeBuyingClub ItemType = "BUYING_CLUB"
)

type SearchSortKey string

const (
	SortKeyUpdated             SearchSortKey = "updated"
	SortKeyCreated             SearchSortKey = "created"
	SortKeyItemDisplaySequence SearchSortKey = "itemDisplaySequence"
	SortKeyManageNumber        SearchSortKey = "manageNumber"
	SortKeyPurchasableStart    SearchSortKey = "purchasablePeriodStart"
	SortKeyPurchasableEnd      SearchSortKey = "purchasablePeriodEnd"
	SortKeyPointCampaignStart  SearchSortKey = "pointCampaignStart"
	SortKeyPointCampaignEnd    SearchSortKey = "pointCampaignEnd"
	SortKeyPointRate           SearchSortKey = "pointRate"
	SortKeyReviewCount         SearchSortKey = "reviewCount"
	SortKeyReviewAverageRating SearchSortKey = "reviewAverageRating"
)

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

// SearchRequest 指定した条件から通常商品・予約商品・定期購入商品の商品情報を検索 API Request
type SearchRequest struct {
	Title                     string        `json:"title,omitempty"`
	Tagline                   string        `json:"tagline,omitempty"`
	ManageNumber              string        `json:"manageNumber,omitempty"`
	ItemNumber                string        `json:"itemNumber,omitempty"`
	ArticleNumber             string        `json:"articleNumber,omitempty"`
	VariantID                 string        `json:"variantId,omitempty"`
	MerchantDefinedSkuID      string        `json:"merchantDefinedSkuId,omitempty"`
	GenreID                   string        `json:"genreId,omitempty"`
	ItemType                  *ItemType     `json:"itemType,omitempty"`
	StandardPriceFrom         *int64        `json:"standardPriceFrom,omitempty"`
	StandardPriceTo           *int64        `json:"standardPriceTo,omitempty"`
	IsVariantStockout         *bool         `json:"isVariantStockout,omitempty"`
	IsItemStockout            *bool         `json:"isItemStockout,omitempty"`
	PurchasablePeriodFrom     *string       `json:"purchasablePeriodFrom,omitempty"`
	PurchasablePeriodTo       *string       `json:"purchasablePeriodTo,omitempty"`
	IsHiddenItem              *bool         `json:"isHiddenItem,omitempty"`
	IsHiddenVariant           *bool         `json:"isHiddenVariant,omitempty"`
	IsSearchable              *bool         `json:"isSearchable,omitempty"`
	IsYamiichi                *bool         `json:"isYamiichi,omitempty"`
	PointApplicablePeriodFrom *string       `json:"pointApplicablePeriodFrom,omitempty"`
	PointApplicablePeriodTo   *string       `json:"pointApplicablePeriodTo,omitempty"`
	IsOptimizedPoint          *bool         `json:"isOptimizedPoint,omitempty"`
	PointRate                 *int          `json:"pointRate,omitempty"`
	MaxPointRate              *int          `json:"maxPointRate,omitempty"`
	CategoryID                string        `json:"categoryId,omitempty"`
	IsBackOrder               *bool         `json:"isBackOrder,omitempty"`
	IsPostageIncluded         *bool         `json:"isPostageIncluded,omitempty"`
	CreatedFrom               *string       `json:"createdFrom,omitempty"`
	CreatedTo                 *string       `json:"createdTo,omitempty"`
	UpdatedFrom               *string       `json:"updatedFrom,omitempty"`
	UpdatedTo                 *string       `json:"updatedTo,omitempty"`
	SortKey                   SearchSortKey `json:"sortKey,omitempty"`
	SortOrder                 SortOrder     `json:"sortOrder,omitempty"`
	Offset                    *int          `json:"offset,omitempty"`
	Hits                      *int          `json:"hits,omitempty"`
	CursorMark                string        `json:"cursorMark,omitempty"`
	IsCategoryIncluded        *bool         `json:"isCategoryIncluded,omitempty"`
	IsReviewIncluded          *bool         `json:"isReviewIncluded,omitempty"`
	IsInventoryIncluded       *bool         `json:"isInventoryIncluded,omitempty"`
	IsSubscription            *bool         `json:"isSubscription,omitempty"`
	BasePriceFrom             *int          `json:"basePriceFrom,omitempty"`
	BasePriceTo               *int          `json:"basePriceTo,omitempty"`
}

// Encode implements GetRequest
func (r *SearchRequest) Encode() string {
	values, _ := query.Values(r)
	return values.Encode()
}

// SearchResponse 搜索结果
type SearchResponse struct {
	rmsModel.BaseResponse
	NumFound       int                `json:"numFound"`
	Offset         int                `json:"offset"`
	NextCursorMark string             `json:"nextCursorMark"`
	Results        []SearchResultItem `json:"results"`
}

type SearchResultItem struct {
	Item      Item                   `json:"item"`
	Category  *SearchResultCategory  `json:"category,omitempty"`
	Review    *SearchResultReview    `json:"review,omitempty"`
	Inventory *InventorySkuContainer `json:"inventory,omitempty"`
}

type SearchResultCategory struct {
	CategoryIds []string `json:"categoryIds"`
}

type SearchResultReview struct {
	Count         *int     `json:"count,omitempty"`
	AverageRating *float32 `json:"averageRating,omitempty"`
}

// InventorySkuContainer 重点：Go 处理 C# 字典继承的方式
type InventorySkuContainer struct {
	// SKU数据：Key 是变体 ID (variantId)
	Skus    map[string]InventorySku `json:"skus"` // 在实际 JSON 中可能是平铺的，需根据 API 响应调整标签
	Created string                  `json:"created"`
	Updated string                  `json:"updated"`
}

type InventorySku struct {
	OperationLeadTime OperationLeadTime `json:"operationLeadTime"`
	ShipFromIds       []int             `json:"shipFromIds"`
}

type OperationLeadTime struct {
	NormalDeliveryTimeID    *int `json:"normalDeliveryTimeId,omitempty"`
	BackOrderDeliveryTimeID *int `json:"backOrderDeliveryTimeId,omitempty"`
}
