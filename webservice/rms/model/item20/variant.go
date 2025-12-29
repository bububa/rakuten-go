package item20

// Variant SKU 变体信息
type Variant struct {
	// 自定义 SKU ID (MaxLength: 96)
	MerchantDefinedSkuID string `xml:"merchantDefinedSkuId,omitempty" json:"merchantDefinedSkuId,omitempty" validate:"omitempty,max=96"`

	// SKU 属性值映射 (例如: {"color": "Red", "size": "L"})
	// 对应 C# 的 Dictionary<string, string>
	SelectorValues map[string]string `xml:"selectorValues" json:"selectorValues"`

	// SKU 图片 (通常乐天 API 2.0 每个 SKU 只允许 1 张图)
	Images []Image `xml:"images>image,omitempty" json:"images,omitempty" validate:"max=1"`

	// 在库戻しフラグ (取消订单后是否自动归还库存)
	RestockOnCancel *bool `xml:"restockOnCancel,omitempty" json:"restockOnCancel,omitempty"`

	// 缺货时是否接受订单
	BackOrderFlag *bool `xml:"backOrderFlag,omitempty" json:"backOrderFlag,omitempty"`

	// 在库有时的交付日期管理 ID
	NormalDeliveryDateID *int `xml:"normalDeliveryDateId,omitempty" json:"normalDeliveryDateId,omitempty"`

	// 缺货（预约）时的交付日期管理 ID
	BackOrderDeliveryDateID *int `xml:"backOrderDeliveryDateId,omitempty" json:"backOrderDeliveryDateId,omitempty"`

	// 最大购买数量限制
	OrderQuantityLimit *int `xml:"orderQuantityLimit,omitempty" json:"orderQuantityLimit,omitempty"`

	// 参考价格 (如：厂家希望零售价)
	ReferencePrice *ReferencePrice `xml:"referencePrice,omitempty" json:"referencePrice,omitempty"`

	// SKU 特色设定 (内部类映射)
	Features *VariantFeature `xml:"features,omitempty" json:"features,omitempty"`

	// 是否隐藏该 SKU (仓库存放)
	Hidden *bool `xml:"hidden,omitempty" json:"hidden,omitempty"`

	// 售价 (由于乐天 API 接受数字字符串，使用 string 保证精度)
	// 正则: \d{1,9}
	StandardPrice string `xml:"standardPrice,omitempty" json:"standardPrice,omitempty" validate:"omitempty,regexp=^\\d{1,9}$"`

	// 定期购买销售价格设定
	SubscriptionPrice *SubscriptionPrice `xml:"subscriptionPrice,omitempty" json:"subscriptionPrice,omitempty"`

	// 组合商品用目录 ID 列表
	ArticleNumberForSet []string `xml:"articleNumberForSet>articleNumber,omitempty" json:"articleNumberForSet,omitempty" validate:"max=20"`

	// 目录 ID 信息 (JAN/EAN/UPC 等)
	ArticleNumber *ArticleNumber `xml:"articleNumber,omitempty" json:"articleNumber,omitempty"`

	// 配送信息
	Shipping *Shipping `xml:"shipping,omitempty" json:"shipping,omitempty"`

	// 规格详情 (Spec)
	Specs []VariantSpec `xml:"specs>spec,omitempty" json:"specs,omitempty"`

	// 属性信息 (Attribute)
	Attributes []VariantAttribute `xml:"attributes>attribute,omitempty" json:"attributes,omitempty"`
}

// VariantFeature 对应 C# 的 Variant.Feature
type VariantFeature struct {
	RestockNotification *bool `xml:"restockNotification,omitempty" json:"restockNotification,omitempty"`
	Noshi               *bool `xml:"noshi,omitempty" json:"noshi,omitempty"`
}

// VariantSpec 对应 C# 的 Variant.Spec
type VariantSpec struct {
	Label string `xml:"label" json:"label" validate:"required,max=40"`
	Value string `xml:"value" json:"value" validate:"required,max=140"`
}

// VariantAttribute 对应 C# 的 Variant.Attribute
type VariantAttribute struct {
	Name   string   `xml:"name" json:"name" validate:"required"`
	Values []string `xml:"values>value" json:"values" validate:"required"`
	Unit   string   `xml:"unit,omitempty" json:"unit,omitempty"`
}

// ReferencePriceDisplayType 表示価格種别枚举
type ReferencePriceDisplayType string

const (
	DisplayTypeReferencePrice ReferencePriceDisplayType = "REFERENCE_PRICE" // 选取的显示价格文案
	DisplayTypeShopSetting    ReferencePriceDisplayType = "SHOP_SETTING"    // 遵循店铺设定
	DisplayTypeOpenPrice      ReferencePriceDisplayType = "OPEN_PRICE"      // 开放价格
)

// ReferencePrice 表示価格（二重価格设定）
type ReferencePrice struct {
	// 表示価格種別
	// [Required]
	DisplayType *ReferencePriceDisplayType `xml:"displayType" json:"displayType" validate:"required"`

	// 表示価格文言 (1: 本店通常价格, 2: 厂家希望零售价, 4: 参考商品价格导航数据)
	// 条件必填: 当 DisplayType 为 REFERENCE_PRICE 时必填
	Type *int `xml:"type,omitempty" json:"type,omitempty"`

	// 表示価格 (具体的数值)
	// 条件必填: 当 DisplayType 为 REFERENCE_PRICE 且 Type 不为 4 时必填
	// 允许值: 1～999999999
	Value string `xml:"value,omitempty" json:"value,omitempty" validate:"omitempty,regexp=^\\d{1,9}$"`
}

// SubscriptionPrice 定期購入販売価格設定
type SubscriptionPrice struct {
	// 基礎価格 (定期购买的基础售价)
	// [Required]
	BasePrice string `xml:"basePrice" json:"basePrice" validate:"required"`

	// 回別価格 (个别轮次价格设定)
	// 对应 C# 中的内部类 IndividualPrices
	Individual *IndividualPrices `xml:"individual,omitempty" json:"individual,omitempty"`
}

// IndividualPrices 个别轮次价格设定
type IndividualPrices struct {
	// 初回価格 (首回折扣价)
	FirstPrice string `xml:"firstPrice,omitempty" json:"firstPrice,omitempty"`
}

// ArticleNumber カタログID情報
type ArticleNumber struct {
	// カタログID (JAN/EAN/UPC/ISBN 等)
	// 条件必填：如果不提供 exemptionReason，则此项必填
	Value string `xml:"value,omitempty" json:"value,omitempty"`

	// カタログIDなしの理由 (无目录 ID 的理由)
	// 常见值：
	// 1: 厂家没有提供 ID
	// 2: 属于定制商品
	// 3: 属于套装商品
	ExemptionReason *int `xml:"exemptionReason,omitempty" json:"exemptionReason,omitempty"`
}

// Shipping 配送情報
type Shipping struct {
	// 個別送料 (个别运费金额)
	Fee string `xml:"fee,omitempty" json:"fee,omitempty"`

	// 送料無料フラグ (免运费标志)
	// true: 送料込 (含运费/免运), false: 送料別 (运费另计)
	PostageIncluded *bool `xml:"postageIncluded,omitempty" json:"postageIncluded,omitempty"`

	// 地域別個別送料管理番号 (按地区设置的个别运费模板 ID)
	ShopAreaSoryoPatternID *int `xml:"shopAreaSoryoPatternId,omitempty" json:"shopAreaSoryoPatternId,omitempty"`

	// 配送方法セット管理番号 (配送方法组管理编号)
	ShippingMethodGroup string `xml:"shippingMethodGroup,omitempty" json:"shippingMethodGroup,omitempty"`

	// 送料区分情報 (运费区分信息 - 通常涉及特定尺寸或重量分类)
	PostageSegment *PostageSegment `xml:"postageSegment,omitempty" json:"postageSegment,omitempty"`

	// 海外配送管理番号 (海外配送模板 ID)
	OverseasDeliveryID *int `xml:"overseasDeliveryId,omitempty" json:"overseasDeliveryId,omitempty"`

	// 単品配送設定 (单品配送设定)
	// 0: 不设定, 1: 设定
	SingleItemShipping *int `xml:"singleItemShipping,omitempty" json:"singleItemShipping,omitempty"`
}

// PostageSegment 运费区分详细（需根据具体 API 定义补充）
type PostageSegment struct {
	LocalSegment int `xml:"localSegment,omitempty" json:"localSegment,omitempty"`
}
