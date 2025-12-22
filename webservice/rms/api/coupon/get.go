package coupon

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
	"github.com/bububa/rakuten-go/webservice/rms/model/coupon"
)

// Get get a coupon
func Get(ctx context.Context, clt *rms.Client, req *coupon.GetRequest) (*coupon.Coupon, error) {
	var resp rmsModel.XMLResponse[*coupon.Coupon]
	if err := clt.Get(ctx, "/es/1.0/coupon/search", req, &resp, clt.TokenKey()); err != nil {
		return nil, err
	}
	return resp.Content, nil
}
