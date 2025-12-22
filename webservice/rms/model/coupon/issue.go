package coupon

import (
	"encoding/xml"

	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

// IssueRequest Issue Coupon Request
type IssueRequest struct {
	XMLName xml.Name       `xml:"couponIssueRequest" json:"-"`
	Coupon  *CouponToIssue `json:"coupon,omitempty" xml:"coupon"`
}

func (r *IssueRequest) ContentType() string {
	return "application/xml"
}

func (r *IssueRequest) Encode() []byte {
	req := rmsModel.NewXMLRequest(r)
	return req.Encode()
}

type Coupon struct {
	CouponToIssue
	CouponCode string `xml:"couponCode" json:"couponCode,omitempty"`
}

type CouponToIssue struct {
	CouponName      string `xml:"couponName" json:"couponName"`
	CouponCaption   string `xml:"couponCaption" json:"couponCaption"`
	CouponStartDate string `xml:"couponStartDate" json:"couponStartDate"`
	CouponEndDate   string `xml:"couponEndDate" json:"couponEndDate"`
	CouponImage     string `xml:"couponImage,omitempty" json:"couponImage,omitempty"`
	IssueCount      int    `xml:"issueCount" json:"issueCount,omitempty"`
	// ItemType         discountTypeによって、値が異なります。
	// 1： 単一商品
	// 3： 複数商品
	// 4： 受注
	// 5： 送料無料
	ItemType int `xml:"itemType" json:"itemType,omitempty"`
	// DiscountType
	// 1： 定額値引き
	// 2： 定率値引き
	// 4： 送料無料
	DiscountType        int `xml:"discountType" json:"discountType,omitempty"`
	DiscountFactor      int `xml:"discountFactor" json:"discountFactor,omitempty"`
	MemberAvailMaxCount int `xml:"memberAvailMaxCount" json:"memberAvailMaxCount,omitempty"`
	// MultRankCard
	// 0： 条件なし
	// 1： レギュラー
	// 2： シルバー
	// 3： ゴールド
	// 4： プラチナ
	// 5： ダイヤモンド
	MultiRankCond []int `xml:"multiRankCond>rankCond,omitempty" json:"multiRankCond,omitempty"`
	CombineFlag   int   `xml:"combineFlag" json:"combineFlag,omitempty"`
	DisplayFlag   int   `xml:"displayFlag" json:"displayFlag,omitempty"`
	// Items 対象商品は3000個まで指定可能です。
	Items           []Item           `xml:"items>item,omitempty" json:"items,omitempty"`
	OtherConditions *OtherConditions `xml:"otherConditions,omitempty" json:"otherConditions,omitempty"`
}

// Item 対象商品
type Item struct {
	ItemURL string `xml:"itemUrl" json:"itemUrl"`
}

type OtherConditions struct {
	OtherCondition OtherCondition `xml:"otherCondition" json:"otherCondition"`
}

type OtherCondition struct {
	// ConditionTypeCode
	// RS001： デバイス指定（設定のない場合は自動的に設定される）
	// RS002： 販売方法(設定のない場合は自動的に設定される）
	// RS003： 利用金額
	// RS004： 利用個数
	// ※39ショップに限り、 利用個数(RS004) を設定した場合、複数商品(「itemType」に「3」)が設定可能です。
	// ※ 対象商品リストが設定されている場合はクーポン発行条件を参照ください。
	ConditionTypeCode string `xml:"conditionTypeCode" json:"conditionTypeCode,omitempty"`
	// StartValue
	// 0： PC （設定のない場合は自動的に設定される）
	// 1： モバイル （設定のない場合は自動的に設定される）
	// RS002	0： 通常購入
	// RS003	1 ～ 999999999： 金額
	// RS004	0 ～ 999999999： 個数
	StartValue string `xml:"startValue" json:"startValue,omitempty"`
}

type IssuedCoupon struct {
	XMLName xml.Name `xml:"coupon" json:"-"`
	// CouponCode クーポンコード
	CouponCode string `xml:"couponCode" json:"couponCode,omitempty"`
	// PCGetURL 取得URL (PC)
	PCGetURL string `xml:"pcGetUrl" json:"picGetUrl,omitempty"`
}

// IssueResponse Issue Coupon Request
type IssueResponse rmsModel.XMLResponse[*IssuedCoupon]
