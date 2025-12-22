package coupon

import "github.com/bububa/rakuten-go/util"

type GetRequest struct {
	CouponCode string `json:"couponCode,omitempty"`
}

// Encode implements GetRequest
func (r *GetRequest) Encode() string {
	values := util.NewURLValues()
	values.Set("couponCode", r.CouponCode)
	ret := values.Encode()
	util.ReleaseURLValues(values)
	return ret
}
