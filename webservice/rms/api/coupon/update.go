package coupon

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
	"github.com/bububa/rakuten-go/webservice/rms/model/coupon"
)

// Update update coupon
func Update(ctx context.Context, clt *rms.Client, req *coupon.UpdateRequest) error {
	var resp rmsModel.EmptyXMLResponse
	return clt.Post(ctx, "/es/1.0/coupon/update", req, &resp, clt.TokenKey())
}
