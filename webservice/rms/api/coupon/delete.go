package coupon

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
	"github.com/bububa/rakuten-go/webservice/rms/model/coupon"
)

// Delete delete coupon
func Delete(ctx context.Context, clt *rms.Client, req *coupon.DeleteRequest) error {
	var resp rmsModel.EmptyXMLResponse
	return clt.Post(ctx, "/es/1.0/coupon/delete", req, &resp, clt.TokenKey())
}
