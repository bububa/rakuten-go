package coupon

import (
	"context"
	"strconv"

	"github.com/bububa/rakuten-go/util"
	"github.com/bububa/rakuten-go/webservice/rms"
	rmsModel "github.com/bububa/rakuten-go/webservice/rms/model"
	"github.com/bububa/rakuten-go/webservice/rms/model/coupon"
)

// StopThanksCoupon stop thanks coupon
func StopThanksCoupon(ctx context.Context, clt *rms.Client, couponID int64) (int64, error) {
	var resp rmsModel.XMLResponse[*coupon.ThanksCouponID]
	if err := clt.Put(ctx, util.StringsJoin("/es/1.0/thankscoupon/", strconv.FormatInt(couponID, 10), "/issuestatus/stop"), nil, &resp, clt.TokenKey()); err != nil {
		return 0, err
	}
	return resp.Content.ThanksCouponID, nil
}
