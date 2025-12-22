package coupon

import (
	"encoding/xml"
	"strconv"

	"github.com/bububa/rakuten-go/util"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
)

type SearchRequest struct {
	CouponName      string `xml:"couponName,omitempty" json:"couponName,omitempty"`
	CouponCode      string `xml:"couponCode,omitempty" json:"couponCode,omitempty"`
	ItemURL         string `xml:"itemUrl,omitempty" json:"itemUrl,omitempty"`
	CouponStartDate string `xml:"couponStartDate,omitempty" json:"couponStartDate,omitempty"`
	CouponEndDate   string `xml:"couponEndDate,omitempty" json:"couponEndDate,omitempty"`
	Hits            int    `xml:"hits,omitempty" json:"hits,omitempty"`
	Page            int    `xml:"page,omitempty" page:"page,omitempty"`
}

// Encode implements GetRequest
func (r *SearchRequest) Encode() string {
	values := util.NewURLValues()
	if r.CouponName != "" {
		values.Set("couponName", r.CouponName)
	}
	if r.CouponCode != "" {
		values.Set("couponCode", r.CouponCode)
	}
	if r.ItemURL != "" {
		values.Set("itemUrl", r.ItemURL)
	}
	if r.CouponStartDate != "" {
		values.Set("couponStartDate", r.CouponStartDate)
	}
	if r.CouponEndDate != "" {
		values.Set("couponEndDate", r.CouponEndDate)
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

type SearchResponse struct {
	rmsModel.XMLResponse[SearchedCoupons]
	AllCount int `xml:"allCount"`
}

type SearchedCoupons struct {
	XMLName xml.Name `xml:"coupons"`
	Coupons []Coupon `xml:"coupon"`
}
