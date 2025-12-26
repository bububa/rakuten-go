package item20

// UpsertItem 用于创建或全量更新商品
// 对应 C# 的 UpsertItem : ItemCommon
type UpsertItem struct {
	ItemCommon
}

// PatchItem 用于部分更新商品
// 对应 C# 的 PatchItem : ItemCommon
type PatchItem struct {
	ItemCommon
}

// Item 代表从 API 获取到的完整商品信息
// 对应 C# 的 Item : ItemCommon
type Item struct {
	ItemCommon

	// 商品管理番号 (管理 URL 中的商品唯一标识)
	// [Required, MaxLength(32)]
	ManageNumber string `xml:"manageNumber" json:"manageNumber" validate:"required,max=32"`

	// 商品创建时间
	Created *string `xml:"created" json:"created"`

	// 商品更新时间
	Updated *string `xml:"updated" json:"updated"`
}

// ItemCommon 是基础结构体
type ItemCommon struct {
	// ItemNumber 商品编号 - MaxLength(32)
	ItemNumber string `xml:"itemNumber,omitempty" json:"itemNumber,omitempty" validate:"max=32"`

	// Title 商品名 - Required
	Title string `xml:"title" json:"title" validate:"required,max=255"`

	// Tagline キャッチコピー - MaxLength(174)
	Tagline string `xml:"tagline,omitempty" json:"tagline,omitempty" validate:"max=174"`

	// ProductDescription 商品说明文
	ProductDescription *ProductDescription `xml:"productDescription,omitempty" json:"productDescription,omitempty"`

	// SalesDescription PC用销售说明文 - MaxLength(10240)
	SalesDescription string `xml:"salesDescription,omitempty" json:"salesDescription,omitempty" validate:"max=10240"`

	// Precautions 医药品说明文
	Precautions *Precautions `xml:"precautions,omitempty" json:"precautions,omitempty"`

	// ImageType 商品种别 - Enum 映射为 string
	ItemType *string `xml:"itemType,omitempty" json:"itemType,omitempty" validate:"omitempty"`

	// Images 商品画像 - MaxOccur: 20
	Images []Image `xml:"images>image,omitempty" json:"images,omitempty" validate:"max=20"`

	// WhiteBgImage 白背景画像
	WhiteBgImage *WhiteBgImage `xml:"whiteBgImage,omitempty" json:"whiteBgImage,omitempty"`

	// Video 动画
	Video *Video `xml:"video,omitempty" json:"video,omitempty"`

	// GenreID ジャンルID - 正则校验 6位数字
	GenreID string `xml:"genreId" json:"genreId" validate:"required,len=6,regexp=^\\d{6}$"`

	// Tags 非制品属性标签ID - MaxOccur: 32
	Tags []int64 `xml:"tags>tag,omitempty" json:"tags,omitempty" validate:"max=32"`

	// HideItem 仓库指定 (隐藏商品)
	HideItem *bool `xml:"hideItem,omitempty" json:"hideItem,omitempty"`

	// UnlimitedInventoryFlag 在库设定无限制
	UnlimitedInventoryFlag *bool `xml:"unlimitedInventoryFlag,omitempty" json:"unlimitedInventoryFlag,omitempty"`

	// CustomizationOptions 商品选项
	CustomizationOptions *CustomizationOption `xml:"customizationOptions,omitempty" json:"customizationOptions,omitempty"`

	// ReleaseDate 预约商品发行日 - 乐天要求格式: yyyy-MM-dd+09:00
	// 建议使用指针，如果是 nil 则不序列化
	ReleaseDate *string `xml:"releaseDate,omitempty" json:"releaseDate,omitempty"`

	// PurchasablePeriod 销售期间指定
	PurchasablePeriod *Period `xml:"purchasablePeriod,omitempty" json:"purchasablePeriod,omitempty"`

	// Subscription 定期购买商品设定
	Subscription *Subscription `xml:"subscription,omitempty" json:"subscription,omitempty"`

	// BuyingClub 颁布会商品设定
	BuyingClub *BuyingClub `xml:"buyingClub,omitempty" json:"buyingClub,omitempty"`

	// Features 其他设定
	Features *Feature `xml:"features,omitempty" json:"features,omitempty"`

	// AccessControl 访问控制
	AccessControl *AccessControl `xml:"accessControl,omitempty" json:"accessControl,omitempty"`

	// Payment 结算情报
	Payment *Payment `xml:"payment,omitempty" json:"payment,omitempty"`

	// PointCampaign 积分变倍情报
	PointCampaign *PointCampaign `xml:"pointCampaign,omitempty" json:"pointCampaign,omitempty"`

	// ItemDisplaySequence 店铺内分类显示顺位
	ItemDisplaySequence *int64 `xml:"itemDisplaySequence,omitempty" json:"itemDisplaySequence,omitempty"`

	// Layout 版面布局设定
	Layout *Layout `xml:"layout,omitempty" json:"layout,omitempty"`

	// VariantSelectors 变体选择器
	VariantSelectors []VariantSelector `xml:"variantSelectors>variantSelector,omitempty" json:"variantSelectors,omitempty"`

	// Variants 变体数据 (C# Dictionary 映射为 Go Map)
	Variants map[string]Variant `xml:"variants,omitempty" json:"variants" validate:"required"`
}

// ProductDescription 商品说明文
type ProductDescription struct {
	// PC用商品説明文 (MaxLength: 10240)
	PC string `xml:"pc,omitempty" json:"pc,omitempty" validate:"max=10240"`

	// スマートフォン用商品説明文 (MaxLength: 10240)
	SP string `xml:"sp,omitempty" json:"sp,omitempty" validate:"max=10240"`
}

// Precautions 医药品说明文
type Precautions struct {
	// 医薬品説明文 (MaxLength: 20480)
	Description string `xml:"description,omitempty" json:"description,omitempty" validate:"max=20480"`

	// 医薬品注意事項 (MaxLength: 20480)
	Agreement string `xml:"agreement,omitempty" json:"agreement,omitempty" validate:"max=20480"`
}

// ImageType 模拟 C# 的枚举
type ImageType string

const (
	ImageTypeCabinet  ImageType = "CABINET"
	ImageTypeGold     ImageType = "GOLD"
	ImageTypeAbsolute ImageType = "ABSOLUTE"
)

// WhiteBgImage 白背景画像
type WhiteBgImage struct {
	// 画像種別 - 使用指针以支持可空性和必填校验
	Type *ImageType `xml:"type" json:"type" validate:"required"`

	// 画像URL (MaxLength: 255)
	Location string `xml:"location" json:"location" validate:"required,max=255"`
}

// Image 商品画像
// 通过嵌入 WhiteBgImage 模拟 C# 的继承 (Image : WhiteBgImage)
type Image struct {
	WhiteBgImage

	// 画像名（ALT）(MaxLength: 255)
	Alt string `xml:"alt,omitempty" json:"alt,omitempty" validate:"max=255"`
}

// VideoType 视频类型枚举
type VideoType string

const (
	VideoTypeHTML VideoType = "HTML"
)

// VideoParameters 视频参数包装器
type VideoParameters struct {
	// 動画のURL (MaxLength: 2048)
	// 虽然 C# 代码中这里带了 StringEnumConverter，但实际字段是 URL 字符串
	Value string `xml:"value" json:"value" validate:"required,max=2048"`
}

// Video 视频信息
type Video struct {
	// 動画種別 (通常固定为 HTML)
	Type *VideoType `xml:"type" json:"type" validate:"required"`

	// 動画パラメータ
	Parameters *VideoParameters `xml:"parameters" json:"parameters" validate:"required"`
}

// CustomizationInputType 模拟 C# 的枚举
type CustomizationInputType string

const (
	InputTypeSingleSelection   CustomizationInputType = "SINGLE_SELECTION"
	InputTypeMultipleSelection CustomizationInputType = "MULTIPLE_SELECTION"
	InputTypeFreeText          CustomizationInputType = "FREE_TEXT" // 常见补充类型
)

// CustomizationOption 商品选项（项目选择肢）
type CustomizationOption struct {
	// 商品オプション项目名 (MaxLength: 255)
	DisplayName string `xml:"displayName" json:"displayName" validate:"required,max=255"`

	// 商品オプション选择肢类型 (SINGLE_SELECTION / MULTIPLE_SELECTION)
	InputType *CustomizationInputType `xml:"inputType" json:"inputType" validate:"required"`

	// 商品オプション必填标志
	// 使用 *bool 以区分 false 和未设置(nil)
	Required *bool `xml:"required,omitempty" json:"required,omitempty"`

	// Select/Checkbox 用选择肢列表 (MaxOccur: 100)
	// XML 映射为 <selections><selection>...</selection></selections>
	Selections []OptionSelection `xml:"selections>selection,omitempty" json:"selections,omitempty" validate:"max=100"`
}

// OptionSelection 具体的选项内容（需根据乐天文档补充字段）
type OptionSelection struct {
	Value string `xml:"value" json:"value"`
}

// Period 销售期间/预约期间设定
type Period struct {
	// 開始日時 (格式: yyyy-MM-ddTHH:mm:ss+09:00)
	// 使用指针以支持 Required 校验和 omitempty
	Start *string `xml:"start" json:"start" validate:"required"`

	// 終了日時 (格式: yyyy-MM-ddTHH:mm:ss+09:00)
	End *string `xml:"end" json:"end" validate:"required"`
}

// Subscription 定期購入商品設定
type Subscription struct {
	// お届け日付指定フラグ (送货日期指定标志)
	// 使用指针以支持 omitempty，当字段为 nil 时不会出现在 JSON/XML 中
	ShippingDateFlag *bool `xml:"shippingDateFlag,omitempty" json:"shippingDateFlag,omitempty"`

	// お届け間隔指定フラグ (送货间隔指定标志)
	ShippingIntervalFlag *bool `xml:"shippingIntervalFlag,omitempty" json:"shippingIntervalFlag,omitempty"`
}

// BuyingClub 颁布会商品设定
type BuyingClub struct {
	// お届け回数 (配送次数)
	// [Required, Range(2, 12)]
	NumberOfDeliveries *int `xml:"numberOfDeliveries" json:"numberOfDeliveries" validate:"required,min=2,max=12"`

	// 商品内訳情報の表示（商品ページへの表示）
	DisplayItems *bool `xml:"displayItems,omitempty" json:"displayItems,omitempty"`

	// 商品内訳情報 (商品明细信息)
	// [ValidateObject(MaxOccur = 12)]
	Items []string `xml:"items>item,omitempty" json:"items,omitempty" validate:"max=12"`

	// お届け日付指定フラグ
	ShippingDateFlag *bool `xml:"shippingDateFlag,omitempty" json:"shippingDateFlag,omitempty"`

	// お届け間隔指定フラグ
	ShippingIntervalFlag *bool `xml:"shippingIntervalFlag,omitempty" json:"shippingIntervalFlag,omitempty"`
}

// SearchEnum 搜索显示枚举
type SearchEnum string

const (
	SearchVisible SearchEnum = "VISIBLE"
	SearchHidden  SearchEnum = "HIDDEN"
)

// InventoryDisplayEnum 库存显示枚举
type InventoryDisplayEnum string

const (
	InventoryDisplayHide     InventoryDisplayEnum = "HIDE"
	InventoryDisplayShow     InventoryDisplayEnum = "SHOW"
	InventoryDisplayLowStock InventoryDisplayEnum = "DISPLAY_LOW_STOCK"
)

// ReviewEnum 评价显示枚举
type ReviewEnum string

const (
	ReviewShow ReviewEnum = "SHOW"
	ReviewHide ReviewEnum = "HIDE"
)

// Feature 商品其他设定
type Feature struct {
	// サーチ表示 (搜索可见性)
	SearchVisibility *SearchEnum `xml:"searchVisibility,omitempty" json:"searchVisibility,omitempty"`

	// 注文ボタン (普通购买按钮)
	DisplayNormalCartButton *bool `xml:"displayNormalCartButton,omitempty" json:"displayNormalCartButton,omitempty"`

	// 定期購入ボタン (定期购买按钮)
	DisplaySubscriptionCartButton *bool `xml:"displaySubscriptionCartButton,omitempty" json:"displaySubscriptionCartButton,omitempty"`

	// 在庫数表示 (库存显示设定)
	InventoryDisplay *InventoryDisplayEnum `xml:"inventoryDisplay,omitempty" json:"inventoryDisplay,omitempty"`

	// 残り在庫数表示閾値 (库存报警阈值)
	// 当 InventoryDisplay 为 DISPLAY_LOW_STOCK 时必填，允许值：1～20
	LowStockThreshold *int `xml:"lowStockThreshold,omitempty" json:"lowStockThreshold,omitempty" validate:"omitempty,min=1,max=20"`

	// 商品問い合わせボタン (商品咨询按钮)
	ShopContact *bool `xml:"shopContact,omitempty" json:"shopContact,omitempty"`

	// レビュー表示设定
	Review *ReviewEnum `xml:"review,omitempty" json:"review,omitempty"`

	// メーカー提供情報表示 (厂家提供信息显示)
	DisplayManufacturerContents *bool `xml:"displayManufacturerContents,omitempty" json:"displayManufacturerContents,omitempty"`
}

// AccessControl アクセスコントロール
type AccessControl struct {
	// 闇市パスワード (暗市密码)
	// MaxLength(32), 正则表达式: [a-z0-9\-\_]+
	AccessPassword string `xml:"accessPassword,omitempty" json:"accessPassword,omitempty" validate:"omitempty,max=32,regexp=^[a-z0-9\\-\\_]+$"`
}

// Payment 決済情報
type Payment struct {
	// 決済情報 (是否含税)
	// true: 税込み, false: 税抜き
	TaxIncluded *bool `xml:"taxIncluded,omitempty" json:"taxIncluded,omitempty"`

	// 消費税税率
	// 允许值: "0" (非课税), "0.08" (8%), "0.1" (10%)
	// nil/空: 店铺设定 (按照乐天约定，如果不传则跟随店铺设定)
	TaxRate string `xml:"taxRate,omitempty" json:"taxRate,omitempty"`

	// 代引料 (是否包含货到付款手续费)
	CashOnDeliveryFeeIncluded *bool `xml:"cashOnDeliveryFeeIncluded,omitempty" json:"cashOnDeliveryFeeIncluded,omitempty"`
}

// PointCampaign 积分变倍信息
type PointCampaign struct {
	// ポイント変倍适用期间
	ApplicablePeriod Period `xml:"applicablePeriod" json:"applicablePeriod" validate:"required"`

	// ポイント情报 (积分倍率详情)
	Benefits Benefits `xml:"benefits" json:"benefits" validate:"required"`

	// 运用型ポイント情报 (积分上限设定)
	Optimization *Optimization `xml:"optimization,omitempty" json:"optimization,omitempty"`
}

// Benefits 积分倍率详情
type Benefits struct {
	// ポイント倍率 (例如: 2 表示 2倍积分)
	PointRate *int `xml:"pointRate" json:"pointRate" validate:"required"`
}

// Optimization 积分优化/上限设定
type Optimization struct {
	// ポイント上限倍率
	MaxPointRate *int `xml:"maxPointRate" json:"maxPointRate" validate:"required"`
}

// Layout レイアウト設定
type Layout struct {
	// 商品ページレイアウト (商品页面布局 ID)
	ItemLayoutID *int `xml:"itemLayoutId,omitempty" json:"itemLayoutId,omitempty"`

	// ヘッダー・フッター・レフトナビのテンプレートID (页眉/页脚/左侧导航模板 ID)
	NavigationID *int `xml:"navigationId,omitempty" json:"navigationId,omitempty"`

	// 表示項目の並び順テンプレートID (显示项目顺序模板 ID)
	LayoutSequenceID *int `xml:"layoutSequenceId,omitempty" json:"layoutSequenceId,omitempty"`

	// 共通説明文(小)テンプレートID (通用说明文-小 模板 ID)
	SmallDescriptionID *int `xml:"smallDescriptionId,omitempty" json:"smallDescriptionId,omitempty"`

	// 共通説明文(大)テンプレートID (通用说明文-大 模板 ID)
	LargeDescriptionID *int `xml:"largeDescriptionId,omitempty" json:"largeDescriptionId,omitempty"`

	// 共通説明文(小)テンプレートID (目次/Showcase 模板 ID)
	// 注：C# 原代码中注释与 smallDescriptionId 重复，此处根据字段名标记为 Showcase
	ShowcaseID *int `xml:"showcaseId,omitempty" json:"showcaseId,omitempty"`
}

// VariantSelectorValue 对应 C# 中的内部类 Value
type VariantSelectorValue struct {
	// バリエーション選択肢 (变体选项值，如 "红色", "XL")
	// MaxLength: 32
	DisplayValue string `xml:"displayValue" json:"displayValue" validate:"required,max=32"`
}

// VariantSelector バリエーション項目設定 (规格选择器)
type VariantSelector struct {
	// 共通説明文(小)テンプレートID (实际上是变体的 Key，如 "color" 或 "size")
	// 注意：C# 注释此处可能为拷贝错误，实际业务中这是变体的唯一标识键
	Key string `xml:"key" json:"key" validate:"required"`

	// バリエーション項目名 (变体展示名称，如 "カラー" 或 "サイズ")
	// MaxLength: 32
	DisplayName string `xml:"displayName" json:"displayName" validate:"required,max=32"`

	// バリエーション項目 (该维度下的所有选项值)
	// [ValidateObject(MaxOccur = 40)]
	// XML 映射为 <values><value>...</value></values>
	Values []VariantSelectorValue `xml:"values>value" json:"values" validate:"required,max=40"`
}
