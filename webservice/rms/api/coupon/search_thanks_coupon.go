package coupon

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/coupon"
)

// SearchThanksCoupon search coupon
func SearchThanksCoupon(ctx context.Context, clt *rms.Client, req *coupon.SearchThanksCouponRequest) (int, []coupon.ThanksCoupon, error) {
	var resp coupon.SearchThanksCouponResponse
	if err := clt.Get(ctx, "/es/1.0/thankscoupon", req, &resp, clt.TokenKey()); err != nil {
		return 0, nil, err
	}
	return resp.AllCount, resp.Content.ThanksCoupons, nil
}
