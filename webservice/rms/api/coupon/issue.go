package coupon

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	"github.com/bububa/rakuten-go/webservice/rms/model/coupon"
)

// Issue issue coupon
func Issue(ctx context.Context, clt *rms.Client, req *coupon.IssueRequest) (*coupon.IssuedCoupon, error) {
	var resp coupon.IssueResponse
	if err := clt.Post(ctx, "/es/1.0/coupon/issue", req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return resp.Content, nil
}
