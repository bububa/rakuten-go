package coupon

import (
	"encoding/xml"
)

// ThanksCouponToIssue corresponds to [XmlRoot("thanksCoupon")]
type ThanksCouponToIssue struct {
	CouponImage   string `xml:"couponImage,omitempty" json:"couponImage,omitempty"`
	CouponName    string `xml:"couponName" json:"couponName"`
	CouponCaption string `xml:"couponCaption" json:"couponCaption"`
	// DiscountType
	// 1： 定額値引き
	// 2： 定率値引き
	// 4： 送料無料
	DiscountType          int `xml:"discountType" json:"discountType"`
	DiscountFactor        int `xml:"discountFactor" json:"discountFactor"`
	CouponUnavailableTerm int `xml:"couponUnavailableTerm" json:"couponUnavailableTerm"`
	CouponTerm            int `xml:"couponTerm" json:"couponTerm"`
	MemberAvailMaxCount   int `xml:"memberAvailMaxCount" json:"memberAvailMaxCount"`
	CombineFlag           int `xml:"combineFlag" json:"combineFlag"`
	// C# [XmlArrayItem("thanksOtherCondition")] maps to Go's nested tag syntax
	ThanksOtherConditions []ThanksOtherCondition `xml:"thanksOtherConditions>thanksOtherCondition,omitempty" json:"thanksOtherConditions,omitempty"`
	// C# [XmlArrayItem("thanksAutoGetCondition")]
	ThanksAutoGetConditions []ThanksAutoGetCondition `xml:"thanksAutoGetConditions>thanksAutoGetCondition,omitempty" json:"thanksAutoGetConditions,omitempty"`
}

// ThanksCoupon embeds ThanksCouponToIssue to mimic inheritance
type ThanksCoupon struct {
	ThanksCouponToIssue
	ThanksCouponID int64  `xml:"thanksCouponId" json:"thanksCouponId,omitempty"`
	ShopID         string `xml:"shopId" json:"shopId,omitempty"`
	ShopName       string `xml:"shopName" json:"shopName,omitempty"`
	ShopURL        string `xml:"shopUrl" json:"shopUrl,omitempty"`
}

// ThanksCouponID is used for simple ID-based requests
type ThanksCouponID struct {
	XMLName        xml.Name `xml:"thanksCoupon" json:"-"`
	ThanksCouponID int64    `xml:"thanksCouponID" json:"thanksCouponID"`
}
