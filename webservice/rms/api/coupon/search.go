package coupon

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/coupon"
)

// Search search coupon
func Search(ctx context.Context, clt *rms.Client, req *coupon.SearchRequest) (int, []coupon.Coupon, error) {
	var resp coupon.SearchResponse
	if err := clt.Get(ctx, "/es/1.0/coupon/search", req, &resp, clt.TokenKey()); err != nil {
		return 0, nil, err
	}
	return resp.AllCount, resp.Content.Coupons, nil
}
