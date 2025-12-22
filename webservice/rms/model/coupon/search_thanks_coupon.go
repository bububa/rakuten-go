package coupon

import (
	"encoding/xml"
	"strconv"

	"github.com/bububa/rakuten-go/util"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

type SearchThanksCouponRequest struct {
	IssueStats     *int   `xml:"issueStats,omitempty"`
	GrantStartDate string `xml:"grantStartDate,omitempty" json:"grantStartDate,omitempty"`
	GrantEndDate   string `xml:"grantEndDate,omitempty" json:"grantEndDate,omitempty"`
	RegDate        string `xml:"regDate,omitempty" json:"regDate,omitempty"`
	Hits           int    `xml:"hits,omitempty" json:"hits,omitempty"`
	Page           int    `xml:"page,omitempty" json:"page,omitempty"`
}

// Encode implements GetRequest
func (r *SearchThanksCouponRequest) Encode() string {
	values := util.NewURLValues()
	if r.IssueStats != nil {
		values.Set("issueStats", strconv.Itoa(*r.IssueStats))
	}
	if r.GrantStartDate != "" {
		values.Set("grantStartDate", r.GrantStartDate)
	}
	if r.GrantEndDate != "" {
		values.Set("grantEndDate", r.GrantEndDate)
	}
	if r.RegDate != "" {
		values.Set("regDate", r.RegDate)
	}
	if r.Hits > 0 {
		values.Set("hits", strconv.Itoa(r.Hits))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}

type SearchThanksCouponResponse struct {
	rmsModel.XMLResponse[*SearchedThanksCoupons]
	AllCount int `xml:"allCount"`
}

type SearchedThanksCoupons struct {
	XMLName       xml.Name       `xml:"thanksCoupons"`
	ThanksCoupons []ThanksCoupon `xml:"thanksCoupon"`
}

type ThanksOtherCondition struct {
	ConditionTypeCode string `xml:"conditionTypeCode" json:"conditionTypeCode,omitempty"`
	StartValue        string `xml:"startValue" json:"startValue,omitempty"`
}

type ThanksAutoGetCondition struct {
	// GetCondCd
	// クーポン獲得金額条件："totalPrice" - 必須項目
	// クーポン獲得期間："grantTerm" - 必須項目
	// 初回購入ユーザー限定："serviceUseHistory" - 任意項目
	GetCondCd      string `xml:"getCondCd" json:"getCondCd,omitempty"`
	StartValue     string `xml:"startValue" json:"startValue,omitempty"`
	EndValue       string `xml:"endValue,omitempty" json:"endValue,omitempty"`
	CompOperatorCd int    `xml:"compOperatorCd" json:"compOperatorCd,omitempty"`
}
