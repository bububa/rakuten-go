package coupon

import (
	"context"
	"strconv"

	"github.com/bububa/rakuten-go/util"
	"github.com/bububa/rakuten-go/webservice/rms"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
	"github.com/bububa/rakuten-go/webservice/rms/model/coupon"
)

// UpdateThanksCoupon update thanks coupon
func UpdateThanksCoupon(ctx context.Context, clt *rms.Client, couponID int64, req *coupon.ThanksCouponToIssue) (int64, error) {
	apiReq := rmsModel.NewXMLRequest(req)
	var resp rmsModel.XMLResponse[*coupon.ThanksCouponID]
	if err := clt.Put(ctx, util.StringsJoin("/es/1.0/thankscoupon/", strconv.FormatInt(couponID, 10)), apiReq, &resp, clt.TokenKey()); err != nil {
		return 0, err
	}
	return resp.Content.ThanksCouponID, nil
}
