package order

import (
	"github.com/bububa/rakuten-go/model"
	"github.com/bububa/rakuten-go/util"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// SearchRequest search order API Request
type SearchRequest struct {
	model.JSONRequest
	// OrderProgressList はステータスリストです。
	// ステータスは以下のいずれかが入ります。
	// 100: 注文確認待ち
	// 200: 楽天処理中
	// 300: 発送待ち
	// 400: 変更確定待ち
	// 500: 発送済
	// 600: 支払手続き中
	// 700: 支払手続き済
	// 800: キャンセル確定待ち
	// 900: キャンセル確定
	OrderProgressList []int `json:"orderProgressList,omitempty"`

	// SubStatusIdList はサブステータスIDリストです。
	// ユーザが作成したサブステータスを指定することができます。
	SubStatusIDList []int `json:"subStatusIdList,omitempty"`

	// DateType は期間検索種別です。
	// この項目は必須です。
	// 種別は以下のいずれかが入ります。
	// 1: 注文日
	// 2: 注文確認日
	// 3: 注文確定日
	// 4: 発送日
	// 5: 発送完了報告日
	// 6: 決済確定日
	DateType int `json:"dateType,omitempty"`

	// StartDatetime は期間検索開始日時です。過去2年以内の注文を指定することが可能です。
	// この項目は必須です。
	StartDatetime string `json:"startDatetime,omitempty"`

	// EndDatetime は期間検索終了日時です。開始日から63日以内を指定することができます。
	// この項目は必須です。
	EndDatetime string `json:"endDatetime,omitempty"`

	// OrderTypeList は販売種別リストです。以下のいずれかを指定することができます。
	// 1: 通常購入
	// 4: 定期購入
	// 5: 頒布会
	// 6: 予約商品
	OrderTypeList []int `json:"orderTypeList,omitempty"`

	// SettlementMethod は支払い方法名です。以下のいずれかを指定することができます。
	// 1: クレジットカード
	// 2: 代金引換
	// 3: 後払い
	// 4: ショッピングクレジット／ローン
	// 5: オートローン
	// 6: リース
	// 7: 請求書払い
	// 9: 銀行振込
	// 12: Apple Pay
	// 13: セブンイレブン（前払）
	// 14: ローソン、郵便局ATM等（前払）
	// 16: Alipay
	// 17: PayPal
	// 21: 後払い決済（楽天市場の共通決済）
	// 27: Alipay（支付宝）
	SettlementMethod int `json:"settlementMethod,omitempty"`

	// DeliveryName は配送方法です。
	DeliveryName string `json:"deliveryName,omitempty"`

	// ShippingDateBlankFlag は発送日未指定有無フラグです。以下のいずれかを指定することができます。
	// 0: 発送日の指定の有無によらず取得
	// 1: 発送日が未指定のものだけを取得
	ShippingDateBlankFlag *int `json:"shippingDateBlankFlag,omitempty"`

	// ShippingNumberBlankFlag はお荷物伝票番号未指定有無フラグです。以下のいずれかを指定することができます。
	// 0: お荷物伝票番号の指定の有無によらず取得
	// 1: お荷物伝票番号が未指定のものだけを取得
	ShippingNumberBlankFlag *int `json:"shippingNumberBlankFlag,omitempty"`

	// SearchKeywordType は検索キーワード種別です。次のいずれかを指定することができます。
	// 0: なし
	// 1: 商品名
	// 2: 商品番号
	// 3: ひとことメモ
	// 4: 注文者氏名
	// 5: 注文者氏名フリガナ
	// 6: 送付先氏名
	// 7: SKU管理番号
	// 8: システム連携用SKU番号
	// 9: SKU情報
	SearchKeywordType *int `json:"searchKeywordType,omitempty"`

	// SearchKeyword は検索キーワードです。32文字以下の入力を受け付けます。
	// 1: 商品名：1024 文字以下
	// 2: 商品番号：127文字以下
	// 3: ひとことメモ：1000文字以下
	// 4: 注文者氏名：254文字以下
	// 5: 注文者氏名フリガナ：254文字以下
	// 6: 送付先氏名：254文字以下
	// 7: SKU管理番号：40文字以下
	// 8: システム連携用SKU番号：96文字以下
	// 9: SKU情報：400文字以下
	SearchKeyword *string `json:"searchKeyword,omitempty"`

	// MailSendType は注文メールアドレス種別です。以下のいずれかを指定することができます。
	// 0: PC/モバイル
	// 1: PC
	// 2: モバイル
	MailSendType *int `json:"mailSendType,omitempty"`

	// OrdererMailAddress は注文者メールアドレスです。完全一致である必要があります。
	OrdererMailAddress string `json:"ordererMailAddress,omitempty"`

	// PhoneNumberType は電話番号種別です。以下のいずれかを指定することができます。
	// 0: 注文者
	// 1: 送付先
	PhoneNumberType *int `json:"phoneNumberType,omitempty"`

	// PhoneNumber は電話番号です。完全一致である必要があります。
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// ReserveNumber は申込番号です。完全一致である必要があります。
	ReserveNumber string `json:"reserveNumber,omitempty"`

	// PurchaseSiteType は購入サイトリストです。以下のいずれかを指定することがあります。
	// 0: すべて
	// 1: PCで注文
	// 2: モバイルで注文
	// 3: スマートフォンで注文
	// 4: タブレットで注文
	PurchaseSiteType *int `json:"purchaseSiteType,omitempty"`

	// AsurakuFlag はあす楽希望フラグです。以下のいずれかを指定することができます。
	// 0: あす楽希望の有無にかかわらず取得
	// 1: あす楽希望のものだけを取得
	AsurakuFlag *int `json:"asurakuFlag,omitempty"`

	// CouponUseFlag はクーポン利用有無フラグです。以下のいずれかを指定することができます。
	// 0: クーポン利用の有無にかかわらず取得
	// 1: クーポン利用のものだけを取得
	CouponUseFlag *int `json:"couponUseFlag,omitempty"`

	// DrugFlag は医薬品受注フラグです。以下のいずれかを指定することができます。
	// 0: 医薬品の有無にかかわらず取得
	// 1: 医薬品を含む注文だけを取得
	DrugFlag *int `json:"drugFlag,omitempty"`

	// OverseasFlag は海外カゴ注文フラグです。以下のいずれかを指定することができます。
	// 0: 海外カゴ注文の有無にかかわらず取得
	// 1: 海外カゴ注文のみ取得
	OverseasFlag *int `json:"overseasFlag,omitempty"`

	// SearchOrderPaginationRequestModel はページングに関する情報です。
	PaginationRequestModel *rmsModel.PaginationRequest `json:"PaginationRequestModel,omitempty"`
}

// Encode implements PostRequest interface
func (r *SearchRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// SearchResponse search order API Response
type SearchResponse struct {
	rmsModel.BaseResponse
	// CommonMessageModelResponseList はメッセージモデルリストです。ここにはエラー情報が含まれます。
	CommonMessageModelResponseList []rmsModel.Message `json:"MessageModelList,omitempty"`

	// OrderNumberList は注文番号リストです。該当する注文番号の一覧が取得できます。
	OrderNumberList []string `json:"orderNumberList"`

	// SearchOrderPaginationResponseModel はページングレスポンスモデルです。ページングに関する情報を取得することができます。
	PaginationResponseModel *rmsModel.PaginationResponse `json:"PaginationResponseModel,omitempty"`
}
