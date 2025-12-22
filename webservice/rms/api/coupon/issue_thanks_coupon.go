package coupon

import (
	"context"

	"github.com/bububa/rakuten-go/webservice/rms"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
	"github.com/bububa/rakuten-go/webservice/rms/model/coupon"
)

// IssueThanksCoupon issue thanks coupon
func IssueThanksCoupon(ctx context.Context, clt *rms.Client, req *coupon.ThanksCouponToIssue) (int64, error) {
	apiReq := rmsModel.NewXMLRequest(req)
	var resp rmsModel.XMLResponse[*coupon.ThanksCouponID]
	if err := clt.Post(ctx, "/es/1.0/thankscoupon", apiReq, &resp, clt.TokenKey()); err != nil {
		return 0, err
	}
	return resp.Content.ThanksCouponID, nil
}
